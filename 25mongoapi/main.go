package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nishaa007/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB api")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":27017", r))
	fmt.Println("Listening at port 27017 ....")

}
