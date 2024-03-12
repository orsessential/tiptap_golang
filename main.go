package main

import "gin-framework/routers"

func main() {
	var PORT = ":9090"

	routers.StartServer().Run(PORT)
}
