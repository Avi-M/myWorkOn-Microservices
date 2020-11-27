package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	ID           string `json:"id,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	OrderedAt    string `json:"orderedAt,omitempty"`
	Items        *Item  `json:"items,omitempty"`
}

type Item struct {
	ItemCode    string `json:"itemCode,omitempty"`
	LineItemID  string `json:"lineItemId,omitempty"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

var orders []Order

func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range orders {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})
}

// Update order
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			var order Order
			_ = json.NewDecoder(r.Body).Decode(&order)
			order.ID = params["id"]
			orders = append(orders, order)
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

func GetordersEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Order
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	orders = append(orders, person)
	json.NewEncoder(w).Encode(orders)
}

func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func main() {
	router := mux.NewRouter()
	orders = append(orders, Order{ID: "1", CustomerName: "Nic", OrderedAt: "2009-11-10 23:00:00", Items: &Item{ItemCode: "D6b89n", LineItemID: "CA", Description: "Work or play, the Redmi 9i is an ideal companion that helps you go through your everyday tasks with ease. ", Quantity: 12}})
	//orders = append(orders, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
	router.HandleFunc("/orders", GetordersEndpoint).Methods("GET")
	router.HandleFunc("/orders/{id}", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/orders/{id}", CreateProductEndpoint).Methods("POST")
	router.HandleFunc("/orders/{id}", DeleteProductEndpoint).Methods("DELETE")
	router.HandleFunc("/order/{id}", updateOrder).Methods("PUT")
	log.Fatal(http.ListenAndServe(":12345", router))
}
