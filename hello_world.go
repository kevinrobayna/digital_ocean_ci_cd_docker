package main
import (
	"net/http"
	"encoding/json"
)
func main() {
	handler := http.NewServeMux()
	///we create a new router to expose our api to our users
	handler.HandleFunc("/api/hello", SayHello)
	//Every time a request is sent to the endpoint ("/api/hello") on port 8080 the function SayHello will be invoked
	http.ListenAndServe("0.0.0.0:8080", handler)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = make(map[string]string)
    response[`message`] = `Hello world`

	json.NewEncoder(w).Encode(response)
}