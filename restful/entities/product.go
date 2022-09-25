package entities

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
)

type Properties struct {
	Color    string  `json:"color,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Category string  `json:"category,omitempty"`
}

type Product struct {
	ID         int64       `json:"id,omitempty"`
	Title      string      `json:"title,omitempty"`
	Properties *Properties `json:"properties,omitempty" gorm:"serializer:json"`
}

func (p *Product) TablelName() string {
	return "products"
}

type JSONSerializer struct {
}

func (j *JSONSerializer) Scan(
	ctx context.Context,
	field *schema.Field,
	dst reflect.Value,
	dbValue interface{},
) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSON value: %#v", dbValue)
		}

		err = json.Unmarshal(bytes, fieldValue.Interface())
	}
	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

func (j *JSONSerializer) Value(
	ctx context.Context,
	field *schema.Field,
	dst reflect.Value,
	fieldValue interface{},
) (interface{}, error) {
	return json.Marshal(fieldValue)
}
