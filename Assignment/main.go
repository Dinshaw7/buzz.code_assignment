package main

import (
	"fmt"
	//log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	service "spotbuzz.com/service"
)

func main() {
	fmt.Println("Hello. This is BitGo Assessment !!")
	router := gin.Default()
    router.GET("/players", service.GetPlayers)
	router.POST("/players", service.AddPlayer)
	router.GET("/players/:id", service.GetPlayerByID)
    router.Run("localhost:8080")
	// txn_ancestry, err := service.GetAncestor(blockNumber)
	// if err != nil {
	// 	fmt.Printf("GetAncestor failed: %v", err)
	// 	return
	// }
}