package main

import (
	"connectdb/db"
	sqlquery "connectdb/sql_query"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	con, errDb := db.ConnectDb(ctx)

	if errDb != nil {
		fmt.Println(errDb)
		return
	}

	fmt.Println("Connected successfuly to db")

	errTable := sqlquery.CreateTable(ctx, con)

	if errTable != nil {
		fmt.Println("error:", errTable)
		return
	}

	if err := sqlquery.InsertRow(ctx, con, "Подготовиться к собесу", "Яндекс", false, time.Now()); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Created")

	if err := sqlquery.InsertRow(ctx, con, "Подготовиться к мидке", "NLP", false, time.Now()); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Created")
	if err := sqlquery.InsertRow(ctx, con, "Найти работу", "headhunter", false, time.Now()); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Created")
	if err := sqlquery.InsertRow(ctx, con, "Решить по паттернам алгосы", "Pattern", false, time.Now()); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Created")

	// if err := sqlquery.UpdateRow(ctx, con, 1); err != nil {
	// 	fmt.Println("error", err)
	// }
	// fmt.Println("Changed")

	// if err := sqlquery.DeleteTable(ctx, con, 3); err != nil {
	// 	fmt.Println("error", err)
	// }
	// fmt.Println("Deleted")

	tasks, err := sqlquery.SelectRows(ctx, con)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tasks)
}
