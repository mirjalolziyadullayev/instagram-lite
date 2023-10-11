package handlers

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, 
	`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<style>
			body {
				background-color: aquamarine;
			}
		</style>
		<p>Hello World</p>
	</body>
	</html>
	`)
}