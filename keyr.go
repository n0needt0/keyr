/**
Copyright 2015 andrew@yasinsky.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keyr

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

//this is Keyr structure it is thread safe
type Keyr struct {
	sync.Mutex
	meta map[string]interface{}
}

//create new keyr
func NewKeyr(data map[string]interface{}) Keyr {
	return Keyr{meta: data}
}

/*return value from map as interface*/
func (bt Keyr) AddKeyVal(key string, val interface{}) {
	bt.Lock()
	bt.meta[key] = val
	bt.Unlock()
}

/*return value from map as interface*/
func (bt Keyr) GetAll() map[string]interface{} {
	bt.Lock()
	defer bt.Unlock()
	return bt.meta
}

/*return value from map as interface*/
func (bt Keyr) GetKeyAsInterface(key string) (interface{}, error) {
	bt.Lock()
	defer bt.Unlock()
	if key == "" {
		return nil, errors.New("Keyr:GetKeyAsInterface (key is empty)")
	}
	if el, ok := bt.meta[key]; ok {
		return el, nil
	}
	return nil, nil
}

/*return value from map as string*/
func (bt Keyr) GetKeyAsString(key string) (string, error) {
	bt.Lock()
	defer bt.Unlock()
	elinterface, err := bt.GetKeyAsInterface(key)
	if err != nil {
		return "", err
	}

	if elinterface == nil {
		return "", nil
	}

	s := ""

	switch T := elinterface.(type) {
	case string:
		s = reflect.ValueOf(T).String()
	case bool:
		s = fmt.Sprintf("%t", reflect.ValueOf(T).Bool())
	case json.Number:
		s = T.String()
	case int, int8, int16, int32, int64:
		s = fmt.Sprintf("%d", reflect.ValueOf(T).Int())
	case float32, float64:
		s = fmt.Sprintf("%g", reflect.ValueOf(T).Float())
	}

	return strings.TrimSpace(s), nil
}

//	return nil, errors.New(fmt.Sprintf("Keyr: GetKeyAsString bad input value %+v", elinterface))

/*return value from map as Int*/
func (bt Keyr) GetKeyAsInt(key string) (int, error) {
	bt.Lock()
	defer bt.Unlock()
	elstr, err := bt.GetKeyAsString(key)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if elstr == "" {
		return 0, nil
	}

	elint, err := strconv.ParseFloat(elstr, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Can not parse to int : %+v", elstr))
	}

	return int(math.Floor(elint + .5)), nil
}

/*return value from map as Float*/
func (bt Keyr) GetKeyAsFloat(key string) (float64, error) {
	bt.Lock()
	defer bt.Unlock()
	elstr, err := bt.GetKeyAsString(key)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if elstr == "" {
		return 0, nil
	}

	elfloat, err := strconv.ParseFloat(elstr, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Can not parse to float64 : %+v", elstr))
	}

	return elfloat, nil
}

func (bt Keyr) GetKeyAsBool(key string) (bool, error) {
	bt.Lock()
	defer bt.Unlock()
	elstr, err := bt.GetKeyAsString(key)
	if err != nil {
		return false, err
	}

	if elstr == "" {
		return false, nil
	}

	elbool, err := strconv.ParseBool(elstr)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Can not parse to bool : %+v", elstr))
	}

	return elbool, nil
}

func (bt Keyr) StringsContainString(hay []string, needle string) bool {
	bt.Lock()
	defer bt.Unlock()
	for _, valid := range hay {
		if valid == needle {
			return true
		}
	}
	return false
}
