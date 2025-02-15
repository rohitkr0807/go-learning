package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	// sum()
	// httpRequestServe()
	result := multiply(10, 20)
	fmt.Println("Result of multiplication:", result)
}

func httpRequestServe() {
	fmt.Println("Hello, World Http!")
	mux := http.NewServeMux()

	// By default, the ServeMux returns a 404 Not Found error for all requests.

	// this handle function is used to handle the request, by defeault it is GET method
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	mux.HandleFunc("GET /comments", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println(r.Method)
		// fmt.Print(r.RequestURI)
		fmt.Fprintln(w, "GET Comments")
		// fmt.Println(w, "hello")
	})

	mux.HandleFunc(("GET /comments/{id}"), func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		id2 := r.PathValue("id")
		if id == "" {
			id = r.PathValue("id")
		}
		fmt.Println("id 1 is" + id)
		fmt.Println("id2 is " + id2)
		fmt.Println("id3 is " + r.PathValue("id"))
		fmt.Print(r.Host)
		fmt.Fprintln(w, "GET Comments =", id2)
		fmt.Fprintln(w, "GET Comments =", id)
	})

	mux.HandleFunc("POST /comments", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " POST Comments")
	})

	mux.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET Comments 2")
	})

	http.ListenAndServe("localhost:8080", mux)
}

// func main() {
// 	fmt.Println("Rohit")
// 	fmt.Println("Rohit")
// 	mains()
// 	sum()
// 	test()
// }

func sum() {
	var a = 10
	b := 20
	c := "rohit"
	d := " kumar rohit"
	result := a + b
	result2 := c + d
	fmt.Println("Sum is: ", result)
	fmt.Println("Sum is: ", result2)
	sumtotal := add(10, 20)
	fmt.Println("Sum of 10 and 20 is: ", sumtotal)
	// add two number
}

func add(a int, b int) int {
	var c = 10
	fmt.Println("C is: ", c)
	return a + b
}
