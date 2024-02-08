package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gowdaganesh005/stocksapi/models"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type response struct {
	ID      int    `json:"id"`
	Message string `json:"msg"`
}

func Connect() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env files:", err)
	}
	dbstring := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal("error connecting the database :", err)
	}
	return db, nil

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Println("Error parsing the request:", err)
	}

	defer r.Body.Close()

	insertid := insertdb(stock)
	res := response{
		ID:      insertid,
		Message: "Stock inserted successfully",
	}
	json.NewEncoder(w).Encode(res)

}
func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Error parsing  the request id  :", err)

	}
	stock := getstock(int64(id))

	json.NewEncoder(w).Encode(stock)

}
func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getallstocks()
	if err != nil {
		log.Println("error in fetching the stocks :", err)
	}
	json.NewEncoder(w).Encode(stocks)

}
func Updatestock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Error parsing the request  id :", err)

	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Println("Error parsing the request:", err)
	}

	defer r.Body.Close()
	updaterows := updatestock(int64(id), stock)
	res := response{
		ID:      id,
		Message: fmt.Sprintf("Stock updated successfully total rows affected is %v", updaterows),
	}
	json.NewEncoder(w).Encode(res)

}
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("error parsing request id:", err)

	}
	deleterows := deletestock(int64(id))
	res := response{
		ID:      id,
		Message: fmt.Sprint("Deleted the row successfully total row deleted:", deleterows),
	}
	json.NewEncoder(w).Encode(res)

}
