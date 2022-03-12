package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type WeekDay int

const (
	Monday WeekDay = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Professor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Course struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Professor
}

type Classroom struct {
	TimeInit string `json:"timeInit"`
	TimeEnd  string `json:"timeEnd"`
	Campus   string `json:"campus"`
	Room     string `json:"room"`
	Course
}

type ClassDay struct {
	Day        WeekDay     `json:"day,string"`
	Classrooms []Classroom `json:"classrooms"`
}

func main() {
	file, err := os.Open("./fixtures/horario.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file opened ok")

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var classweek []ClassDay
	cw := make(map[WeekDay][]Classroom)
	err = json.Unmarshal(data, &classweek)
	if err != nil {
		log.Fatal(err)
	}
    for _, day := range classweek {
		cw[day.Day] = day.Classrooms
	}

}
