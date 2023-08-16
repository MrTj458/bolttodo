package db

import (
	"encoding/binary"
	"time"

	"go.etcd.io/bbolt"
)

var todosBucket = []byte("todos")

func Open(fileName string) (*bbolt.DB, error) {
	db, err := bbolt.Open(fileName, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(todosBucket)
		return err
	})

	return db, nil
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
