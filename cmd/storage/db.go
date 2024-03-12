package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"sluggers/cmd/models"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("libsql://%s-wewantwebsites.turso.io?authToken=%s", dbName, os.Getenv("DB_AUTH_TOKEN"))

	db, err = sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		log.Fatal("failed to connect to the db")
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
	fmt.Println("connected to the db@", dbName)

	return db
}

type Data struct {
	Characters []models.Character `json:"characters"`
}

func Migrate() {
	fmt.Println("begin data.json migration")
	// load json file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Println("couldn't close the json file... not good.")
			log.Fatal(err)
		}
	}()

	byteValue, _ := io.ReadAll(jsonFile)
	fmt.Println("reading the file")

	var objects Data
	err = json.Unmarshal(byteValue, &objects)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to db established")
	// insert JSON data into the tables
	for _, item := range objects.Characters {
		// insert into character table
		res, err := db.Exec("INSERT INTO characters (name, description, ability, team) VALUES (?, ?, ?, ?) RETURNING id", item.Name, item.Description, item.Ability, item.Team)
		if err != nil {
			log.Fatal(err)
		}

		characterId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		// insert stats table
		_, err = db.Exec(fmt.Sprintf("INSERT INTO stats (characterID, pitch, bat, field, run) VALUES (%d, %d, %d, %d, %d)", characterId, item.Stats.Pitch, item.Stats.Bat, item.Stats.Field, item.Stats.Run))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s has been migrated", item.Name)
	}
}
