package handlers

import (
	"fmt"
	"log"

	"github.com/gowdaganesh005/stocksapi/models"
)

func insertdb(stock models.Stock) int {
	db, err := Connect()

	if err != nil {
		log.Println("Error connecting the database", err)
	}
	defer db.Close()
	var id int
	sqlstatement := "INSERT INTO stocks(id,name,price) values($1,$2,$3) returning id"
	err = db.QueryRow(sqlstatement, stock.ID, stock.Name, stock.Price).Scan(&id)
	if err != nil {
		log.Fatal("error running the sql query:", err)

	}
	fmt.Println("Inserted the row in database successfully")
	return id

}

func getstock(id int64) models.Stock {
	db, err := Connect()

	if err != nil {
		log.Println("Error connecting the database", err)
	}

	defer db.Close()
	var stock models.Stock
	sqlstatement := "SELECT * FROM stocks WHERE ID=$1"
	res := db.QueryRow(sqlstatement, id)
	res.Scan(&stock.ID, &stock.Name, &stock.Price)

	return stock

}
func getallstocks() ([]models.Stock, error) {
	db, err := Connect()

	if err != nil {
		log.Println("Error connecting to database:", err)
	}
	defer db.Close()
	var stocks []models.Stock
	sqlstatment := "SELECT * FROM stocks"
	row, err := db.Query(sqlstatment)
	if err != nil {
		log.Fatalln("Unable to execute the query ", err)
	}
	defer row.Close()
	for row.Next() {
		var stock models.Stock
		err := row.Scan(&stock.ID, &stock.Name, &stock.Price)
		if err != nil {
			log.Println("error retriving data :", err)
		}
		stocks = append(stocks, stock)

	}
	return stocks, err

}
func updatestock(id int64, stock models.Stock) int64 {
	db, err := Connect()

	if err != nil {
		log.Println("Error connecting to database:", err)
	}
	defer db.Close()
	sqlstatment := "Update stocks set name=$2 ,price=$3 WHERE id =$1"
	res, err := db.Exec(sqlstatment, stock.ID, stock.Name, stock.Price)
	if err != nil {
		log.Fatal("Unable to execute the query :", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error fetching how many rows affected:", err)
	}
	return rows

}
func deletestock(id int64) int64 {
	db, err := Connect()

	if err != nil {
		log.Println("Error connecting to database:", err)
	}
	defer db.Close()
	sqlstatement := "DELETE FROM stocks WHERE id=$1"
	res, err := db.Exec(sqlstatement, id)
	if err != nil {
		log.Fatal("error deleting the row :", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error fetching the rows affected:", err)
	}
	return rows

}
