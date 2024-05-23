package db

import (
	"fmt"
)

type Collection struct {
	// your collection structure
}

type Pocketbase struct {
	// your database structure
}

func LoadDataFromCollection(db *Pocketbase, collectionName string) (*Collection, error) {
	// Assuming db has a method GetCollection that returns a Collection
	collection, err := db.GetCollection(collectionName)
	if err != nil {
		fmt.Println("Error loading collection:", err)
		return nil, err
	}

	// Assuming Collection has a method LoadData that loads the data
	err = collection.LoadData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return nil, err
	}

	return collection, nil
}
