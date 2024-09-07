package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/Jack-ZL/go_rookie"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
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
		return nil, 0, 0, intParseErr
	}

	offset, intParseErr := strconv.Atoi(c.GetDefaultQuery("offset", strconv.Itoa(DEFAULT_OFFSET)))
	if intParseErr != nil {
		return nil, 0, 0, intParseErr
	}

	if limit > 100 {
		limit = 100
	}

	return filter, int64(limit), int64(offset), nil
}

func DoubleStrMerge(str1, str2 string) string {
	slen1 := len(str1)
	slen2 := len(str2)
	if slen1 == 0 || slen2 == 0 {
		return ""
	}

	pwd_len := 0
	if slen1 >= slen2 {
		pwd_len = slen1
	} else {
		pwd_len = slen2
	}
	var newstr string
	for i := 0; i < pwd_len; i++ {
		if i < slen1 {
			child1 := str1[i : i+1]
			newstr += child1
		}

		if i < slen2 {
			child2 := str2[i : i+1]
			newstr += child2
		}
	}
	return newstr
}

/**
 * Md5Encrypt
 * @Description: md5加密
 * @Author: Jazk-Z
 * @param str
 * @return string
 */
func Md5Encrypt(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * MakeSignature
 * @Description: 生成签名
 * @Author: Jazk-Z
 * @param timestamp
 * @param nonce
 * @return string
 */
func MakeSignature(key ...string) string {
	sort.Strings(key)
	// 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(key, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
