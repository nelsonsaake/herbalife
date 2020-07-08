package src

import (
	"fmt"
	"html/template"
	"net/http"

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

func productResponseWriter(products []Product, writer http.ResponseWriter) {

	tmplFiles := []string{
		"./public/html/default/layout.html",
		"./public/html/default/layoutcontent.html",
		"./public/html/default/headerpage.html",
		"./public/html/default/links.html",
		"./public/html/default/page.html",
		"./public/html/default/searcharea.html",
		"./public/html/default/primarypagecontent.html",
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

	productResponseWriter(products, writer)
}

func search(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	val := request.Form["search"][0]
	products, err := GetAllProductWith(val)
	if err != nil {
		panic(err)
	}

	productResponseWriter(products, writer)
}

func solutionResponseWriter(sln Solution, writer http.ResponseWriter) {

	tmplFiles := []string{
		"./public/html/default/layout.html",
		"./public/html/sln/links.html",
		"./public/html/sln/layoutcontent.html",
		"./public/html/sln/package.html",
		"./public/html/sln/packageimgview.html",
	}

	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "layout", sln)
}

func weigthGain(writer http.ResponseWriter, request *http.Request) {

	sln := weightGainSolution()
	solutionResponseWriter(sln, writer)
}

func Run() {
	fmt.Println(welcomeMessage)

	port := ":8085"
	server := http.Server{
		// Addr: ":" + os.Getenv("PORT"),
		Addr: port,
	}
	fmt.Println("Serving at port ", port)

	c := cors.AllowAll()

	public := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", public))
	http.Handle("/", c.Handler(http.HandlerFunc(allProducts)))
	http.Handle("/search", c.Handler(http.HandlerFunc(search)))
	http.Handle("/wg", c.Handler(http.HandlerFunc(weigthGain)))

	server.ListenAndServe()
}
