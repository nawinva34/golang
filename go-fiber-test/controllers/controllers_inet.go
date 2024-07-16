package controllers

import (
	"errors"
	"go-fiber-test/database"
	"go-fiber-test/models"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BodyParser(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john : log can show timestamp
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func Params(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func Query(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func Validation(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

// 5.1
func Factorial(c *fiber.Ctx) error {
	numStr := c.Params("num")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid number",
		})
	}

	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	return c.JSON(fiber.Map{
		"number":    num,
		"factorial": result,
	})
}

// 5.2
func ConvertTaxIDToASCII(c *fiber.Ctx) error {
	taxID := c.Query("tax_id")
	asciiValues := convertToASCII(taxID)
	return c.JSON(fiber.Map{
		"tax_id": taxID,
		"ascii":  asciiValues,
	})
}
func convertToASCII(text string) []int {
	asciiValues := make([]int, len(text))
	for i, char := range text {
		asciiValues[i] = int(char)
	}
	return asciiValues
}

// 6 Register
var validate = validator.New()

func init() {
	validate.RegisterValidation("username_validate", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(fl.Field().String())
	})
	validate.RegisterValidation("web_validate", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-z0-9.-]+$`).MatchString(fl.Field().String())
	})
	validate.RegisterValidation("tel", func(fl validator.FieldLevel) bool {
		tel := fl.Field().String()
		return regexp.MustCompile(`^\d{10}$`).MatchString(tel)
	})
}

func Register(c *fiber.Ctx) error {
	user := new(models.UserRegister)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := validate.Struct(user)
	if err != nil {
		fieldErrors := make(map[string]string)

		validationErrors := err.(validator.ValidationErrors)
		errorMessages := map[string]string{
			"username":  "ใช้อักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข (0-9) และเครื่องหมาย (_), (-) เท่านั้น เช่น Example_01",
			"password":  "ความยาว 6-20 ตัวอักษร",
			"tel":       "กรุณาใส่หมายเลขโทรศัพท์ 10 หลัก",
			"web_name":  "ความยาว 2-30 ตัวอักษร",
			"web_name2": "ใช้อักษรภาษาอังกฤษตัวเล็ก (a-z), ตัวเลข (0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นขีด (-) ห้ามเว้นวรรค และห้ามใช้ภาษาไทย",
		}

		for _, e := range validationErrors {
			fieldName := strings.ToLower(e.Field())
			if msg, ok := errorMessages[fieldName]; ok {
				fieldErrors[fieldName] = msg
			} else {
				fieldErrors[fieldName] = fieldName + " is required"
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors occurred",
			"errors":  fieldErrors,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
		"result":  user,
	})
}

// CRUD DB
func GetDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at is NOT NULL").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDogsRangeCountByDogId(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("dog_id > ? && dog_id < ?", 50, 100).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs
	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJsonSummary(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNone := 0

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNone++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:        dataResults,
		Name:        "golang-test",
		Count:       len(dogs),
		Sum_red:     sumRed,
		Sum_green:   sumGreen,
		Sum_pink:    sumPink,
		Sum_noColor: sumNone,
	}
	return c.Status(200).JSON(r)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	type ResultData struct {
		Data  []m.DogsRes `json:"data"`
		Name  string      `json:"name"`
		Count int         `json:"count"`
	}
	r := ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs),
	}
	return c.Status(200).JSON(r)
}

// CRUD companies
func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Companies

	db.Find(&companies)
	return c.Status(200).JSON(companies)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Companies

	result := db.Find(&company, "id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&company)
}

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Companies

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Companies
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Companies
	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// CRUD profiles
func GetProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var profiles []m.UserProfiles

	db.Find(&profiles)
	return c.Status(fiber.StatusOK).JSON(profiles)
}

func GetProfile(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var profile m.UserProfiles

	result := db.Where("employee_id = ?", search).First(&profile)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).SendString("User profile not found")
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}
func AddProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.UserProfiles

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(profile); err != nil {
		fieldErrors := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Email" && e.Tag() == "email" {
				fieldErrors[strings.ToLower(e.Field())] = "Invalid email"
			} else if e.Field() == "Age" && e.Tag() == "min" {
				fieldErrors[strings.ToLower(e.Field())] = "Age must be greater than or equal 18"
			} else {
				fieldErrors[strings.ToLower(e.Field())] = e.Field() + " is required"
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors occurred",
			"errors":  fieldErrors,
		})
	}

	var existingEmpId m.UserProfiles
	if err := db.Where("employee_id = ?", profile.EmployeeID).First(&existingEmpId).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Employee ID already exists."})
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	db.Create(&profile)

	return c.Status(fiber.StatusCreated).JSON(profile)
}

func UpdateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile m.UserProfiles

	if err := db.Where("id = ?", id).First(&profile).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User profile not found")
	}

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if profile.EmployeeID != "" {
		var existingEmpId m.UserProfiles
		if err := db.Where("employee_id = ?", profile.EmployeeID).First(&existingEmpId).Error; err == nil && existingEmpId.ID != profile.ID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Employee ID already exists."})
		}
	}

	db.Save(&profile)
	return c.Status(fiber.StatusOK).JSON(profile)
}

func RemoveProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile m.UserProfiles

	if err := db.Where("id = ?", id).First(&profile).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User profile not found")
	}

	db.Delete(&profile)
	return c.SendStatus(fiber.StatusNoContent)
}

func GetProfileAnyAges(c *fiber.Ctx) error {
	db := database.DBConn
	var profiles []m.UserProfiles
	var result m.UserProfileAgesResult
	var dataResult []m.UserProfileResult

	db.Find(&profiles)

	for _, profile := range profiles {
		var genStr string

		switch {
		case profile.Age < 24:
			genStr = "Gen Z"
		case profile.Age >= 24 && profile.Age <= 41:
			genStr = "Gen Y"
		case profile.Age >= 42 && profile.Age <= 56:
			genStr = "Gen X"
		case profile.Age >= 57 && profile.Age <= 75:
			genStr = "Baby Boomer"
		default:
			genStr = "G.I. Generation"
		}

		p := m.UserProfileResult{
			EmployeeID: profile.EmployeeID,
			Name:       profile.Name,
			LastName:   profile.LastName,
			Birthday:   profile.Birthday,
			Age:        profile.Age,
			Email:      profile.Email,
			Tel:        profile.Tel,
			Gen:        genStr,
		}

		dataResult = append(dataResult, p)

		switch genStr {
		case "Gen Z":
			result.SumGenZ++
		case "Gen Y":
			result.SumGenY++
		case "Gen X":
			result.SumGenX++
		case "Baby Boomer":
			result.SumBabyBoomer++
		case "G.I. Generation":
			result.SumGI++
		}
	}

	result.Data = dataResult
	result.Count = len(profiles)
	result.Name = "profile-ages"
	return c.Status(fiber.StatusOK).JSON(result)
}

func SearchProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var profiles []m.UserProfiles

	db.Where("(employee_id = ? OR name LIKE ? OR last_name LIKE ?) AND deleted_at IS NULL", search, "%"+search+"%", "%"+search+"%").Find(&profiles)

	return c.Status(fiber.StatusOK).JSON(profiles)
}
