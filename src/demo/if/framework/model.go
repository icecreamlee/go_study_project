package framework

import (
	"fmt"
	"reflect"
	"strings"
)

type ModelInterface interface {
	GetTable() string
	SetTable(string)
	Find() ModelInterface
	Set(ModelInterface, string, interface{})
	Get(ModelInterface, string) interface{}
	Save(ModelInterface, map[string]interface{}) int64
}

type Model struct {
	tableName string
	ID        int64
}

func (model *Model) GetTable() string {
	return model.tableName
}

func (model *Model) SetTable(table string) {
	model.tableName = table
}

func (model *Model) Set(m ModelInterface, key string, value interface{}) {
	reflect.ValueOf(m).Elem().FieldByName(key).Set(reflect.ValueOf(value))
}

func (model Model) Get(m ModelInterface, key string) interface{} {
	return reflect.ValueOf(m).Elem().FieldByName(key).Interface()
}

func (model *Model) Find() ModelInterface {
	return model
}

func (model *Model) Save(m ModelInterface, data map[string]interface{}) int64 {
	if _, ok := data["id"]; ok {
		model.ID = int64(data["id"].(int))
	}
	//db, _ := GetDB()

	var keys []string
	var values []string
	var args []interface{}
	for key, val := range data {
		keys = append(keys, key)
		values = append(values, "?")
		args = append(args, val)
	}

	var sql string
	if model.ID != 0 {
		args = append(args, model.ID)
		sql = "UPDATE " + m.GetTable() + " SET " + strings.Join(keys, "=?, ") + "=?" + " WHERE id=?"
		execArgs := append([]interface{}{sql}, args...)
		_, err := db.Exec(execArgs...)
		if err != nil {
			fmt.Printf("UPDATE error: %s\n", err.Error())
			return 0
		}
		//db.Table(model.tableName).Where("id=?", model.ID).Updates(data)
	} else {
		sql = "INSERT INTO " + m.GetTable() + " (" + strings.Join(keys, ",") + ") VALUES (" + strings.Join(values, ",") + ")"
		execArgs := append([]interface{}{sql}, args...)
		result, err := db.Exec(execArgs...)
		if err != nil {
			fmt.Printf("Insert error: %s\n", err.Error())
			return 0
		}
		model.ID, err = result.LastInsertId()
		if err != nil {
			fmt.Printf("Insert error: %s\n", err.Error())
			return 0
		}
		//db.Table(model.tableName).Create()
	}
	fmt.Printf("SQL: %s\n, ARGS: %v", sql, args)
	return model.ID
}
