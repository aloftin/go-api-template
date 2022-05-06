package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"template-service/api/rest/routers"
	config "template-service/boot"
	"template-service/services"
)

func main() {
	defer log.Println("Template Service ended")
	errChan := make(chan error)

	// listen for ctrl-c
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// Set up services/repos
	services, err := services.NewServices(
		//services.WithDB(config.DBType, config.GetConnectionString()), // TODO
		services.WithHealth(),
	)

	if err != nil {
		log.Fatal("Error wiring up services: ", err)
	}

	defer services.Close()

	// Get the configured port
	configPort, portConfigured := os.LookupEnv("TS_PORT")

	if !portConfigured {
		configPort = config.DefaultPort
	}

	port, err := strconv.Atoi(configPort)

	if err != nil {
		log.Fatalf("Error configuring port %s", configPort)
	}

	// HTTP transport
	go func() {
		errChan <- routers.ListenREST(*services, port)
	}()

	fmt.Printf("Template Service running on port %d\n", port)
	log.Fatal(<-errChan)
}
