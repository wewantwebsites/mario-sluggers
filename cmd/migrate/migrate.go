package migrate 

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
    "fmt"

    "sluggers/cmd/models"
    "sluggers/cmd/storage"
)

func migrate() {
    // get db
    db := storage.GetDB()
    // load json file
    jsonFile, err := os.Open("data.json")
    if err != nil {
        log.Fatal(err)
    }
    defer func () {
        if err := jsonFile.Close(); err != nil {
            fmt.Println("couldn't close the json file... not good.")
            log.Fatal(err)
        }
    }()


    byteValue, _ := ioutil.ReadAll(jsonFile)

    var objects []models.Character
    json.Unmarshall(byteValue, &objects)

    // insert JSON data into the tables

    for _, item := range objects {
        // insert stats table first to get the generated id
        res, err := db.Exec("INSERT INTO stats (pitch, bat, field, run) VALUES ($1, $2, $3, $4) RETURNING id",
        item.Stats.Pitch, item.Stats.Bat, item.Stats.Field, item.Stats.Run)
        if err != nil { 
            log.Fatal(err)
        }

        var statsID int 
        err = res.Scan(&statsID)
        if err != nil {
            log.Fatal(err)
        }

        // insert into character table with the stats_id
        _, err = db.Exec("INSERT INTO characters (id, name, description, ability, team, stats_id) VALUES ($1, $2, $3, $4, $5, $6)", item.Name, item.Description, item.Ability, item.Team, statsID)
        if err != nil { 
            log.Fatal(err)
        }

        fmt.Println("db has been migrated")
    }
}
