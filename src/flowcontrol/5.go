package main

import (
	"fmt"
	"time"
)

var days = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func main() {
	fmt.Println("What day is today ?")
	switch today := time.Now().Weekday(); int(today)  {
	default:
		fmt.Println(days[today - 1], "\nnumber:", int(today))
	}
}
