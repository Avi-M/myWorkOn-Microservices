package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Order struct {
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

type Item struct {
	LineItemID  uint   `json:"lineItemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

var db *gorm.DB

var (
	od = []Order{
		{OrderID: 123, CustomerName: "Jimmy Johnson", OrderedAt: time.Now()},
		{OrderID: 889, CustomerName: "Howard Hills", OrderedAt: time.Now()},
		{OrderID: 678, CustomerName: "Craig Colbin", OrderedAt: time.Now()},
	}
	it = []Item{
		{LineItemID: 2000, ItemCode: "Toyota", Description: "Tundra", Quantity: 1, OrderID: 123},
		{LineItemID: 2001, ItemCode: "Honda", Description: "Accord", Quantity: 1, OrderID: 889},
		{LineItemID: 2002, ItemCode: "Nissan", Description: "Sentra", Quantity: 2, OrderID: 678},
		{LineItemID: 2003, ItemCode: "Ford", Description: "F-150", Quantity: 3, OrderID: 123},
	}
)

func initDB() {

	var err error
	dataSourceName := "root:password@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.Exec("CREATE DATABASE ordersdetail")
	db.Exec("USE ordersdetail")
	db.AutoMigrate(&Order{}, &Item{})
}

func main() {
	initDB()

	for index := range od {
		db.Create(&od[index])
	}

	for index := range it {
		db.Create(&it[index])
	}
	router := mux.NewRouter()
	router.HandleFunc("/postorders", createOrder).Methods("POST")
	router.HandleFunc("/getorders/{orderId}", getOrder).Methods("GET")
	router.HandleFunc("/getallorders", getOrders).Methods("GET")
	router.HandleFunc("/updateorders/{orderId}", updateOrder).Methods("PUT")
	router.HandleFunc("/deleteorders/{orderId}", deleteOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []Order
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]

	var order Order
	db.Preload("Items").First(&order, inputOrderID)
	json.NewEncoder(w).Encode(order)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Save(&updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedOrder)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
	idToDelete := uint(id64)

	db.Where("order_id = ?", idToDelete).Delete(&Item{})
	db.Where("order_id = ?", idToDelete).Delete(&Order{})
	w.WriteHeader(http.StatusNoContent)
}
