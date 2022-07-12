package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Trade struct {
	UpdateId     int64  `json:"u"`
	Symbol       string `json:"s"`
	BestBidPrice string `json:"b"`
	BestBidQty   string `json:"B"`
	BestAskPrice string `json:"a"`
	BestAskQty   string `json:"A"`
}

type marketOrder struct {
	id           string
	order_id     string
	asset_pair   string
	order_type   string
	price        string
	size         string
	executed_at  int64
	ob_update_id int64
}

func (t marketOrder) addMarketOrder() {
	db := returnDB()

	stmt, _ := db.Prepare("INSERT INTO trades (id, order_id, asset_pair, order_type, price, size, executed_at, ob_update_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(t.id, t.order_id, t.asset_pair, t.order_type, t.price, t.size, t.executed_at, t.ob_update_id)
	defer stmt.Close()

	closeDB(db)

	fmt.Printf("Trade executed %v %v at %v\n", t.asset_pair, t.order_type, t.executed_at)
	fmt.Println("")
}

func execution(c *websocket.Conn, input chan Trade, order order) {
	for trade := range input {
		json, _ := json.Marshal(trade)
		fmt.Println(string(json))
		if order.order_type == "sell" && trade.BestBidPrice >= order.price {
			tsize, _ := strconv.ParseFloat(trade.BestBidQty, 32)
			osize, _ := strconv.ParseFloat(order.size, 32)
			remain, v := countSize(tsize, osize)
			createOrder(trade, order, v).addMarketOrder()
			if remain <= 0.0 {
				break
			}
			order.size = fmt.Sprintf("%f", remain)

		}
		if order.order_type == "buy" && trade.BestAskPrice <= order.price {
			tsize, _ := strconv.ParseFloat(trade.BestAskQty, 32)
			osize, _ := strconv.ParseFloat(order.size, 32)
			remain, v := countSize(tsize, osize)
			createOrder(trade, order, v).addMarketOrder()
			if remain <= 0.0 {
				break
			}
			order.size = fmt.Sprintf("%f", remain)
		}

	}
	closeWSConnection(c)
}

func createOrder(t Trade, o order, s string) marketOrder {

	uniqueId := uuid.New().String()
	now := time.Now()
	mo := marketOrder{
		id:           uniqueId,
		order_id:     o.id,
		asset_pair:   t.Symbol,
		order_type:   o.order_type,
		price:        t.BestBidPrice,
		size:         s,
		executed_at:  now.Unix(),
		ob_update_id: t.UpdateId,
	}
	return mo
}

func countSize(t float64, o float64) (float64, string) {

	remain := 0.0
	v := t - o

	if v <= 0 {
		remain = -v
		v = t
		return remain, fmt.Sprintf("%f", v)
	}

	return remain, fmt.Sprintf("%f", o)

}
