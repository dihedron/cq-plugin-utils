package transform

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

// Transform is the type of function that transforms a value into
// another value.
type Transform func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column, v any) (any, error)

// Apply returns a ColumnResolver that is made up of a sequence of
// Transforms whose purpose is that of extracting information from
// the input, changing it along the way and returning the final
// value to be stored in the database.
func Apply(transforms ...Transform) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var (
			value any
			err   error
		)

		for _, transform := range transforms {
			value, err = transform(ctx, meta, r, c, value)
			if err != nil {
				return err
			}
		}
		return r.Set(c.Name, value)
	}
}

// OnObjectField gets the value from a Golang object by extracting
// the value associated with the field as per the path.
func OnObjectField(path string) Transform {
	return func(ctx context.Context, _ schema.ClientMeta, r *schema.Resource, _ schema.Column, _ any) (any, error) {
		value := funk.Get(r.Item, path, funk.WithAllowZero())
		return value, nil
	}
}

// AssertType checks if the input value is of the given (generic) type.
func AssertType[T any]() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, r *schema.Resource, _ schema.Column, v any) (any, error) {
		if v == nil {
			return nil, nil
		}
		if _, ok := v.(T); !ok {
			var t T
			return nil, fmt.Errorf("invalid data type: expected %T, got %T", t, v)
		}
		return v, nil
	}
}

// GetElementAt gets the n-th element from a slice; it returns an error
// if the value is not an array.
func GetElementAt(n int) Transform {
	return func(ctx context.Context, _ schema.ClientMeta, r *schema.Resource, _ schema.Column, v any) (any, error) {
		value := funk.ElementAt(v, n)
		return value, nil
	}
}

// GetMapEntry returns the value associated with the given map key; it assumes
// that the input value is a map where the key is of type K, and the value
// is of type V.
func GetMapEntry[K comparable, V any](key K) Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if v != nil {
			switch value := v.(type) {
			case map[K]V:
				if t, ok := value[key]; ok {
					return t, nil
				}
			}
		}
		return nil, nil
	}
}

// TrimString assumes that the current valule is a string and trims it
// of its spaces.
func TrimString() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if v == nil {
			return nil, nil
		}
		if s, ok := v.(string); ok {
			return strings.TrimSpace(s), nil
		}
		return nil, fmt.Errorf("invalid type: expected string, got %T", v)
	}
}

// NilIfZero returns nil if the current value is the zero value of
// its respective type.
func NilIfZero() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if v == nil {
			return nil, nil
		}
		if funk.IsZero(v) {
			return nil, nil
		}
		return v, nil
	}
}

// ToString converts the current value into its string representation.
func ToString() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if v == nil {
			return nil, nil
		}
		switch v := v.(type) {
		case string:
			return v, nil
		default:
			return fmt.Sprintf("%v", v), nil
		}
	}
}

// ToInt converts the value to an int; numeric types are cast into an int,
// whereas strings are parsed.
func ToInt() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		switch v := v.(type) {
		case int:
			return v, nil
		case int8:
			return int(v), nil
		case int16:
			return int(v), nil
		case int32:
			return int(v), nil
		case int64:
			return int(v), nil
		case uint:
			return v, nil
		case uint8:
			return int(v), nil
		case uint16:
			return int(v), nil
		case uint32:
			return int(v), nil
		case uint64:
			return int(v), nil
		case float32:
			return int(v), nil
		case float64:
			return int(v), nil
		case string:
			if strings.TrimSpace(v) != "" {
				return strconv.ParseInt(v, 10, 32)
			}
		default:
			// log.Printf("ToInt: unsupported type: %T", v)
			return nil, fmt.Errorf("unsupported value type: %T", v)
		}
		return nil, nil
	}
}

// ToBool converts the current value into its bool representation.
func ToBool() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if v == nil {
			return nil, nil
		}
		switch v := v.(type) {
		case bool:
			return v, nil
		case string:
			if strings.TrimSpace(v) != "" {
				return strconv.ParseBool(v)
			}
			return v, nil
		default:
			return nil, fmt.Errorf("incompatible type: %T", v)
		}
	}
}

// OrDefault sets the current value to the given default value
// if it is nil.
func OrDefault(value any) Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if v == nil {
			return value, nil
		}
		return v, nil
	}
}

// RemapValue remaps the current value to a different value according
// to the input map; it can be used to convert e.g. integer values into
// their string representations.
func RemapValue[K comparable, V any](remap map[K]V) Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		if key, ok := v.(K); ok {
			if value, ok := remap[key]; ok {
				return value, nil
			} else {
				return v, nil
			}
		}
		var t K
		return nil, fmt.Errorf("invalid data type: expected %T, got %T", t, v)
	}
}

// DecodeBase64 decodes the input value (if it is a string) from base64
// and returns it as a string.
func DecodeBase64() Transform {
	return func(ctx context.Context, _ schema.ClientMeta, _ *schema.Resource, _ schema.Column, v any) (any, error) {
		log.Printf("DecodeBase64: got %v (type %T)", v, v)
		if v == nil {
			return nil, nil
		}
		if v, ok := v.(string); ok {
			decoded, err := base64.StdEncoding.DecodeString(v)
			if err == nil {
				return string(decoded), nil
			} else {
				return nil, err
			}
		}
		return nil, fmt.Errorf("invalid data type: expected string, got %T", v)
	}
}
