package main

import (
	"encoding/json"
)

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (wd WeekDay) String() string {
	daysNames := []string{
		"Domingo",
		"Lunes",
		"Martes",
		"Miércoles",
		"Jueves",
		"Viernes",
		"Sábado",
	}

	if Sunday <= wd && wd <= Saturday {
		return daysNames[wd]
	}
	return ""
}

type Professor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Course struct {
	Name string `json:"name"`
	Code string `json:"code"`
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

func RequestedCourses(data []byte, wd WeekDay) ([]Classroom, error) {
	var weekClasses []ClassDay
	err := json.Unmarshal(data, &weekClasses)
	if err != nil {
		return []Classroom{}, err
	}

	var dayclasses []Classroom

	for _, day := range weekClasses {
		if day.Day == wd {
			return day.Classrooms, nil
		}
	}

	return dayclasses, nil
}
