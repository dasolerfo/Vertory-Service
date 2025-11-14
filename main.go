package vertoryservice

import (
	"database/sql"
	"log"

	"github.com/dasolerfo/Vertory-Service/api"
	db "github.com/dasolerfo/Vertory-Service/db/model"
	"github.com/dasolerfo/Vertory-Service/factory"
)

func main() {

	config, err := factory.LoadConfig("")
	if err != nil {
		log.Fatal("Error! No pots carregar la configur ", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error! No et pots connectar a la base de dades: ", err)
	}

	store := db.New(testDB)
	router, err := api.NewServer(config, store)

	router.Start("/")

}
