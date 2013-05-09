// Copyright 2013 Landon Wainwright. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package configuration

import (
	"testing"
)

// Struct for testing multiple value types
type testpair struct {
	key   string
	value interface{}
}

// The various values that we want to save and load from the 
var tests = []testpair{
	{"true", true},
	{"false", false},
	{"int1", 1},
	{"int100000", 100000},
	{"float641.4", 1.4},
	{"float64500.5", 500.5},
	{"string", "string"},
	{"array", []interface{}{1, 2, 3}},
	{"map", map[string]interface{}{"map1": 1, "map2": 2, "map3": 3, "map4": 4}},
}

// Tests adding and retrieving values to a memory backed configuration
func TestEmptyConfiguration(t *testing.T) {

	// Try to create a new configuration
	config := Empty()

	// Try to add some items to the list
	for _, pair := range tests {
		switch v := pair.value.(type) {
		case bool:
			config.SetBool(pair.key, pair.value.(bool))
			if config.GetBool(pair.key) != pair.value {
				t.Error(
					"Expected", pair.value,
					"Type", v,
					"Got", config.GetBool(pair.key),
				)
			}
		case int:
			config.SetInt(pair.key, pair.value.(int))
			if config.GetInt(pair.key) != pair.value {
				t.Error(
					"Expected", pair.value,
					"Type", v,
					"Got", config.GetInt(pair.key),
				)
			}
		case float64:
			config.SetFloat(pair.key, pair.value.(float64))
			if config.GetFloat(pair.key) != pair.value {
				t.Error(
					"Expected", pair.value,
					"Type", v,
					"Got", config.GetFloat(pair.key),
				)
			}
		case string:
			config.SetString(pair.key, pair.value.(string))
			if config.GetString(pair.key) != pair.value {
				t.Error(
					"Expected", pair.value,
					"Type", v,
					"Got", config.GetString(pair.key),
				)
			}
		case []interface{}:
			config.SetArray(pair.key, pair.value.([]interface{}))
			//if config.GetArray(pair.key) != pair.value {
			//	t.Error(
			//		"Expected", pair.value,
			//		"Type", v,
			//		"Got", config.GetArray(pair.key),
			//	)
			//}
		case map[string]interface{}:
			config.SetMap(pair.key, pair.value.(map[string]interface{}))
			//if config.GetMap(pair.key) != pair.value {
			//	t.Error(
			//		"Expected", pair.value,
			//		"Type", v,
			//		"Got", config.GetMap(pair.key),
			//	)
			//}
		default:
			{
				t.Error(
					"Unexpected type", v,
				)
			}
		}
	}
}
