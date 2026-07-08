package gobyexample

import (
	"fmt"
	"net/http"
)

func httpServerHelloHandler(w http.ResponseWriter, req *http.Request) {

	// use fmt.Fprintln to write a string to the response writer (nicer for quick output)
	fmt.Fprintln(w, "hello")

	// use w.Write to write a byte slice to the response writer (lower level, but more flexible)
	w.Write([]byte("Yooo!"))

}

// this prints all the headers in the request to the response writer
func httpServerHeadersHandler(w http.ResponseWriter, req *http.Request) {

	for name, values := range req.Header {
		for _, h := range values {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// handler for the /inspect endpoint that prints out the request method, path, and user-agent header to the console, and returns a greeting to the client based on a query parameter
func httpServerInspectHandler(w http.ResponseWriter, r *http.Request) {

	// print the request method, path, and user-agent header to the console
	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("user-agent:", r.Header.Get("User-Agent"))

	// get the value of the "name" query in the URL, e.g. /inspect?name=Reda
	name := r.URL.Query().Get("name")

	// set the response header and status code before writing the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintf(w, "hello %s", name)
}

func ShowHTTPServer() {
	// register the handlers for the two endpoints
	http.HandleFunc("/hello", httpServerHelloHandler)
	http.HandleFunc("/headers", httpServerHeadersHandler)
	http.HandleFunc("/inspect", httpServerInspectHandler)
	// start the server on port 8080
	http.ListenAndServe("127.0.0.1:8080", nil)
}
