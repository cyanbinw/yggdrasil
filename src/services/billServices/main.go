package main

import (
	"yggdrasil/src/common/router"
	"yggdrasil/src/services/billServices/controller"
)

func main() {
	route := new(controller.BillRoute)
	router.Register(route, "bill")
	router.InitRouter()
}
