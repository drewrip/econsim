package market

import(
	"github.com/dgraph-io/badger"
	"time"
	"fmt"
)

type Market struct{
	ClassName string
	Database *badger.DB
	TickSpeed time.Duration
}

func (m *Market) StartMarket(){
	c := time.NewTicker(m.TickSpeed).C
	for{
		select{
		case <-c:
			fmt.Printf("Processing new events for market %s\n", m.ClassName)
		}
	}
}