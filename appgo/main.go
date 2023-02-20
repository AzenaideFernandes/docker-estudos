package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
func main() {
	println("Hello world!")
	insertproducts()
}

//insert products using sql.db
//Path: appgo/main.go
func insertproducts() {
	//db, err := sql.Open("mysql","root:root@tcp(localhost:3307)/test")
	db, err := sql.Open("mysql","root:root@tcp(mysql:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//insert products
	_, err = db.Exec("INSERT INTO products (id, name) VALUES ('1', 'Carrinho')")
	if err != nil {
		panic(err.Error())
	}
}