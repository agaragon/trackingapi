package dal


type DbBase interface {
	Save(object interface{},tableName string) error
	FilterById(object []byte, tableName string) interface{}
}