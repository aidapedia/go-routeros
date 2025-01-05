package main

import (
	"context"
	"fmt"
	"log"

	routeros "github.com/aidapedia/airouteros"
	ipBuilder "github.com/aidapedia/airouteros/builder/ip"
	"github.com/aidapedia/airouteros/model"
)

func main() {
	routerBuilder := routeros.NewRouterOS(&routeros.Options{})
	err := routerBuilder.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer routerBuilder.Close()

	ip := ipBuilder.NewIPBuilder(routerBuilder)
	hotspot := ipBuilder.NewIPHotspotBuilder(ip)
	active := ipBuilder.NewIPHotspotUserProfileBuilder(hotspot)

	// Example of using Print
	hotspotRes, err := hotspot.Print(context.Background(), model.PrintRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range hotspotRes {
		fmt.Println(a)
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
