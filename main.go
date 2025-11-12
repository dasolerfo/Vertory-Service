package vertoryservice

import (
	"database/sql"
	"log"

	db "github.com/dasolerfo/Vertory-Service/db/model"
)

func main() {
	testDB, err := sql.Open("config.DBDriver", "config.DBSource")
	if err != nil {
		log.Fatal("Error! No et pots connectar a la base de dades: ", err)
	}

	store := db.New(testDB)
	router, err := api.NewServer()

}
