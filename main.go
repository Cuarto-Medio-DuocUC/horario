package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Cuarto-Medio-DuocUC/horario/table"
)

//go:embed fixtures/horario.json
var rawData embed.FS

func main() {

    file, err := rawData.Open("fixtures/horario.json")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

	today := int(time.Now().Weekday())

	var wD int
	flag.IntVar(&wD, "dia", today, "Día de la semmana 1 Lunes a 6 Sábado")
	flag.Parse()

	if wD > int(Saturday) || wD < int(Sunday) {
		fmt.Printf("El día %d no es parte del rango de la semana\n", wD)
		os.Exit(1)
	}

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	weekDay := WeekDay(wD)
    coursesData, err := ParseFile(rawData)
    if err != nil {
        log.Fatal(err)
    }

	courses, ok := coursesData[weekDay]
	if !ok {
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
