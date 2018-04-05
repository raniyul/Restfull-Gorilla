package controllers

import (
	"day15/DBSelling/officer/data"
	"encoding/json"
	"net/http"
)

func GetOfficer(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	d := DBInitial(context.DB) // Untuk koneksi
	defer d.Close()

	repo := &data.OfficerRepository{d} // memanggil koneksi dari receiver koneksi

	officer := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(officer)
	w.Write(mdl)

}
