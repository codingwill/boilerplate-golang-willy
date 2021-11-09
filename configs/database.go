package configs

import (
	"fmt"
	"log"
	"projectstructuring/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

var DB *sqlx.DB

func InitDB(conf models.Config) {
	var err error
	//postgresql
	confStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s",
		conf.Username,
		conf.Password,
		conf.Database,
		conf.Port,
	)
	DB, err = sqlx.Connect("postgres", confStr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("DB init completed. \nDB in Config: ", DB)

}
