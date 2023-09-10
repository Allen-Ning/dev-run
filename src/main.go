package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Allen-Ning/dev-run/config"
	"github.com/Allen-Ning/dev-run/repositories"
	"github.com/Allen-Ning/dev-run/runtime"
	"github.com/Allen-Ning/dev-run/services"
)

const (
	downloadDir = "./downloads"
	configFile  = "repos.yaml"
)

func main() {
	action := flag.String("action", "clone", "Action to perform: 'clone', 'docker-up', 'list-services', or 'run-service'")
	targetService := flag.String("service", "", "Target service to run")
	flag.Parse()

	config, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
	}

	switch *action {
	case "clone":
		repositories.CloneRepositories(config, downloadDir)
	case "docker-up":
		runtime.RunDockerCompose(config, downloadDir)
	case "list-services":
		services, err := services.ListServices(config, downloadDir)
		if err != nil {
			log.Fatalf("Failed to list services: %v\n", err)
		}
		fmt.Printf("Services:\n")
		for _, service := range services {
			fmt.Printf("- Repo: %s, Service: %s\n", service.Repo, service.Service)
		}
	case "run-service":
		if *targetService == "" {
			log.Fatalf("Please provide a service name using the -service flag\n")
		}
		err = services.RunTargetService(config, downloadDir, *targetService)
		if err != nil {
			log.Fatalf("Failed to run target service: %v\n", err)
		}
	default:
		log.Fatalf("Invalid action: %s. Choose 'clone', 'docker-up', 'list-services', or 'run-service'.\n", *action)
	}
}
