package src

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type Link struct {
	Id      string
	Message string
}

type Tmpl struct {
	Links    []Link
	Products []Product
}

var allProductsLinks []Link

func init() {
	allProductsLinks = []Link{
		// Link{
		// 	Id:      "idtop",
		// 	Message: "Top | ",
		// },
		Link{
			Id:      "id1",
			Message: "Herbalife Products",
		},
	}
}

func writeHtmlProductResponse(products []Product, writer http.ResponseWriter) {

	tmplFiles := []string{
		"./public/html/layout.html",
		"./public/html/layoutcontent.html",
		"./public/html/headerpage.html",
		"./public/html/links.html",
		"./public/html/page.html",
		"./public/html/searcharea.html",
		"./public/html/primarypagecontent.html",
	}

	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		panic(err)
	}

	data := Tmpl{
		Links:    allProductsLinks,
		Products: products,
	}

	t.ExecuteTemplate(writer, "layout", data)
}

func allProducts(writer http.ResponseWriter, request *http.Request) {

	products, err := GetAllProducts()
	if err != nil {
		panic(err)
	}

	writeHtmlProductResponse(products, writer)
}

func search(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	val := request.Form["search"][0]
	products, err := GetAllProductWith(val)
	if err != nil {
		panic(err)
	}

	writeHtmlProductResponse(products, writer)
}

func Run() {
	fmt.Println(welcomeMessage)

	port := ":8085"
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
		// Addr: port,
	}
	fmt.Println("Serving at port ", port)

	c := cors.AllowAll()

	public := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", public))
	http.Handle("/", c.Handler(http.HandlerFunc(allProducts)))
	http.Handle("/search", c.Handler(http.HandlerFunc(search)))

	server.ListenAndServe()
}
