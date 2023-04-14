package xformutils

import (
	"log"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
	log.Printf("name transformer: %v\n", field)
	if cq := field.Tag.Get(NameTag); cq != "" {
		return cq, nil
	}
	return transformers.DefaultNameTransformer(field)
}

// TagTypeTransformer checks if a struct field is annotated with a "cq-type" tag
// and if so uses its value as the type of the column in the generated schema.
func TagTypeTransformer(field reflect.StructField) (schema.ValueType, error) {
	log.Printf("type transformer: %v\n", field)
	if cq := field.Tag.Get(TypeTag); cq != "" {
		switch strings.ToLower(cq) {
		case "bool", "boolean":
			return schema.TypeBool, nil
		case "int", "integer":
			return schema.TypeInt, nil
		case "float":
			return schema.TypeFloat, nil
		case "uuid":
			return schema.TypeUUID, nil
		case "string":
			return schema.TypeString, nil
		case "byte-array", "[]byte", "bytearray":
			return schema.TypeByteArray, nil
		case "string-array", "[]string", "stringarray":
			return schema.TypeStringArray, nil
		case "int-array", "[]int", "intarray":
			return schema.TypeIntArray, nil
		case "timestamp":
			return schema.TypeTimestamp, nil
		case "json":
			return schema.TypeJSON, nil
		case "uuid-array", "[]uuid", "uuidarray":
			return schema.TypeUUIDArray, nil
		case "inet":
			return schema.TypeInet, nil
		case "inet-array", "[]inet", "inetarray":
			return schema.TypeInetArray, nil
		case "cidr":
			return schema.TypeCIDR, nil
		case "cidr-array", "[]cidr", "cidrarray":
			return schema.TypeCIDRArray, nil
		case "mac", "mac-addr", "mac-address":
			return schema.TypeMacAddr, nil
		case "mac-array", "mac-addr-array", "mac-address-array", "[]mac", "[]mac-addr", "[]mac-address":
			return schema.TypeMacAddrArray, nil
		case "time-interval":
			return schema.TypeTimeIntervalDeprecated, nil
		case "end":
			return schema.TypeEnd, nil
		}
	}
	return transformers.DefaultTypeTransformer(field)
}
