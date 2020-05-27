package main

import (
	"fmt"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/configs"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/internal/entities/product"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/internal/routes"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/pkg/repository/adapter"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/pkg/repository/database"
	"log"
	"net/http"
)

func main() {
	configs := config.GetConfig()
	entity := &product.Product{}

	connection := database.GetConnection(configs.Dialect, configs.DatabaseURI)
	connection.Table(entity.TableName()).AutoMigrate(entity)
	repository := adapter.NewAdapter(connection)

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	log.Println("service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}
