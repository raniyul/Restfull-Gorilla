package controllers

import (
	"day15/DBSelling/item/data"
	"encoding/json"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	d := DBInitial(context.DB) // Untuk koneksi
	defer d.Close()

	repo := &data.ItemRepository{d} // memanggil koneksi dari receiver koneksi

	item := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)

}
