package reflect_more_test

import (
	"github.com/pkg/errors"
	"reflect"
)

type Employee struct {
	EmployeeID string
	Name string `format:"normal"`
	Age int
}

func (e *Employee) updataAge(age int)  {
	e.Age= age
}

func fillBySetting(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("type must be a pointer")
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to struct type")
		}
	}

	if settings == nil {
		return errors.New("")
	}

	var(
		field reflect.StructField
		ok bool
	)

	for k,v := range settings {
		if field, ok = reflect.ValueOf(st).Elem().Type().FieldByName(k); !ok{
			continue
		}

		if field.Type == reflect.TypeOf(v){
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}


type Customer struct {
	CookieID string
	Name string
	Age int
}



