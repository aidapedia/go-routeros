package main

import (
	"fmt"
	"log"

	routeros "github.com/aoida/router-os"
	"github.com/aoida/router-os/builder"
	"github.com/aoida/router-os/builder/entity"
)

func main() {
	routerBuilder := routeros.NewRouterOSBuilder("192.168.1.1:8728", "", "")
	err := routerBuilder.Connect()
	if err != nil {
		log.Fatal(err)
	}

	ip := builder.NewIPBuilder(routerBuilder)
	hotspot := builder.NewIPHotspotBuilder(ip)
	active := builder.NewIPHotspotActiveBuilder(hotspot)

	// Example of using Print
	hotspotRes, err := hotspot.Print(entity.PrintRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range hotspotRes {
		fmt.Println(a)
	}

	// Example of using Print
	activeRes, err := active.Print(entity.PrintRequest{
		Where: []entity.Where{
			{
				Field:    "address",
				Value:    "10.10.10.53",
				Operator: entity.OperatorEqual,
			},
		},
	})
	for _, a := range activeRes {
		fmt.Println(a)
	}
	fmt.Println(err)

	routerBuilder.Close()
}
