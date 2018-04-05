package data

import (
	"database/sql"
	"day15/DBSelling/item/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ItemRepository struct {
	DB *sql.DB
}

// Untuk inilai return get all butuh struktur dari Item
// 2.a. Buat model dari Item
func GetAll(db *ItemRepository) []models.Item { // modelItem ==> nama file di folder model
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select ItemCode, ItemName, BuyingPrice, SellingPrice, ItemAmount, Pieces FROM tblItem")

	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	item := []models.Item{}
	for result.Next() {
		var p models.Item
		if err := result.Scan(&p.ItemCode, &p.ItemName, &p.BuyingPrice, &p.SellingPrice, &p.ItemAmount, &p.Pieces); err != nil {
			fmt.Println("petugasCanceError")
			fmt.Println(err)
			return nil
		}
		item = append(item, p)
	}
	return item
}
