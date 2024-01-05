package main

import (
	"log"

	"github.com/yimsoijoi/migo/svc"
)

func main() {
	/*
		TODO:
			- Get config from env
		    - Docker
			- README
			- (xlsx styling)

	*/
	dbName := "db-name"
	db, err := svc.ConnectDB(
		"my-db.com",
		"db-user",
		"db-password",
		dbName,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	dbModel, err := svc.ExecDB(db)
	if err != nil {
		log.Fatalln(err)
	}

	excelModel := svc.ToExcelModel(dbModel)

	file := svc.BuildXlsx(excelModel)
	if file == nil {
		log.Fatalln("build failed")
	}

	if err = file.SaveAs(dbName + `.xlsx`); err != nil {
		log.Fatalln("failed to save: ", dbName, err)
	}

	log.Println("Done...")
}
