package main

import (
	"fmt"
	"net/http"

)

type( 
	Model struct {
	data []string
}
    View struct{}

    Controller struct {
		model Model
		view  View
	}
)
func (m *Model) Create(item string) {
	m.data = append(m.data, item)
}

func (m *Model) Read() []string {
	return m.data
}

func (m *Model) Update(index int, item string) {
	if index >= 0 && index < len(m.data) {
		m.data[index] = item
	}
}

func (m *Model) Delete(index int) {
	if index >= 0 && index < len(m.data) {
		m.data = append(m.data[:index], m.data[index+1:]...)
	}
}

func (v *View) Render(data []string) string {
	return fmt.Sprintf("Data: %v", data)
}

//---

func (c *Controller) Create(item string) {
	c.model.Create(item)
}

func (c *Controller) Read() string {
	data := c.model.Read()
	return c.view.Render(data)
}

func (c *Controller) Update(index int, item string) {
	c.model.Update(index, item)
}

func (c *Controller) Delete(index int) {
	c.model.Delete(index)
}

func main() {
	controller := Controller{}

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		item := r.URL.Query().Get("item")
		controller.Create(item)
		fmt.Fprint(w, "Item created")
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		data := controller.Read()
		fmt.Fprint(w, data)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		index := r.URL.Query().Get("index")
		item := r.URL.Query().Get("item")
		controller.Update(index, item)
		fmt.Fprint(w, "Item updated")
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		index := r.URL.Query().Get("index")
		controller.Delete(index)
		fmt.Fprint(w, "Item deleted")
	})

	http.ListenAndServe(":8080", nil)
}
