package dogego_module_dbseeder

import (
	"github.com/jinzhu/gorm"
	"log"
)

type DBSeeder struct {
	masterDB *gorm.DB
	slaveDB  *gorm.DB
}

func NewSeeder(masterDB *gorm.DB, slaveDB *gorm.DB) *DBSeeder {
	return &DBSeeder{
		masterDB: masterDB,
		slaveDB:  slaveDB,
	}
}

func (seeder *DBSeeder) Seed(table_name string, seed_list []interface{}) *DBSeeder {
	for _, seed_data := range seed_list {
		datas := make([]interface{}, 0)
		err := seeder.slaveDB.Table(table_name).Where(&datas, seed_data).Error

		if err != nil {
			log.Fatal(err)
		}

		if len(datas) == 0 {
			err := seeder.masterDB.Create(seed_data).Error

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return seeder
}
