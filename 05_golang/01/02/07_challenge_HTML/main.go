package main

import "fmt"
import "os"
import "strconv"

const miTokm = 1.60934

func main() {
	number, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")
	fmt.Printf("Miles: %f<br>\n", number)
	fmt.Printf("KM: %f<br>\n", number*miTokm)
	fmt.Println("</body>")
	fmt.Println("</html>")
}

/*
at terminal

go install
07_challenge_HTML 43 > tmpfile.html

*/