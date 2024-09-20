package main

import (
	"FullStackForm/configs"
	"FullStackForm/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}


