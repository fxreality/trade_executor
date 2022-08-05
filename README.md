# trade_executor
golang project for swissblock

run project and ,
- first it will create database and the tables
- then you can insert the values like :

Enter asset pair: btcusdt
buy/sell: buy
Enter price: 1980
Enter size: 2

After that the market prices will be written, and if the price match the trades will be stored into db, and reported on console.

--------------------------------------------------------------------------------------------------------------------------------

func splitOrder(order_size int, children int) []int {

	result := make([]int, children)
	remain := order_size

	for i := 0; i < children-1; i++ {
		first := (remain / children) * (children / 2)
		first = rand.Intn(first-(remain/children)) + (remain / children)
		result[i] = first
		remain = remain - first
	}

	result[children-1] = remain

	return result
}
