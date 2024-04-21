package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"strings"
)

func getOrSetKey() error {
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		return errors.New("error opening database: " + err.Error())
	}
	defer db.Close()

	//create the bucket ifit doesnt exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return errors.New("error creating database bucket: " + err.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}

	var apiKey string
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MyBucket"))
		bytes := bucket.Get([]byte("OPENAI_API_KEY"))
		apiKey = string(bytes)
		return nil
	})

	if apiKey == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter OPENAI API Key: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return errors.New("Error reading input: " + err.Error())
		}
		apiKey = strings.TrimSpace(input)
		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("MyBucket"))
			err := bucket.Put([]byte("OPENAI_API_KEY"), []byte(apiKey))
			if err != nil {
				return errors.New("Error setting key to db: " + err.Error())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	err = os.Setenv("OPENAI_API_KEY", apiKey)
	if err != nil {
		return errors.New("error setting env variable key: " + err.Error())
	}
	return nil
}

func deleteKey() error {
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		return errors.New("error opening database: " + err.Error())
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MyBucket"))
		if bucket == nil {
			return nil
		}
		err := bucket.Delete([]byte("OPENAI_API_KEY"))
		if err != nil {
			return errors.New("error deleting key from db: " + err.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
