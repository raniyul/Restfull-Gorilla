package main

import (
	"day15/DBSelling/item/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":1234", router))
}
