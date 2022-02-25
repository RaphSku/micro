package main

import (
	"flag"
	"fmt"

	"github.com/RaphSku/micro/internal/server"
)

func main() {
	databaseNamePtr := flag.String("database", "", "A string which specifies the database name")
	portPtr := flag.String("port", ":8090", "A string which specifies the port on which the server will run")
	flag.Parse()

	fmt.Printf("Server will start running on port %s and database %s!", *portPtr, *databaseNamePtr)
	server.CreateServer(*portPtr, *databaseNamePtr)
}
