package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	name  string
	color string
	size  string
}

type Payment_card struct {
	Id           uint64  `json:"id"`
	Item         Product `json:"item"`
	UnitPrice    uint64  `json:"price"`
	Quaintity    uint8   `json:"quantity"`
	ShippingFree bool    `json:"shipping"`
	SubTotal     uint64  `json:"total"`
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/checkout.html")
	if err != nil {
		http.NotFound(w, r)
	}
	tmpl.Execute(w, nil)
	// tmpl.ExecuteTemplate(w, "checkout", nil)
}

func singlePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/singlePage.html")
	if err != nil {
		http.NotFound(w, r)
	}
	tmpl.Execute(w, nil)
	// tmpl.ExecuteTemplate(w, "singlePage", nil)
}

func shopCartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/shoppingCart.html")
	if err != nil {
		http.NotFound(w, r)
	}
	tmpl.Execute(w, nil)
	// tmpl.ExecuteTemplate(w, "shoppingCart", nil)

}

func productHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/product.html")
	if err != nil {
		http.NotFound(w, r)
	}
	tmpl.Execute(w, nil)
	// tmpl.ExecuteTemplate(w, "product", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/index.html")
	if err != nil {
		http.NotFound(w, r)
	}
	tmpl.Execute(w, nil)
	// tmpl.ExecuteTemplate(w, "index", nil)
}

func handlerResponse() {
	port := ":9090"
	fmt.Println("Listen on port ", port)

	//------------------HandleHTML------------------
	http.HandleFunc("index", indexHandler)
	http.HandleFunc("product", productHandler)
	http.HandleFunc("shoppingCart", shopCartHandler)
	http.HandleFunc("singlePage", singlePageHandler)
	http.HandleFunc("checkout", checkoutHandler)

	//------------------HandleCSS------------------
	http.Handle("/source", http.StripPrefix("/source", http.FileServer(http.Dir("./source/"))))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}

}

func main() {
	handlerResponse()
}
