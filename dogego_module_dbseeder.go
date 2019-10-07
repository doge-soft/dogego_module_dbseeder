package dogego_module_dbseeder

import "github.com/jinzhu/gorm"

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

func (seeder *DBSeeder) Seed(seed_list []interface{}) {

}
