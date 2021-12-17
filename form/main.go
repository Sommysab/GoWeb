package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Product data type for export
type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template

var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error

	db, err := sql.Open("mysql", "root:"+os.Getenv("MYSQL_PASSWORD")+"@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/productsearch", productSearchHandler)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func productSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch.html", nil)
		return
	}
	r.ParseForm()
	var P Product
	name := r.FormValue("productName")
	// fmt.Println("name:", name)

	stmt := "SELECT * FROM products WHERE name = ?;"
	// func (db *DB) QueryRow(query string, args ...interface{}) *Row
	row := db.QueryRow(stmt, name)
	// func (r *Row) Scan(dest ...interface{}) error
	err := row.Scan(&P.ID, &P.Name, &P.Price, &P.Description)
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w, "productsearch.html", P)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
