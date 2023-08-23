package slessdb

import (
	dbm "github.com/tendermint/tm-db"
)

type DebugDB struct {
	dbm.MemDB
}

var _ dbm.DB = (*DebugDB)(nil)

func NewDebugDB() *DebugDB {
	database := &DebugDB{
		*dbm.NewMemDB(),
	}
	return database
}

func (db *DebugDB) Get(key []byte) ([]byte, error) {
	// fmt.Printf("Get")
	return db.MemDB.Get(key)
}

func (db *DebugDB) Has(key []byte) (bool, error) {
	// fmt.Println("Has")
	return db.MemDB.Has(key)
}

func (db *DebugDB) Set(key, value []byte) error {
	// fmt.Println("Set")
	return db.MemDB.Set(key, value)
}

func (db *DebugDB) SetSync(key, value []byte) error {
	// fmt.Println("SetSync")
	return db.MemDB.SetSync(key, value)
}

func (db *DebugDB) Delete(key []byte) error {
	// fmt.Println("Delete")
	return db.MemDB.Delete(key)
}

func (db *DebugDB) DeleteSync(key []byte) error {
	// fmt.Println("DeleteSync")
	return db.MemDB.DeleteSync(key)
}

func (db *DebugDB) Iterator(start, end []byte) (dbm.Iterator, error) {
	// fmt.Printf("Iterator %s %s\n", string(start), string(end))
	return db.MemDB.Iterator(start, end)
}

func (db *DebugDB) ReverseIterator(start, end []byte) (dbm.Iterator, error) {
	// fmt.Printf("ReverseIterator %s %s\n", string(start), string(end))
	return db.MemDB.ReverseIterator(start, end)
}

func (db *DebugDB) Close() error {
	// fmt.Println("Close")
	return db.MemDB.Close()
}

func (db *DebugDB) NewBatch() dbm.Batch {
	// fmt.Println("NewBatch")
	return db.MemDB.NewBatch()
}

func (db *DebugDB) Print() error {
	// fmt.Println("Print")
	return db.MemDB.Print()
}

func (db *DebugDB) Stats() map[string]string {
	// fmt.Println("Stats")
	return db.MemDB.Stats()
}
