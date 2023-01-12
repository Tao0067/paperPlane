package main

import "go_xx/router"

func main()  {
	r := router.Router()
	r.Run(":8090")
}
