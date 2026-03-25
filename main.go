package main

import (
	"connectdb/db"
	"fmt"
	"net/http"
)

func main() {
	_, errDb := db.ConnectDb()

	if errDb != nil {
		fmt.Println(errDb)
		return
	}

	fmt.Println("Connected successfuly to db")
	fmt.Println("Started server on port 9091")

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Error", err)
	}

}
