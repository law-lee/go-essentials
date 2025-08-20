package datetime

import (
	"fmt"
	"log"
	"time"
)

func Run() {
	now := time.Now()
	fmt.Printf("time now=%v\n", now)

	newTime := now.Add(100 * time.Second)
	fmt.Printf("增加100秒之后时间: %v\n", newTime)

	newTime = now.AddDate(1, 1, 1)
	fmt.Printf("增加1年又1个月又1天后时间：%v\n", newTime)

	unixTime := now.Unix()
	fmt.Printf("转化成unix时间: %d\n", unixTime)

	year, month, day := newTime.Date()
	fmt.Printf("转化成年月日：%d, %d, %d\n", year, month, day)

	// format time
	t := time.Date(2017, 9, 4, 3, 38, 45, 0, time.UTC)
	fmt.Println(t.Format("2006-01-02 15:04:05.000 MST"))
	fmt.Println(t.Format("2006-01-02 15pm"))
	fmt.Println(t.Format("Jan 06 Mon 2 01"))
	fmt.Println(t.Format("January 6 Mon 2 1"))
	fmt.Println(t.Format("Month: Jan '1', '01', _2"))

	// parse time
	s := "2017-04-09 03:38:45.000 UTC"
	t, err := time.Parse("2006-01-02 15:04:05.000 MST", s)
	if err != nil {
		log.Fatalf("time.Parse() failed wiht '%s'\n", err)
	}
	fmt.Printf("year: %d, month: %d, day: %d\n", t.Year(), t.Month(), t.Day())

	// compare time and date
	var date1 = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    var date2 = time.Date(2017, time.July, 25, 16, 22, 42, 123, time.UTC)
    var date3 = time.Date(2017, time.July, 25, 16, 22, 42, 123, time.UTC)

    fmt.Println(date1.Before(date2)) // true, because date1 is before date2
    fmt.Println(date1.After(date2)) // false, because date1 is not after date2

    fmt.Println(date2.Before(date1)) // false, because date2 is not before date1
    fmt.Println(date2.After(date1)) // true, because date2 is after date1

    fmt.Println(date1.Equal(date2)) // false, not the same moment
    fmt.Println(date1.Equal(date3)) // true, different objects but representing the exact same time

    fmt.Println(date1 != date2) // true, different moments
    fmt.Println(date1 != date3) // false, not different moments

    fmt.Println(date1.After(date3)) // false, because date1 is not after date3 (that are the same)
    fmt.Println(date1.Before(date3)) // false, because date1 is not before date3 (that are the same)

    fmt.Println(!(date1.Before(date3))) // true, because date1 is not before date3
    fmt.Println(!(date1.After(date3))) // true, because date1 is not after date3
}
