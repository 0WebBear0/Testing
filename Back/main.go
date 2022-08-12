package main

import (
	"Polytech-lyceum_back/models"
	"Polytech-lyceum_back/router"
)

func main() {
	models.DB()
	router.Router()
}
