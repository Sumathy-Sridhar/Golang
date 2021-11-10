package main

import (
	"fmt"
	"time"
)

func main() {

	// to get current location and time
	tm := time.Now()
	fmt.Println("Location: ", tm.Location(), "Time: ", tm)

	//to retrieve specific location and time (using loadfunction)
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Location : ", location, " Time : ", tm.In(location)) // America/New_York

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shangha

	//to use Weekday and YearDay function
	t1, _ := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	fmt.Println(t1.YearDay()) // 315
	fmt.Println(t1.Weekday()) // Wednesday

	//current time in various format
	currentTime := time.Now()
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))
	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))

	fmt.Println("Current Time in String: ", currentTime.String())
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		fmt.Println(err)
	}
	hr, min, sec := timeStamp.Clock()

	fmt.Println("\n**************")
	fmt.Println("Year   :", currentTime.Year())
	fmt.Println("Month  :", currentTime.Month())
	fmt.Println("Day    :", currentTime.Day())
	fmt.Println("Hour   :", hr)
	fmt.Println("Min    :", min)
	fmt.Println("Sec    :", sec)

	fmt.Println("\n**************")
	year, month, day := time.Now().Date()
	fmt.Println("Year   :", year)
	fmt.Println("Month  :", month)
	fmt.Println("Day    :", day)

	fmt.Println("\n**************")
	t := time.Now()

	y := t.Year()
	mon := t.Month()
	d := t.Day()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	n := t.Nanosecond()

	fmt.Println("Year   :", y)
	fmt.Println("Month   :", mon)
	fmt.Println("Day   :", d)
	fmt.Println("Hour   :", h)
	fmt.Println("Minute :", m)
	fmt.Println("Second :", s)
	fmt.Println("Nanosec:", n)

	//convert time to MST.EST
	tme, err := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tme)
	loc, err = time.LoadLocation("MST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)
	tme = tme.In(loc)
	fmt.Println(t.Format(time.RFC822))

	loc, err = time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)
	tme = tme.In(loc)
	fmt.Println(tme.Format(time.RFC822))

}
