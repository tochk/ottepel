package main

import  (
	"log"
	"net/http"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {


	http.HandleFunc("/auth/", authHandler)
	log.Println("Server started at port :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}