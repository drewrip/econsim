package market

import(
	"github.com/dgraph-io/badger"
)

func (m *Market) AddUser(un string, name string){
	err := m.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(un), []byte(name))
		return err
	})
	check(err)
}