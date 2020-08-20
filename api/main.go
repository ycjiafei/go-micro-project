package main

import (
	"github.com/ycjiafei/go-micro-project/api/routes"
)

func main() {
	routes.InitRoutes().Run()
}
