package models

import (
	"log"
	"delivery/database/postgres"
	"strconv"
	"fmt"
)


type Item struct {
	ID      int
	Model   string
	Price   int
	Company string
}

func GetAllProducts() []Item {

	var items []Item

	rows, err := postgres.DB.Query("SELECT id, model, price, company FROM products")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close() 

	
	for rows.Next() {
		var item Item 
		err := rows.Scan(&item.ID, &item.Model, &item.Price, &item.Company)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}

		items = append(items, item)
	}


	if err := rows.Err(); err != nil {
		log.Fatal("Error during rows iteration:", err)
	}

	return items
}

func AddProduct(model string, company string, price string) error {
    priceInt, err := strconv.Atoi(price)
    if err != nil {
        return fmt.Errorf("invalid price: %v", err)
    }

    query := "INSERT INTO products (model, company, price) VALUES ($1, $2, $3)"
    _, err = postgres.DB.Exec(query, model, company, priceInt)
    if err != nil {
        return fmt.Errorf("failed to insert product: %v", err)
    }

    return nil
}