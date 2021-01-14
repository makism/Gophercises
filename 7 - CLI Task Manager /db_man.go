package main

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
)

var Db *bolt.DB

func DbOpenOrCreate(dbFilename string) *bolt.DB {
	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func DbInit() {
	if err := Db.Update(func(tx *bolt.Tx) error {
		fmt.Println("Creating bucket...")
		_, err := tx.CreateBucket([]byte("todolist"))
		return err
	}); err != nil {
		fmt.Println(err)
	}
}

func DbAddRecord(item string) {
	Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todolist"))

		id, _ := b.NextSequence()
		iid := int(id)
		err := b.Put(itob(iid), []byte(item))

		fmt.Printf("Inserting item: %i=%s.\n", iid, item)

		return err
	})
}

func DbListAll() {
	Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todolist"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			id := binary.BigEndian.Uint64(k)
			item := string(v)

			if item == "" {
				item = "\"empty\""
			}

			fmt.Printf("[%d] \t %s\n", id, item)
		}

		return nil
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

