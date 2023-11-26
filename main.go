package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var err error
var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/go")
	if err != nil {
		panic(err)
	}
	// make router (gin server)
	router := gin.Default()
	// routers
	router.GET("/", getAllProducts)
	router.POST("/", CreateProducts)

	// if router has error
	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

func getAllProducts(c *gin.Context) {
	stmt := "SELECT * FROM products"
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []product
	for rows.Next() {
		var p product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	c.IndentedJSON(http.StatusOK, products)
}

func CreateProducts(c *gin.Context) {
	var prod product
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ins, err := db.Exec("INSERT INTO products(name, price, description) VALUES (?, ?, ?)", prod.Name, prod.Price, prod.Description)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffec, _ := ins.RowsAffected()
	lastInserted, _ := ins.LastInsertId()
	fmt.Println("ID Of Last Row inserted:", lastInserted)
	fmt.Println("Number of rows affected:", rowsAffec)
	c.IndentedJSON(http.StatusCreated, "Data Successfully created!")
}
