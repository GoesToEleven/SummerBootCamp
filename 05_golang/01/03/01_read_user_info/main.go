package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
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
	//	fmt.Print("Enter email: ")
	//	fmt.Fprint(os.Stdout, "Enter email: ")
	fmt.Fprint(os.Stderr, "Enter email: ")
	var email string
	fmt.Scanln(&email)
	var gravatarHash = getGravatarHash(email)
	//	fmt.Fprint(os.Stdout, `<!DOCTYPE html>
	fmt.Print(`<!DOCTYPE html>
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
