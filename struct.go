package main

import (
	"fmt"
)

// Struct Person dengan field Name dan Age
type Person struct {
	Name string
	Age  int
}

// Fungsi CreatePerson yang mengembalikan pointer ke struct Person
func CreatePerson(name string, age int) *Person {
	if age < 0 {
		panic("Invalid age") // Menghasilkan panic jika umur kurang dari 0
	}
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Fungsi PrintPerson yang mencetak informasi person dengan menggunakan defer
func PrintPerson(p *Person) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// Fungsi HandlePerson yang memanggil CreatePerson dan PrintPerson
func HandlePerson(name string, age int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	person := CreatePerson(name, age)
	PrintPerson(person)
}

func main() {
	// Memanggil HandlePerson dengan data valid
	fmt.Println("Handling person with valid data:")
	HandlePerson("Toto", 20)

	// Memanggil HandlePerson dengan data yang memicu panic
	fmt.Println("\nHandling person with invalid data:")
	HandlePerson("Toto", -1)
}
