package main

import (
	"fmt"
	"github.com/do-work/learnGo/calendar"
	"log"
)

func main() {
	date := calendar.Date{}

	err := date.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}
