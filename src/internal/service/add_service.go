package service

import (
	"UmairAhmedImran/internal/db"
	"UmairAhmedImran/internal/model"
	"UmairAhmedImran/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

var jsonValue model.BoltDbStruct = model.BoltDbStruct{}

func AddCommand(title string, dbData model.BoltDbStruct) error {

	dbData.Title = title
	db := db.Get()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))

		if b == nil {
			// ask user to create a new bucket or go for health check or init again?
			return fmt.Errorf("bucket %s does not exists", BucketName)
		}

		value := b.Get([]byte(title))

		if value == nil {
			bytesData, err := json.Marshal(dbData)

			if err != nil {
				return err
			}

			utils.ShowingInfo.Printf("Marshal data: %s", bytesData)

			if err := b.Put([]byte(title), bytesData); err != nil {
				return err
			}
		} else {
			err := json.Unmarshal(value, &jsonValue)
			if err != nil {
				return err
			}

			jsonValue.Notes = append(jsonValue.Notes, dbData.Notes...)

			data, err := json.Marshal(jsonValue)
			if err != nil {
				return err
			}
			err = b.Put([]byte(title), data)
			return err
		}
		return nil
	})

}
