package main

import(
	"fmt"
	"github.com/drewrip/econsim/market"
	"github.com/dgraph-io/badger"
	"log"
	"time"
)

func main(){
	deweys := market.Business{BusinessName: "Dewey's Pizza", BusinessType: market.Pizza}
	fmt.Println(deweys)

	testMarket := &market.Market{TickSpeed: 1 * time.Second}
	db, err := badger.Open(badger.DefaultOptions("./data"))
	if err != nil{
		log.Fatalf("[ERR] %v\n", err)
	}

	testMarket.Database = db

	fmt.Println(*testMarket)
	go testMarket.StartMarket()
	for{}
}