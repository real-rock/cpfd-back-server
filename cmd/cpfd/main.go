package main

import (
	"cpfd-back/internal/api"
)

func main() {
	a := NewApp()
	a.PrintAppInfo()
	router := api.NewRouter()
	router.Run()
}
