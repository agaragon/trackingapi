package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "trackingApp/shared/logger"
	. "trackingApp/shared/customError"
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
		LogError(err)
		return &ErrorSaving{"Unable to save on database"}
	}
	return nil

}