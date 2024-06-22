package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() er:%v", err)
		return
	}
	fmt.Fprintf(w, "Post request succesfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s \n", name)
	fmt.Fprintf(w, "Address=%s \n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// "/" - means the root route , send that to the fileserver

	http.HandleFunc("/form", formHandler)
	// at /form , formHandler func(http.responsewriter,*http.request)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server starting at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

/* FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html".
To use the operating system's file system implementation, use [http.Dir]:
http.Handle("/", http.FileServer(http.Dir("/tmp")))
To use an [fs.FS] implementation, use [http.FileServerFS] instead.
*/
