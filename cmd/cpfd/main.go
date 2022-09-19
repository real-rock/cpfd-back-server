package main

import (
	"cpfd-back/internal/api"
)

func main() {
	setApp()
	router := api.NewRouter()
	router.Run()
}
