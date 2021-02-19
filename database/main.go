package main

import (
	"context"
	"log"
	"os"

	"github.com/my-packages/database/mysql"
)

func main() {
	initor := mysql.NewInitDatabase()
	ctx := context.Background()
	data, err := initor.ReadJsondataFromFile("./segment_fields.json")
	if err != nil {
		log.Fatalf("Read data from file error: %v", err)
		os.Exit(1)
	}
	dao := &mysql.SectionFieldsDAO{
		Section: "segment",
		Fields:  data,
	}
	if err = initor.InsertBatches(ctx, dao); err != nil {
		log.Fatal("Insert data error")
		os.Exit(1)
	}
}
