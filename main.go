package main

import (
	"os"
	"sellingGorilla/controller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initial := controller.App{}
	initial.Initializer(
		os.Getenv("root"),
		os.Getenv("root"),
		os.Getenv("DBSelling"),
	)
	initial.Run(":8898")
}
