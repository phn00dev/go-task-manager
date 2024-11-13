package seeders

import "gorm.io/gorm"

type DBSeeder struct {
	db *gorm.DB
}

func NewDbSeeder(db *gorm.DB) *DBSeeder {
	return &DBSeeder{db: db}
}

func (dbSeeder DBSeeder) GetDBSeeder() error {
	// admin seeder
	if err := dbSeeder.adminSeeder(); err != nil {
		return err
	}
	// team seeder
	if err := dbSeeder.teamSeeder(); err != nil {
		return err
	}
	return nil
}
