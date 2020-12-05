package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func DbOpenOrCreate(dbFilename string) *bolt.DB {
	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func DbInit(db *bolt.DB, mappings []yamlRecords) {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("mappings"))
		return err
	}); err != nil {
		fmt.Println(err)
	}

	for _, kvp := range mappings {
		//key := kvp.Path[1:]
		key := kvp.Path
		value := kvp.Url

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("mappings"))
			fmt.Printf("Inserting pair: %s=%s.\n", key, value)
			err := b.Put([]byte(key), []byte(value))
			return err
		})
	}
}

func DbFetchForKey(db *bolt.DB, key string) []byte {
	var v []byte = nil

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mappings"))
		v = b.Get([]byte(key))

		return nil // no error
	})

	return v
}
