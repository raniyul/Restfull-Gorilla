package controllers

import (
	"day15/DBSelling/selling/data"
	"encoding/json"
	"net/http"
)

func GetSelling(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	d := DBInitial(context.DB) // Untuk koneksi
	defer d.Close()

	repo := &data.SellingRepository{d} // memanggil koneksi dari receiver koneksi

	selling := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(selling)
	w.Write(mdl)

}
