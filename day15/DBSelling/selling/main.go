package main

import (
	"day15/DBSelling/selling/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":9012", router))
}
