package main

import (
	"fmt"
	"goTestProject/animal"
	"goTestProject/realization"
)

func main() {
	animals := []animal.Animal{realization.Dog{}, realization.Cat{}, realization.Llama{}, realization.JavaProgrammer{}}
	for _, an := range animals {
		fmt.Println(an.Speak())
	}
}
