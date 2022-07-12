package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func tickerWS(assetPair string) *websocket.Conn {
	url := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@bookTicker", assetPair)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	checkErr(err)

	return conn
}

func streamData(c *websocket.Conn, input chan Trade) {
	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			var trade Trade
			json.Unmarshal(message, &trade)
			input <- trade
		}
		close(input)
	}()
}

func closeWSConnection(ws *websocket.Conn) {
	defer ws.Close()
}
