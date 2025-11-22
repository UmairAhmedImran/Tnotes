package service

import (
	"UmairAhmedImran/internal/db"
	"UmairAhmedImran/internal/utils"
	"fmt"

	"github.com/boltdb/bolt"
)

func ViewCommand(title ...string) error {

	CheckInit()

	db := db.Get()

	if len(title) == 0 {

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BucketName))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("key=%s, value=%s\n", k, v)
			}
			return nil
		})

	} else if len(title) == 1 {

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BucketName))
			value := b.Get([]byte(title[0]))
			if value == nil {
				utils.ShowingWarning.Println("There is no such title")
			}
			return nil
		})
	} else {
		utils.ShowingError.Println("Error too many arguments")
	}
	return nil
}
