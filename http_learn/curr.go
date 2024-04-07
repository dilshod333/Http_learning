package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func changeCurr(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error ", http.StatusNotFound)
			return
		}

		dollarStr := r.Form.Get("dollar")

		dollarInput, err := strconv.Atoi(dollarStr)
		if err != nil {
			http.Error(w, "Invalid input for dollar amount", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, "Invalid input for converted result", http.StatusBadRequest)
			return
		}

		convertedSum := dollarInput * 12572

		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Currency Converter</title>
		</head>
		<body>
			<h1>Welcome TO Converter</h1>
			<form method="POST">
				<label for="dollar">Enter the money:</label>
				<input type="text" name="dollar" id="dollar" placeholder="Enter the money $" value="%s">
				<br>
				<label for="sum">Converted result:</label>
				<input type="text" name="sum" id="sum" placeholder="Converted result" value="%d sum ">
				<br>
				<button type="submit" name="submit">Convert</button>
			</form>
		</body>
		</html>
		`, dollarStr, convertedSum)

		return
	}

	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Currency Converter</title>
	</head>
	<body>
		<h1>Welcome TO Converter</h1>
		<form method="POST">
			<label for="dollar">Enter the money:</label>
			<input type="text" name="dollar" id="dollar" placeholder="Enter the money">
			<br>
			<label for="sum">Converted result:</label>
			<input type="text" name="sum" id="sum" placeholder="Converted result">
			<br>
			<button type="submit" name="submit">Convert</button>
		</form>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", changeCurr)
	http.ListenAndServe(":8080", nil)
}
