package controllers

import (
	"go-fiber-test/models"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
