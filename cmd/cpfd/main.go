package main

import (
	"cpfd-back/internal/api"
)

func main() {
	router := api.NewRouter()
	router.Run()
}
