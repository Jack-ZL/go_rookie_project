package common

import (
	"github.com/Jack-ZL/go_rookie"
	"reflect"
	"strconv"
)

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			if f.IsNil() {
				f.Set(reflect.MakeMap(ft.Type))
			}
		case reflect.Slice:
			if f.IsNil() {
				f.Set(reflect.MakeSlice(ft.Type, 0, 0))
			} else {
				for i := 0; i < f.Len(); i++ {
					_v := f.Index(i)
					if _v.Type().Kind() == reflect.Ptr && _v.Type().Elem().Kind() == reflect.Struct {
						initializeStruct(_v.Type().Elem(), _v.Elem())
					}
				}
			}
		case reflect.Chan:
			if f.IsNil() {
				f.Set(reflect.MakeChan(ft.Type, 0))
			}
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			if f.IsNil() {
				fv := reflect.New(ft.Type.Elem())
				if ft.Type.Elem().Kind() == reflect.Struct {
					initializeStruct(ft.Type.Elem(), fv.Elem())
				}
				f.Set(fv)
			} else {
				initializeStruct(ft.Type.Elem(), f.Elem())
			}
		default:
		}
	}
}

func RefineNil(t any) {
	initializeStruct(reflect.TypeOf(t).Elem(), reflect.ValueOf(t).Elem())
}

func GetFilterAndLimitOffset(c *go_rookie.Context) (map[string]string, int64, int64, error) {
	filter := make(map[string]string)
	for param, value := range c.Query() {
		if param != "limit" && param != "offset" {
			if len(value) == 0 {
				filter[param] = ""
				continue
			}
			filter[param] = value[0]
		}
	}

	limit, intParseErr := strconv.Atoi(c.GetDefaultQuery("limit", strconv.Itoa(DEFAULT_LIMIT)))
	if intParseErr != nil {
		Log.Error("Could not part the limit parameter")
		return nil, 0, 0, intParseErr
	}

	offset, intParseErr := strconv.Atoi(c.GetDefaultQuery("offset", strconv.Itoa(DEFAULT_OFFSET)))
	if intParseErr != nil {
		Log.Error("Could not part the offset parameter")
		return nil, 0, 0, intParseErr
	}

	if limit > 100 {
		limit = 100
	}

	return filter, int64(limit), int64(offset), nil
}
