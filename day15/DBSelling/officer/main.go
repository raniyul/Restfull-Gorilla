package main

import (
	"day15/DBSelling/officer/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":5678", router))
}
