package main

import "github.com/slimhappy/web_backend_simple/routers"

func main() {
	err := routers.Run()
	if err != nil {
		panic(err)
	}
}
