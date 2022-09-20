package main

import (
	"context"
	"log"
	"os"

	"github.com/kumin/GolangMaster/database/mysql"
)

func InitSectionFields(ctx context.Context) {
}

func InitFieldValues(ctx context.Context) {
}

func main() {
	initor := mysql.NewInitDatabase()
	ctx := context.Background()
	data, err := initor.ReadJsondataFromFile("./mysql/targeting_fields.json")
	if err != nil {
		log.Fatalf("Read data from file error: %v", err)
		os.Exit(1)
	}
	dao := &mysql.SectionFieldsDAO{
		Section: "targeting",
		Fields:  data,
	}
	if err = initor.InsertBatches(ctx, dao); err != nil {
		log.Fatal("Insert data error")
		os.Exit(1)
	}
}
