package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	createFile()
	writeFile()
	readFile()
}

func createFile() {
	var file, err = os.Create("./files/person.json")
	defer file.Close() //Close file created
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Created Successfully")
}

func writeFile() {
	var file, err = os.OpenFile("./files/person.json", os.O_WRONLY, 0644)
	defer file.Close() //Close file after write
	if err != nil {
		log.Fatal(err)
	}

	var people = []person{
		{
			ID:   1,
			Name: "Marco Días",
			Age:  19,
		},
		{
			ID:   2,
			Name: "Juan Perez",
			Age:  20,
		},
		{
			ID:   3,
			Name: "Frankie Rivers",
			Age:  15,
		},
	}

	var b, errC = json.Marshal(people)
	if errC != nil {
		log.Fatal(errC)
	}
	file.Write(b)

	fmt.Println("File Updated Successfully")
}

func readFile() {
	var file, err = os.OpenFile("./files/person.json", os.O_RDONLY, 0644)
	defer file.Close() //Close file after read
	if err != nil {
		log.Fatal(err)
	}

	var content, errC = ioutil.ReadAll(file)
	if errC != nil {
		log.Fatal(errC)
	}
	fmt.Println("Content:", string(content))
}
