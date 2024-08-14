package main

import "fmt"

func main() {
	animal := getAnimals()
	result := preparePrint(animal)
	fmt.Println(result)

}

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {

	animals := []Animal{
		{Type: "Cat", Name: "Watson", Age: 10},
		{Type: "Dog", Name: "Kirill", Age: 27},
		{Type: "Bear", Name: "Zver", Age: 15},
	}
	return animals
}

func preparePrint(animals []Animal) string {
	var str string
	for _, animal := range animals {
		str += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return str
}
