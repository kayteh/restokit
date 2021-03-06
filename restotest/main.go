package main

import (
	"fmt"

	"github.com/kayteh/restokit"
	"github.com/kayteh/restokit/restotest/api"
)

func main() {
	resto := restokit.NewRestokit("127.0.0.1:4665")
	api.FetchAPIRoutes(resto.Router)
	fmt.Println("started 127.0.0.1:4665")
	resto.Start()
}
