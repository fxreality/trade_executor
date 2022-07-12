package main

func main() {

	createTables()

	newOrder := newOrder()
	newOrder.addOrder()

	c := tickerWS(newOrder.asset_pair)
	inputOne := make(chan Trade)

	streamData(c, inputOne)

	execution(c, inputOne, newOrder)

}
