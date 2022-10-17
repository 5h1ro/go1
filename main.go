package main

import (
	"fmt"
	"os"
	_ "strconv"
)

type Person struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var people []Person = []Person{
		{
			Id:        "1",
			Nama:      "Nurhakiki",
			Alamat:    "Madiun",
			Pekerjaan: "Programmer",
			Alasan:    "Tuntutan",
		},
		{
			Id:        "2",
			Nama:      "Bima",
			Alamat:    "Madiun",
			Pekerjaan: "Fotografer",
			Alasan:    "Hobi",
		},
	}

	var args = os.Args
	fmt.Println(getPersonById(args[1], people))

}

func getPersonById(id string, people []Person) string {

	var conditional bool = false

	var peopleIndex = 0

	for index, value := range people {
		if value.Id == id {
			conditional = true
			peopleIndex = index
			break
		}
	}

	if conditional {
		return fmt.Sprintf("Data ditemukan => %+v", people[peopleIndex])
	} else {
		return "Data tidak ditemukan"
	}
}
