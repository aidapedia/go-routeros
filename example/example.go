package main

import (
	"context"
	"fmt"
	"log"

	routeros "github.com/aidapedia/go-routeros"
	"github.com/aidapedia/go-routeros/driver"
	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/module"
)

func main() {
	routerBuilder := routeros.NewRouterOS(&routeros.Options{
		Address:  "192.168.1.3:8728",
		Username: "",
		Password: "",
	})
	err := routerBuilder.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer routerBuilder.Close()

	active, err := driver.New(routerBuilder, module.HotspotModule)
	if err != nil {
		log.Fatal(err)
	}

	// Example of using Print
	activeRes, err := active.Print(context.Background(), model.PrintRequest{
		Where: []model.Where{
			// {
			// 	Field:    "address",
			// 	Value:    "10.10.10.53",
			// 	Operator: model.OperatorEqual,
			// },
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, a := range activeRes {
		fmt.Println(a)
	}
}
