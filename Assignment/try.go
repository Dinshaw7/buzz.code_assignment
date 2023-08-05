package main
import (
	"net/http"
    "github.com/gin-gonic/gin"
	//"strconv"
)

//player represents the information about the player
type player struct {
	ID int 				`json:"id"`
	Name string			`json:"name"`
	Country string		`json:"country"`
	Score int			`json:"score"`
}

//players slice to seed record player data
var players = []player {
	{ID: 1, Name: "Yash Rastogi", Country: "IN", Score:80},
    {ID: 2, Name: "Jeru", Country: "JP", Score: 82},
    {ID: 3, Name: "Sarah Vaughan ", Country: "US", Score: 72},
	{ID: 4, Name: "Clifford Brown", Country: "CN", Score: 99},
}

func main () {
	router := gin.Default()
	router.GET("/players", getPlayers)
	router.GET("/players", getPlayerByID)
	
	//router.POST("/players", postPlayers)
	router.Run("localhost:8080")
}

//getPlayers => gives the information of all players in JSON
func getPlayers(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, players)
}

//postPlayers => creates a new entry for a player 
func postPlayers(c *gin.Context) {

	var newPlayer player
	if err := c.BindJSON(&newPlayer); err != nil {
        return
    }
	//adds newPlayer to the slice
	players = append(players, newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

// getPlayerByID locates the player whose ID value matches the id 
func getPlayerByID(c *gin.Context) {
	var id int

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range players {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "player not found"})
}