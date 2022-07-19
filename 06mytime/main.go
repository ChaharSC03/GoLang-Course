package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the time package")

	presentTime :=time.Now()

	fmt.Println("Present time is:",presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	// the above format is mandotary as--> 01-02-2006 
	// and for days alawys put--> Monday

	createDate :=time.Date(2020,time.March,04,22,15,0,0,time.UTC)

	// with format
	fmt.Println("created date is :",createDate.Format("01-02-2006 Monday"))



}
