package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type RouteComment struct {
	Route   string
	Comment string
}

type RunData struct {
	WelcomeMessage string
	RouteComments  []RouteComment
}

var welcomeMessage string
var routecomments []RouteComment

func init() {
	welcomeMessage = "Herbalife server on the go..."
	routecomments = []RouteComment{}
}

func GetRunData() RunData {
	return RunData{
		WelcomeMessage: welcomeMessage,
		RouteComments:  routecomments,
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("index_tmpl.html")
	if err != nil {
		panic(err)
	}
	data := GetRunData()
	t.Execute(writer, data)
}

func allProducts(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("all_tmpl.html")
	if err != nil {
		panic(err)
	}

	products, err := GetAllProducts()
	if err != nil {
		panic(err)
	}

	t.Execute(writer, products)
}

func search(writer http.ResponseWriter, request *http.Request) {

	t, err := template.ParseFiles("all_tmpl.html")
	if err != nil {
		panic(err)
	}

	request.ParseForm()
	val := request.Form["search"][0]
	products, err := GetAllProductWith(val)
	if err != nil {
		panic(err)
	}

	t.Execute(writer, products)
}

func setupServerAndRun() {
	fmt.Println(welcomeMessage)

	port := ":8085"
	server := http.Server{
		Addr: os.Getenv("PORT"),
	}
	fmt.Println("Serving at port ", port)

	c := cors.AllowAll()

	// http.Handle("/", c.Handler(http.HandlerFunc(index)))
	routecomment := RouteComment{
		Route:   "/",
		Comment: "Default route",
	}
	routecomments = append(routecomments, routecomment)

	http.Handle("/products", c.Handler(http.HandlerFunc(allProducts)))
	routecomment = RouteComment{
		Route:   "/products",
		Comment: "returns a list of all products in a presentation.",
	}
	routecomments = append(routecomments, routecomment)

	http.Handle("/search", c.Handler(http.HandlerFunc(search)))
	routecomment = RouteComment{
		Route:   "/search",
		Comment: "returns a list of all products in a presentation after search.",
	}
	routecomments = append(routecomments, routecomment)

	css := http.FileServer(http.Dir("./"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	routecomment = RouteComment{
		Route:   "/css/",
		Comment: "make sure the css file are served",
	}
	routecomments = append(routecomments, routecomment)

	server.ListenAndServe()
}

func checkThis() {
	res, _ := GetAllProductWith("shake")
	fmt.Println(res)
}

func main() {
	// checkThis()
	setupServerAndRun()
}
