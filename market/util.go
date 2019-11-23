package market

import (
	"log"
)

func check(err error){
	if err != nil{
		log.Fatalf("[ERR] %v\n", err)
	}
}