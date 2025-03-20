package main

import "api/routes"

func main() {
	router := routes.SetUpRouter()
	router.Run(":8080")
}
