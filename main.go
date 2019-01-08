package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func sendResponse(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["a"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'a' is missing")
		return
	}

	a, err := strconv.Atoi(keys[0])

	if err != nil {
		return
	}

	keys, ok = r.URL.Query()["b"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'b' is missing")
		return
	}
	b, err := strconv.Atoi(keys[0])
	if err != nil {
		log.Println("Bad number!")
		return
	}

<<<<<<< HEAD
  fmt.Fprintf(w, "%d + %d = %d", a, b, sum(a,b))
=======
	fmt.Fprintf(w, "%d + %d = %d", a, b, sumValues(a, b))
>>>>>>> bf38ea1 (Formatted code)
}

func main() {
	http.HandleFunc("/", sendResponse)
	log.Fatal(http.ListenAndServe(":7000", nil))
}

<<<<<<< HEAD
func sum (a, b int) int {
  t := b + a
  return t
=======
func sumValues(a, b int) int {
	t := b + a
	return t
>>>>>>> bf38ea1 (Formatted code)
}
