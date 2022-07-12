package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type order struct {
	id         string
	asset_pair string
	order_type string
	price      string
	size       string
	created_at int64
}

func newOrder() order {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter asset pair: ")
	assetPair, _ := reader.ReadString('\n')
	if assetPair != "\n" {
		assetPair = strings.TrimSuffix(assetPair, "\n")
	}
	fmt.Print("buy/sell: ")
	orderType, _ := reader.ReadString('\n')
	if orderType != "\n" {
		orderType = strings.TrimSuffix(orderType, "\n")
	}
	fmt.Print("Enter price: ")
	price, _ := reader.ReadString('\n')
	if price != "\n" {
		price = strings.TrimSuffix(price, "\n")
	}
	fmt.Print("Enter size: ")
	size, _ := reader.ReadString('\n')
	if size != "\n" {
		size = strings.TrimSuffix(size, "\n")
	}

	uniqueId := uuid.New().String()
	now := time.Now()

	order := order{
		id:         uniqueId,
		asset_pair: assetPair,
		order_type: orderType,
		price:      price,
		size:       size,
		created_at: now.Unix(),
	}

	return order
}

func (o order) addOrder() {

	db := returnDB()

	stmt, _ := db.Prepare("INSERT INTO orders (id, asset_pair, order_type, price, size, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	stmt.Exec(o.id, o.asset_pair, o.order_type, o.price, o.size, o.created_at)
	defer stmt.Close()

	closeDB(db)

	fmt.Printf("Order saved %v %v \n", o.asset_pair, o.order_type)
}
