package service

import (
	"strconv"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	model "spotbuzz.com/src/model"

)

var players = []model.Player {
	{ID: 1, Name: "Yash Rastogi", Country: "IN", Score:80},
    {ID: 2, Name: "Jeru", Country: "JP", Score: 82},
    {ID: 3, Name: "Sarah Vaughan ", Country: "US", Score: 72},
	{ID: 4, Name: "Clifford Brown", Country: "CN", Score: 99},
}

func GetPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
	fmt.Println("HALLO")
}

func /*POST*/ AddPlayer(c *gin.Context) {
	var newPlayer model.Player 
	// Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newPlayer); err != nil {
		fmt.Println(err)
        return
    }
	// Add the new album to the slice.
	players = append(players , newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

func GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	ID , _ := strconv.Atoi(id)
    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, player := range players {
        if player.ID == ID {
            c.IndentedJSON(http.StatusOK, player)
            return
        }
    }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "player not found"})
}