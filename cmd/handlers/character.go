package handlers

import (
	"log"
	"net/http"
	"sluggers/cmd/models"
	"sluggers/cmd/storage"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	characters = make(map[int]*models.Character)
	lock       = sync.Mutex{}
	hydrated   = false
)

// GetAllCharacters godoc
// @Summary     Get all mario sluggers characters
// @Description get all characters
// @Tags        Characters
// @Accept	    json
// @Produce     json
// @Success     200 {array}     models.Character
// @Router      /all [get]
func GetAllCharacters(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	if hydrated {
		return c.JSON(http.StatusOK, characters)
	}

	db := storage.GetDB()
	rows, err := db.Query("SELECT characters.ID, Name, Description, Ability, Team, Bat, Pitch, Field, Run FROM characters INNER JOIN stats ON characters.ID = stats.CharacterID")
	if err != nil {
		c.Logger().Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var ch = models.Character{}
		if err := rows.Scan(
			&ch.ID,
			&ch.Name,
			&ch.Description,
			&ch.Ability,
			&ch.Team,
			&ch.Stats.Bat,
			&ch.Stats.Field,
			&ch.Stats.Pitch,
			&ch.Stats.Run); err != nil {
			log.Fatal(err)
		}

		if characters[ch.ID] == nil {
			characters[ch.ID] = &ch
		}
	}

	hydrated = true
	return c.JSON(http.StatusOK, characters)
}

// GetCharacter godoc
// @Summary     Get mario sluggers characters by their id
// @Description Returns a Mario Sluggers character or an error
// @Tags       	Characters
// @Produce     json
// @Success     200 	{object}   models.Character
// @Router      /:id		[get]
func GetCharacter(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	if characters[id] != nil {
		return c.JSON(http.StatusOK, characters[id])
	}

	db := storage.GetDB()
	rows, err := db.Query(
		"SELECT Name, Description, Ability, Team, Bat, Pitch, Field, Run FROM characters INNER JOIN stats ON characters.ID = stats.CharacterID WHERE characters.ID = ?", id)
	if err != nil {
		c.Logger().Error(err)
	}
	defer rows.Close()

	var ch = models.Character{ID: id} // we got the record using this id
	for rows.Next() {                 // highlander
		if err := rows.Scan(
			&ch.Name,
			&ch.Description,
			&ch.Ability,
			&ch.Team,
			&ch.Stats.Bat,
			&ch.Stats.Pitch,
			&ch.Stats.Field,
			&ch.Stats.Run); err != nil {
			log.Fatal(err)
		}

		if characters[id] == nil {
			characters[id] = &ch
		}
	}
	var status = http.StatusOK
	if ch.Name == "" {
		status = http.StatusNotFound
	}

	return c.JSON(status, &ch)
}
