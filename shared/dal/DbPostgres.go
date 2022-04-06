package dal

import (
	"github.com/jinzhu/gorm"
)

type DbPostgres struct {
	Uri string 
}

func (db *DbPostgres) Save(object interface{},tableName string) error{

	dbConn, err := gorm.Open("postgres",db.Uri)
	if err != nil {
		return err
	}
	err = dbConn.Create(object).Error
	if err != nil {
		return err
	}
	return nil

}