package main

import (
	"time"
	"fmt"
	"reflect"
)

func main() {
	var userBday string
	fmt.Print("Enter your birthday (month/day/year): ")
	fmt.Scanln(&userBday)

	const shortForm = "01/02/2006"
	date1, _ := time.Parse(shortForm, userBday)
	fmt.Println(date1)

	date2 := time.Now()
	difference := date2.Sub(date1)

//	fmt.Println(difference)
//	fmt.Println(reflect.TypeOf(difference))
	fmt.Println(difference.Hours())
	fmt.Println(difference.Minutes())


	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Enter your social security number: ")
	time.Sleep(1200 * time.Millisecond)
	fmt.Println("And your credit card number: ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("And your bank account: ")
	time.Sleep(900 * time.Millisecond)
	fmt.Println("Just kidding.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You've been alive ...")
	time.Sleep(300 * time.Millisecond)
	fmt.Println(difference.Seconds(),"seconds")
	time.Sleep(300 * time.Millisecond)
	fmt.Println(difference.Minutes(),"minutes")
	time.Sleep(300 * time.Millisecond)
	fmt.Println(difference.Hours()/24,"hours")
	time.Sleep(300 * time.Millisecond)
	fmt.Println(difference.Hours()/24/30,"days")
	time.Sleep(300 * time.Millisecond)
	fmt.Println(difference.Hours()/24,"months")



}
