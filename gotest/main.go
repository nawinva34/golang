package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Switch() // num0
	// num1() // num1
	// num1_2() //num1.2
	// num2() //num2
	// num3() //num3
	num4() //num4
	// num5() //num5
	// num6() //num6
	// sp()
}

// num0
func Switch() {
	i := 2

	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Your i not in case.")
	}
}

// num1 =============
func num1() {
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
			count++
		}
	}

	fmt.Printf("\nจำนวนเลขที่หาร 3 ลงตัวที่อยู่ระหว่าง 1 ถึง 100 คือ : %d\n", count)
}

// num1.2 ===========
func num1_2() {
	base := 20
	exponent := 0
	result := exponent_fnc(base, exponent)
	fmt.Printf("%d ยกกำลัง %d คือ %d\n", base, exponent, result)
}
func exponent_fnc(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// num2 ==============
func num2() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	min := findMin(x)
	max := findMax(x)
	fmt.Printf("Minimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)
}

func findMin(x []int) int {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}

func findMax(x []int) int {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

// num3 ==============
func num3() {
	totalNines := countNines()
	totalNinesInput := totalNinesInput(10000)
	fmt.Printf("มีเลข 9 จำนวน %d ตัวในระหว่าง 1 ถึง 1000.\n", totalNines)
	fmt.Printf("มีเลข 9 จำนวน %d ตัวในระหว่าง 1 ถึง 10000.\n", totalNinesInput)
}

func countNines() int {
	count := 0
	for i := 1; i <= 1000; i++ {
		str := strconv.Itoa(i)
		for _, v := range str {
			if v == '9' {
				count++
			}
		}
	}
	return count
}

func totalNinesInput(x int) int {
	count := 0
	for i := 1; i <= x; i++ {
		str := strconv.Itoa(i)
		for _, v := range str {
			if v == '9' {
				count++
			}
		}
	}
	return count
}

// num4 ================
func num4() {
	myWords := "AW SOME GO!"
	result := ""
	input := "ine t"

	for _, char := range myWords {
		if char != ' ' {
			result += string(char)
		}
	}

	fmt.Println(result)
	fmt.Println(cutText(input))
}

func cutText(input string) string {
	result := ""

	for _, char := range input {
		if char != ' ' {
			result += string(char)
		}
	}

	return result
}

// num5 ================
type Employee struct {
	Name    string
	Age     string
	Address string
}

func num5() {
	peoples := map[string]map[string]string{
		"emp_01": {
			"Name":    "Mr. Tim Carry",
			"Age":     "25",
			"Address": "142rd roads, Virginia, 22202",
		},
		"emp_02": {
			"Name":    "Ms. Laura MaCoy",
			"Age":     "22",
			"Address": "123/western Street, New York, 12304",
		},
		"emp_03": {
			"Name":    "Mr. John Ashley",
			"Age":     "24",
			"Address": "152rd roads, California, 53320",
		},
	}

	for _, value := range peoples {
		employee := Employee{
			Name:    value["Name"],
			Age:     value["Age"],
			Address: value["Address"],
		}
		fmt.Printf("\nName: %s (Age: %s)\nAddress: %s\n\n", employee.Name, employee.Age, employee.Address)
	}
}

// num6 ================
type Company struct {
	Name    string
	Age     string
	Address string
}

func num6() {
	var c Company
	c.Name = "Tim Carry"
	c.Age = "25"
	c.Address = "142rd roads, Virginia, 22202"
	fmt.Printf("\nName: %s (Age: %s)\nAddress: %s\n\n", c.Name, c.Age, c.Address)
}

func sp() {
	x := 6

	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// for i := 1; i <= x; i++ {
	// 	for j := 1; j <= x-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
}
