package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var attemp int
var targ int
var mAttemp int = 10

func randomIt() {
	rand.Seed(time.Now().UnixNano())
	targ = rand.Intn(100) + 1
}
func check(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println(r.Method)
		if attemp == 0 {
			randomIt()
			attemp = mAttemp
		}

		fmt.Fprintf(w, `
			<!Doctype html>
			<html>
			<head>
				<title>NUmber Guessing Game</title>
				<style>
					body {
						font-family: Arial,sans-sarif;
						text-align; center;
						margin-top: 50px;
						input[type="text"] {
							padding: 10px;
							font-size: 16px;
						}
					}
				</style>
			</head>
			<body>
					<h1>Guess The number (1-100)</h1>
					<form method="POST>
						<input type="text" name="guess" placeholder="Enter your guess">
						<input type="submit" value="Guess">
					</form>
					<p> Attempts remaining %d</p>
			</body>
			</html>
			`, attemp)

	} else if r.Method == "POST" {
		fmt.Println(r.Method)
		guess, _ := strconv.Atoi(r.FormValue("guess"))
		if guess == targ {
			fmt.Fprintf(w, "<h2>Congrulations Dilshod 😎</h2>")
			attemp = 0
		} else if guess < targ {
			fmt.Fprintf(w, "<h2>Too low try again bro 😅</h2>")
		} else {
			fmt.Fprintf(w, "<h2>Too high try again bro 😅</h2>")
		}
		attemp--
		if attemp == 0 {
			fmt.Fprintf(w, "<p>Out of attempts</p>")
		}
	}

}

func main() {
	http.HandleFunc("/", check)
	http.ListenAndServe(":8080", nil)
}
