package main

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	for _, item := range items {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	var updatedItem Item
	json.NewDecoder(r.Body).Decode(&updatedItem)
	for i, item := range items {
		if item.Name == name {
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	for i, item := range items {
		if item.Name == name {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}

func main() {
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/item", getItem)
	http.HandleFunc("/create-item", createItem)
	http.HandleFunc("/update-item", updateItem)
	http.HandleFunc("/delete-item", deleteItem)

	http.ListenAndServe(":8080", nil)
}
