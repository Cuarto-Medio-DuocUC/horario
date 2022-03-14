package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Cuarto-Medio-IFPT-DuocUC-2022/horario/table"
)

func main() {
	today := int(time.Now().Weekday())

	var wD int
	flag.IntVar(&wD, "dia", today, "Día de la semmana 1 Lunes a 6 Sábado")
	flag.Parse()

	if wD > int(Saturday) || wD < int(Sunday) {
		fmt.Printf("El día %d no es parte del rango de la semana\n", wD)
		os.Exit(1)
	}

	file, err := os.Open("./fixtures/horario.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	weekDay := WeekDay(wD)
	courses, err := RequestedCourses(data, weekDay)
	if err != nil {
		log.Fatal(err)
	}
	if len(courses) == 0 {
		fmt.Printf("No hay clases hoy %s\n", weekDay)
		os.Exit(0)
	}

	table := table.NewTable(os.Stdout)
	table.MergeDay(len(courses) > 1)

	for _, course := range courses {
		datum := []string{
			weekDay.String(),
			course.TimeInit,
			course.TimeEnd,
			course.Course.Code,
			course.Course.Name,
			course.Campus,
			course.Room,
		}
		table.Append(datum)
	}

	fmt.Println()
	table.Render()
}
