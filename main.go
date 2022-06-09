package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello World from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request)  {
    // Check if current request URL path exactly matches "/". I fit doesn't
    // then http.NotFound() function to send a 404 response to the client.
    // Importantly,we then return from the ahndler. If we don't return the handler
    // would keep executing and also write the "Hello from Snippetbox" message
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello from Snippetbox"))
}

// Add snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("create new snippet..."))
}

func main() {
    // Use the http.NewServeMux() function to initialize a new servemux, then
    // regiseter the home function as the handler for the "/" URL pattern
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    // Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen  on(in this case ":4000")
    // and th servemux we just created. If http.ListenAndServe returns a error
    // we use the log.Fatal() function to log the eror message and exit. Note
    // that  any error returned by http.ListenAndServe() is always non-nil.
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}
