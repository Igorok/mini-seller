package main

import (
	"log"
	"mini-seller/infrastructure/mongohelper"
	"mini-seller/infrastructure/mongohelper/testdata"
	"mini-seller/infrastructure/viperhelper"
)

func main() {
	vip := viperhelper.Viper{ConfigType: "", ConfigName: "", ConfigPath: "infrastructure/viperhelper"}
	vip.Read()

	db, err := mongohelper.Connect()
	if err != nil {
		log.Fatal(err)
	}

	testdata.InsertOrganizations(db)
	testdata.InsertProducts(db)
	testdata.InsertEmployee(db)
}
