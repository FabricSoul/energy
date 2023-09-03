//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package tools

import (
	"bytes"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
)

// CommandExists 命令是否存在
func CommandExists(name string) bool {
	_, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	return true
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func StrToInt64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}
func StrToFloat64(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func StrToFloat32(value string) float32 {
	v, _ := strconv.ParseFloat(value, 32)
	return float32(v)
}

func ToInt(v interface{}) int {
	switch v.(type) {
	case string:
		return int(StrToInt64(v.(string)))
	case float32:
		return int(math.Round(float64(StrToFloat32(v.(string)))))
	case float64:
		return int(math.Round(StrToFloat64(v.(string))))
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	case int:
		return v.(int)
	case int8:
		return int(v.(int8))
	case int16:
		return int(v.(int16))
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case uintptr:
		return int(v.(uintptr))
	default:
		return 0
	}
}

func ToRNilString(v interface{}, new string) string {
	if v == nil {
		return new
	}
	return v.(string)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func RenderTemplate(templateText string, data map[string]any) ([]byte, error) {
	tmpl, err := template.New("").Parse(templateText)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func HttpRequestGET(url string) ([]byte, error) {
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Compare compare2 < compare1 = true
func Compare(compare1, compare2 string) bool {
	if compare1[0] == 'v' {
		compare1 = compare1[1:]
	}
	if compare2[0] == 'v' {
		compare2 = compare2[1:]
	}
	compare1 = strings.Split(compare1, "-")[0]
	compare2 = strings.Split(compare2, "-")[0]
	cv := strings.Split(compare1, ".")
	ev := strings.Split(compare2, ".")
	c0, _ := strconv.Atoi(cv[0])
	c1, _ := strconv.Atoi(cv[1])
	c2, _ := strconv.Atoi(cv[2])
	e0, _ := strconv.Atoi(ev[0])
	e1, _ := strconv.Atoi(ev[1])
	e2, _ := strconv.Atoi(ev[2])
	if e0 < c0 {
		return true
	}
	if e1 < c1 {
		return true
	}
	if e2 < c2 {
		return true
	}
	return false
}
