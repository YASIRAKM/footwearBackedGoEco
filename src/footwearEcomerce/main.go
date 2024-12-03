package main

import (
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/db"
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/router"
	"log"
	"os"
)

func main() {

	db.Init()
	defer db.CloseDB()
	e := router.New()

	port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT environment variable is not set")
    }

    // Start the server
    e.Logger.Fatal(e.Start(":" + port))

}
