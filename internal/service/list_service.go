package service

import (
	"UmairAhmedImran/internal/db"
	"UmairAhmedImran/internal/model"
	"UmairAhmedImran/internal/utils"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
)

func ListCommand(recursively bool) error {

	db := db.Get()

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()
		fmt.Println()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			utils.ShowingInfo.Printf("Title: %s\n", key)

			if recursively {
				var curRecord model.BoltDbStruct
				err := json.Unmarshal(value, &curRecord)

				if err != nil {
					utils.ShowingError.Println("Error loading data from db")
				}

				for noteIndex := 0; noteIndex < len(curRecord.Notes); noteIndex += 1 {

					currentNote := curRecord.Notes[noteIndex]
					noteFirstLine := currentNote.Content
					formattedID := currentNote.Id[:8]
					maxPreviewLength := 40

					if len(noteFirstLine) > maxPreviewLength {
						noteFirstLine = noteFirstLine[:maxPreviewLength]
						noteFirstLine = strings.Join([]string{noteFirstLine, "..."}, "")
					}

					utils.ShowingInfo.Printf("   ↳ ID: %s | Content: %s \n", formattedID, noteFirstLine)
				}

				utils.ShowingInfo.Println(strings.Repeat("-", 120))

			}

			fmt.Println()
		}

		return nil
	})

	return nil
}
