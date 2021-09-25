package main

import (
	"fmt"
	"strings"
)

type VowelsFinder  interface {
	FindVowels() []rune
}

type MyString string

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId, basicpay, pf int
}

type Contract struct {
	empId, basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s [] SalaryCalculator)  {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func (ms MyString) FindVowels() []rune {
	const vowelsString = "ueoai"
	var vowels []rune
	for _, c := range ms {
		if strings.Contains(vowelsString, string(c)) {
			vowels = append(vowels, c)
		}
	}
	return vowels
}

func main() {
	var i VowelsFinder = MyString("Nguyen Manh")
	vowels := i.FindVowels()
	for i, v := range vowels {
		fmt.Printf("vowels %d is %c \n", i, v)
	}

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
