package data

import (
	"database/sql"
	"day15/DBSelling/selling/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SellingRepository struct {
	DB *sql.DB
}

// Untuk inilai return get all butuh struktur dari Selling
// 2.a. Buat model dari Selling
func GetAll(db *SellingRepository) []models.Selling { // modelSelling ==> nama file di folder model
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select Invoice, InvoiceDate, Item, Total, Paid, Returned, OfficeCode FROM tblSelling")

	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	selling := []models.Selling{}
	for result.Next() {
		var p models.Selling
		if err := result.Scan(&p.Invoice, &p.InvoiceDate, &p.Item, &p.Total, &p.Paid, &p.Returned, &p.OfficeCode); err != nil {
			fmt.Println("petugasCanceError")
			fmt.Println(err)
			return nil
		}
		selling = append(selling, p)
	}
	return selling
}
