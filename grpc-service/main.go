package main

import (
	"fmt"
	"github.com/kytruong0712/goffee-shop/grpc-service/protogen/golang/orders"
	product "github.com/kytruong0712/goffee-shop/grpc-service/protogen/golang/products"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}

	fmt.Println(string(bytes))
}
