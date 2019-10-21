package framework

type ModelInterface interface {
	Table(string) string
	Find() *ModelInterface
	Set(string, interface{}) bool
	Get(string) interface{}
	Save()
}

type Model struct {
	TableName string
	ID        int
}

func (model *Model) GetTable() string {
	return model.TableName
}

func (model *Model) SetTable(table string) {
	model.TableName = table
}

func (model *Model) find(id int) *Model {
	model.ID = id
	return model
}

func (model *Model) Save() {
}
