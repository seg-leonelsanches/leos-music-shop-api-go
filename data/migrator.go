package data

import (
	"segment/leos-music-shop-api-go/models"
)

func Migrate() {
	// Migrate the schema
	Db.AutoMigrate(&models.Keyboard{})
	Db.AutoMigrate(&models.Manufacturer{})
	Db.AutoMigrate(&models.MigrationState{})

	// Create
	var migrationState models.MigrationState
	result := Db.First(&migrationState, "name = ?", "Initial")
	if result.RowsAffected == 0 {
		Db.Create(&models.MigrationState{
			Id:   "1",
			Name: "Initial",
		})
		Db.Create([]models.Keyboard{
			{Id: "1", Model: "Williams Allegro III", Manufacturer: "Williams", Price: 349.99},
			{Id: "2", Model: "Yamaha P-125", Manufacturer: "Yamaha", Price: 699.99},
			{Id: "3", Model: "Casio CDP-S100", Manufacturer: "Casio", Price: 449.99},
		})
		Db.Create([]models.Manufacturer{
			{Id: "1", Name: "Williams"},
			{Id: "2", Name: "Yamaha"},
			{Id: "3", Name: "Casio"},
		})
	}

}
