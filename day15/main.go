package main

import (
	"day15/DBSelling/item/routers"
	"log"
	"net/http"
)

func main() {
	// P.1 buat routing
	//buat fungsi routing
	router := routers.InitRouters()

	// buat configurasi server
	log.Fatal(http.ListenAndServe(":8887", router))
}
