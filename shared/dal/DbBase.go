package dal


type DbBase interface {
	Save(object interface{},tableName string) error
}