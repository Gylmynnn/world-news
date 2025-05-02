package main

import (
	"github.com/Gylmynnn/world-news/router"
)

func main() {
	err := router.InitRouter()
	if err != nil {
		panic("failed to start server")
	}
}

