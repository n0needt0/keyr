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
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {

	newmeta := map[string]interface{}{"test": "blah"}
	k := NewKeyr(newmeta)

	result, err := k.GetKeyAsInterface("")
	assert.NotEqual(t, err, nil)

	result, err = k.GetKeyAsInterface("test")
	assert.NotEqual(t, result, nil)

	result, err = k.GetKeyAsInterface("test1")
	assert.Equal(t, result, nil)
}

func TestString(t *testing.T) {

	newmeta := map[string]interface{}{"test": "blah", "test1": 2}
	k := NewKeyr(newmeta)

	result, _ := k.GetKeyAsString("test")
	assert.Equal(t, result, "blah")

	result, _ = k.GetKeyAsString("test1")
	assert.Equal(t, result, "2", fmt.Sprintf("%+v", result))
}

func TestInt(t *testing.T) {

	newmeta := map[string]interface{}{"test": "blah", "test1": 1, "test2": "2", "test3": 3.3, "test4": "4.4", "test5": "4.9", "test6": json.Number("6")}
	k := NewKeyr(newmeta)

	result, err := k.GetKeyAsInt("test")
	assert.NotEqual(t, err, nil)

	result, _ = k.GetKeyAsInt("test1")
	assert.Equal(t, result, 1)

	result, _ = k.GetKeyAsInt("test2")
	assert.Equal(t, result, 2)

	result, _ = k.GetKeyAsInt("test3")
	assert.Equal(t, result, 3)

	result, _ = k.GetKeyAsInt("test4")
	assert.Equal(t, result, 4)

	result, _ = k.GetKeyAsInt("test5")
	assert.Equal(t, result, 5)

	result, _ = k.GetKeyAsInt("test6")
	assert.Equal(t, result, 6)
}

func TestFloat(t *testing.T) {

	newmeta := map[string]interface{}{"test": "blah", "test1": 1, "test2": "2", "test3": 3.3, "test4": "4.4", "test5": "4.9", "test6": json.Number("6")}
	k := NewKeyr(newmeta)

	result, err := k.GetKeyAsFloat("test")
	assert.NotEqual(t, err, nil)

	result, _ = k.GetKeyAsFloat("test1")
	assert.Equal(t, result, float64(1))

	result, _ = k.GetKeyAsFloat("test2")
	assert.Equal(t, result, float64(2))

	result, _ = k.GetKeyAsFloat("test3")
	assert.Equal(t, result, float64(3.3))

	result, _ = k.GetKeyAsFloat("test4")
	assert.Equal(t, result, float64(4.4))

	result, _ = k.GetKeyAsFloat("test5")
	assert.Equal(t, result, float64(4.9))

	result, _ = k.GetKeyAsFloat("test6")
	assert.Equal(t, result, float64(6))
}

func TestBool(t *testing.T) {

	newmeta := map[string]interface{}{"test": "blah", "test1": 1, "test2": "true", "test3": true, "test4": json.Number("1")}
	k := NewKeyr(newmeta)

	result, err := k.GetKeyAsBool("test")
	assert.NotEqual(t, err, nil)

	result, _ = k.GetKeyAsBool("test1")
	assert.Equal(t, result, true)

	result, _ = k.GetKeyAsBool("test2")
	assert.Equal(t, result, true)

	result, _ = k.GetKeyAsBool("test3")
	assert.Equal(t, result, true)

	result, _ = k.GetKeyAsBool("test4")
	assert.Equal(t, result, true)
}
