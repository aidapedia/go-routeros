package main

import (
	"fmt"
	"log"

	routeros "github.com/aidapedia/airouteros"
	ipBuilder "github.com/aidapedia/airouteros/builder/ip"
	"github.com/aidapedia/airouteros/model"
)

func main() {
	routerBuilder := routeros.NewRouterOS("127.0.0.1:8728", "", "")
	err := routerBuilder.Connect()
	if err != nil {
		log.Fatal(err)
	}

	ip := ipBuilder.NewIPBuilder(routerBuilder)
	hotspot := ipBuilder.NewIPHotspotBuilder(ip)
	active := ipBuilder.NewIPHotspotActiveBuilder(hotspot)

	// Example of using Print
	hotspotRes, err := hotspot.Print(model.PrintRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range hotspotRes {
		fmt.Println(a)
	}

	// Example of using Print
	activeRes, err := active.Print(model.PrintRequest{
		Where: []model.Where{
			{
				Field:    "address",
				Value:    "10.10.10.53",
				Operator: model.OperatorEqual,
			},
		},
	})
	for _, a := range activeRes {
		fmt.Println(a)
	}
	fmt.Println(err)

	routerBuilder.Close()
}
