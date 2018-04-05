package data

import (
	"database/sql"
	"day15/DBSelling/officer/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OfficerRepository struct {
	DB *sql.DB
}

// Untuk inilai return get all butuh struktur dari Officer
// 2.a. Buat model dari Officer
func GetAll(db *OfficerRepository) []models.Officer { // modelOfficer ==> nama file di folder model
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select OfficerCode, OfficerName, OfficerPassword, OfficerStatus FROM tblOfficer")

	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	officer := []models.Officer{}
	for result.Next() {
		var p models.Officer
		if err := result.Scan(&p.OfficerCode, &p.OfficerName, &p.OfficerPassword, &p.OfficerStatus); err != nil {
			fmt.Println("officerCanceError")
			fmt.Println(err)
			return nil
		}
		officer = append(officer, p)
	}
	return officer
}
