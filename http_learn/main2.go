// // // package main

// // // import (
// // // 	"fmt"
// // // 	"net/http"
// // // )

// // // func hello(w http.ResponseWriter, r *http.Request) {
// // // 	fmt.Fprint(w, "Hello Dilshodbek Welocome to your first golang webserver: ")

// // // }
// // // func dilshod(w http.ResponseWriter, r *http.Request) {
// // // 	fmt.Fprint(w, "Salom Dilshod nima xizmatlar ")

// // // }
// // // func main() {
// // // 	http.HandleFunc("/", hello)
// // // 	http.HandleFunc("dilshod", dilshod)
// // // 	http.ListenAndServe(":8080", nil)
// // // }

// // package main

// // import (
// // 	"fmt"
// // 	"net/http"
// // )

// // func result(w http.ResponseWriter, r *http.Request) {
// // 	switch r.URL.Path {
// // 	case "/":
// // 		fmt.Fprint(w, "Hey it main page ")
// // 	case "/Dilshod/":
// // 		fmt.Fprint(w, "Hey Dilshod how are you doing ")
// // 	default:
// // 		http.Error(w, "Not found", http.StatusNotFound)

// // 	}

// // }

// // func main() {
// // 	http.HandleFunc("/", result)
// // 	http.HandleFunc("/Dilshod", result)
// // 	http.ListenAndServe(":8080", nil)

// // }
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func result(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		fmt.Fprintf(w, `
// 			<!DOCTYPE html>
// 			<html lang="en">
// 			<head>
// 				<meta charset="UTF-8">
// 				<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 				<title>Main Page</title>
// 				<style>
// 					body {
// 						background-color: #f0f0f0;
// 						font-family: Arial, sans-serif;
// 						margin: 0;
// 						padding: 0;
// 					}
// 					.container {
// 						max-width: 800px;
// 						margin: 50px auto;
// 						padding: 20px;
// 						background-color: #black;
// 						border-radius: 10px;
// 						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
// 					}
// 					h1 {
// 						color: #blue;
// 					}
// 				</style>
// 			</head>
// 			<body>
// 				<div class="container">
// 					<h1>Hey it's the main page</h1>
// 				</div>
// 			</body>
// 			</html>
// 		`)
// 	case "/Dilshod/":
// 		fmt.Fprintf(w, `
// 			<!DOCTYPE html>
// 			<html lang="en">
// 			<head>
// 				<meta charset="UTF-8">
// 				<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 				<title>Greeting</title>
// 				<style>
// 					body {
// 						background-color: #f0f0f0;
// 						font-family: Arial, sans-serif;
// 						margin: 0;
// 						padding: 0;
// 					}
// 					.container {
// 						max-width: 800px;
// 						margin: 50px auto;
// 						padding: 20px;
// 						background-color: #000000;
// 						border-radius: 10px;
// 						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
// 					}
// 					h1 {
// 						color: #333;
// 					}
// 				</style>
// 			</head>
// 			<body>
// 				<div class="container">
// 					<h1>Hey Dilshod, how are you doing?</h1>
// 					<h2> How can I help you now</h2>
// 				</div>
// 			</body>
// 			</html>
// 		`)
// 	default:
// 		http.Error(w, "Not found", http.StatusNotFound)
// 	}
// }

//	func main() {
//		http.HandleFunc("/", result)
//		http.ListenAndServe(":8080", nil)
//	}
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var target int
var attempts int
var maxAttempts = 10

func generateTarget() {
	rand.Seed(time.Now().UnixNano())
	target = rand.Intn(100) + 1
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if attempts == 0 {
			generateTarget()
			attempts = maxAttempts
		}

		fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Number Guessing Game</title>
                <style>
                    body {
                        font-family: Arial, sans-serif;
                        text-align: center;
                        margin-top: 50px;
                    }
                    input[type="text"] {
                        padding: 10px;
                        font-size: 16px;
                    }
                    input[type="submit"] {
                        padding: 10px 20px;
                        font-size: 16px;
                        background-color: #007bff;
                        color: #fff;
                        border: none;
                        cursor: pointer;
                    }
                </style>
            </head>
            <body>
                <h1>Guess the Number (1-100)</h1>
                <form method="POST">
                    <input type="text" name="guess" placeholder="Enter your guess" autofocus>
                    <input type="submit" value="Guess">
                </form>
                <p>Attempts remaining: %d</p>
            </body>
            </html>
        `, attempts)
	} else if r.Method == "POST" {
		guess, _ := strconv.Atoi(r.FormValue("guess"))
		if guess == target {
			fmt.Fprintf(w, "<h2>Congratulations! You guessed the number correctly.</h2>")
			attempts = 0
		} else if guess < target {
			fmt.Fprintf(w, "<h2>Too low. Try again!</h2>")
		} else {
			fmt.Fprintf(w, "<h2>Too high. Try again!</h2>")
		}
		attempts--
		if attempts == 0 {
			fmt.Fprintf(w, "<p>Out of attempts. The correct number was %d.</p>", target)
		}
	}
}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
