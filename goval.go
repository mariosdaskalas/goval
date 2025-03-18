package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the Golang credit card validator.")
	http.HandleFunc("/validate", validateHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	if number == "" {
		http.Error(w, "No number provided", http.StatusBadRequest)
		return
	}

	// Parse the number
	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Validate the number
	if isValidLuhn(num) {
		fmt.Fprintf(w, "The number %s is valid according to the Luhn algorithm\n", number)
	} else {
		fmt.Fprintf(w, "The number %s is not valid according to the Luhn algorithm\n", number)
	}

}

// isValidLuhn checks if a number is valid according to the Luhn algorithm
func isValidLuhn(number int64) bool {
	str := strconv.FormatInt(number, 10)
	var sum int
	double := false

	// Iterate over the digits from right to left
	for i := len(str) - 1; i >= 0; i-- {
		digit := int(str[i] - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}
