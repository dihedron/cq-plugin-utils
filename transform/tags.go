package transform

import (
	"reflect"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const (
	// NameTag is the tag applied to structs fields to specify a column name.
	NameTag = "cq-name"
	// TypeTag is the tag applied to structs fields to specify a column type.
	TypeTag = "cq-type"
)

// TagNameTransformer checks if a struct field is annotated with a "cq-name" tag
// and if so uses its value as the name of the column in the generated schema.
func TagNameTransformer(field reflect.StructField) (string, error) {
	// log.Printf("name transformer: %v\n", field)
	if cq := field.Tag.Get(NameTag); cq != "" {
		return cq, nil
	}

	return transformers.DefaultNameTransformer(field)
}

// TagTypeTransformer checks if a struct field is annotated with a "cq-type" tag
// and if so uses its value as the type of the column in the generated schema.
func TagTypeTransformer(field reflect.StructField) (arrow.DataType, error) {
	// log.Printf("type transformer: %v\n", field)
	if cq := field.Tag.Get(TypeTag); cq != "" {
		switch strings.ToLower(cq) {

		case "bool":
			return arrow.FixedWidthTypes.Boolean, nil
		case "int":
			return arrow.PrimitiveTypes.Int64, nil
		case "float":
			return arrow.PrimitiveTypes.Float64, nil
		case "uuid":
			return types.ExtensionTypes.UUID, nil
		case "string":
			return arrow.BinaryTypes.String, nil
		case "[]byte":
			return arrow.BinaryTypes.Binary, nil
		case "[]string":
			return arrow.ListOf(arrow.BinaryTypes.String), nil
		case "[]int":
			return arrow.ListOf(arrow.PrimitiveTypes.Int64), nil
		case "timestamp":
			return arrow.FixedWidthTypes.Timestamp_us, nil
		case "json":
			return types.ExtensionTypes.JSON, nil
		case "[]uuid":
			return arrow.ListOf(types.ExtensionTypes.UUID), nil
		case "inet":
			return types.ExtensionTypes.Inet, nil
		case "[]inet":
			return arrow.ListOf(types.ExtensionTypes.Inet), nil
		case "cidr":
			return types.ExtensionTypes.Inet, nil
		case "[]cidr":
			return arrow.ListOf(types.ExtensionTypes.Inet), nil
		case "mac", "macaddr":
			return types.ExtensionTypes.MAC, nil
		case "[]mac", "[]Macaddr":
			return arrow.ListOf(types.ExtensionTypes.MAC), nil
		}
	}
	return transformers.DefaultTypeTransformer(field)
}

// var defaultCaser = caser.New()

// func DefaultNameTransformer(field reflect.StructField) (string, error) {
// 	name := field.Name
// 	if jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]; len(jsonTag) > 0 {
// 		// return empty string if the field is not related api response
// 		if jsonTag == "-" {
// 			return "", nil
// 		}
// 		if nameFromJSONTag := defaultCaser.ToSnake(jsonTag); schema.ValidColumnName(nameFromJSONTag) {
// 			return nameFromJSONTag, nil
// 		}
// 	}
// 	return defaultCaser.ToSnake(name), nil
// }
// */
