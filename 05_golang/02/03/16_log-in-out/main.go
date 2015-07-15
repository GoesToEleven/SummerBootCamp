package main

import (

	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		io.WriteString(res, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>

<h1>LOG IN</h1>
<form method="post">
    <input type="text" name="userName" id="userName">
    <input type="text" name="password" id="password">
    <input type="submit">
</form>

</body>
</html>`)
	})

	http.ListenAndServe(":9000", nil)
}
