package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time
	now = time.Now()
	date := now.Second() + now.Minute() + now.Hour() + now.Day() + int(now.Month()) + now.YearDay()
	fmt.Println("Now:", date)
}
