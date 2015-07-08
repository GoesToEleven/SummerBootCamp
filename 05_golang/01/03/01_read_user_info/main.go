package main

import (
	"fmt"
	"strings"
	"crypto/md5"
	"io"
	"encoding/hex"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}

func main() {
	fmt.Print("Enter first name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("Enter email: ")
	var email string
	fmt.Scanln(&email)
	var gravatarHash = getGravatarHash(email)
	fmt.Println(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<img src="http://www.gravatar.com/avatar/` + gravatarHash + `" alt="user picture">
</body>
</html>`)

}
