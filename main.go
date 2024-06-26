package main

import (
	"fmt"
	"go-hexagonal-example/core"
	"go-hexagonal-example/db"
	"go-hexagonal-example/utils"
	"log"

)

func main() {
	// _ = pagi.Paginate[interface{}]()
	// paging
	fmt.Println(goease.ToLowerCase("Hello"))

	param := utils.ProvideFiberHttpServiceParams()
	fiberSrv := utils.InitializeHTTPService(param)

	database, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}
	db.DB = database

	repo := core.NewRepository(db.DB)
	srv := core.NewServices(repo)
	handler := core.NewHandlers(srv)

	fiberSrv.Get("/", handler.HandleGetTodoes)

	portString := fmt.Sprintf(":%v", param.Port)
	err = fiberSrv.Listen(portString)

	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}

}
