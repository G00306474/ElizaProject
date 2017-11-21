// A simple web application example.

// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13

package main

import (
	"fmt"
	"net/http"
	//"strconv"
	//"math/rand"
	//"time"
	"regexp"
)

//func userinputhandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //Input is passed here from index.html and hello is added to it
//}
/*func userinputhandler(w http.ResponseWriter, r *http.Request) {
	// Get the user input from the request.
	input := r.URL.Query().Get("value")
	// Declare an output string.
	var output string;

	// Try to convert the user input from a string to an unsigned 64-bit integer.
	if uintinput, err := strconv.ParseUint(input, 10, 64); err == nil {
		// If we got an integer, and it's less than 10 then convert it to a word.
		if uintinput < 10 {
			unit := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			output = unit[uintinput]
		} else {
			// Otherwise just return the input.
			output = input
		}
	} else {
		output = input
	}
	fmt.Fprintf(w, "%s", output)
}
*/
func userinputhandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value"))
	input := r.URL.Query().Get("value")
	var output string
	if input == "i need help" {
		output = "you came to the right place "
	}
	if input == "i feel sad" {
		output = "tell me more "
	}
	if input == "i'm tired" {
		output = "you need more sleep "
	}
	


	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		output = "Why don’t you tell me more about your father?"
	}

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		output = re.ReplaceAllString(input, "How do you know you are $1?")
	}

	/*answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	output = answers[rand.Intn(len(answers))]*/


	fmt.Fprintf(w, "%s", output)


}





func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("webFolder"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":3000", nil)
}
