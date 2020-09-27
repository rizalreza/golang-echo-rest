package models

import (
	"database/sql"
	"net/http"

	"github.com/rizalreza/golang-echo-rest/database"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func GetAllProduct() (Response, error) {
	var obj Product

	var arrObj []Product

	var res Response

	connection := database.CreateConnection()

	sqlStatement := "SELECT * FROM products"

	rows, err := connection.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Name, &obj.Category, &obj.Price)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func GetProductById(id int) (Response, error) {
	var obj Product
	var res Response

	connection := database.CreateConnection()

	sqlStatement := "SELECT * FROM products WHERE id = ?"

	row := connection.QueryRow(sqlStatement, id)

	switch err := row.Scan(&obj.ID, &obj.Name, &obj.Category, &obj.Price); err {
	case sql.ErrNoRows:
		res.Status = http.StatusNotFound
		res.Message = "Product not found"
	case nil:
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = obj
	default:
		panic(err)
	}

	return res, nil
}

func UpdatePoduct(id, price int, name, category string) (Response, error) {
	var res Response

	connection := database.CreateConnection()

	sqlStatement := "UPDATE products SET name = ?, category = ?, price = ? WHERE id = ?"

	statement, err := connection.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := statement.Exec(name, category, price, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteProduct(id int) (Response, error) {
	var res Response

	connection := database.CreateConnection()

	sqlStatement := "DELETE FROM products WHERE id = ?"

	statement, err := connection.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := statement.Exec(id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected": rowsAffected,
	}

	return res, nil

}
