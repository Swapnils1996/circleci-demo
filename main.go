package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/phati/circleci-demo/config"
	"github.com/phati/circleci-demo/server"
)

func main() {
	fmt.Println("Listining to port: " + config.ReadEnvString("PORT"))
	log.Fatal(http.ListenAndServe(":"+config.ReadEnvString("PORT"), server.InitRouter()))
}
