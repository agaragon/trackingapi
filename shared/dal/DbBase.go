package dal


type DbBase interface {
	Save(object interface{},tableName string) error
	// Filter(object map[string]interface{}) interface{}
	Filter(object []byte, tableName string) interface{}
}