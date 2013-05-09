// Copyright 2013 Landon Wainwright. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package configuration

import (
	"bytes"
	"encoding/json"
	"os"
)

// Holds the JSON configuration data
type Configuration struct {
	filePath   string
	properties map[string]interface{}
}

// Creates a new empty configuration file
func Empty() *Configuration {
	return NewFromFile("")
}

// Creates a new empty configuration file that can be read into or written to disk
func NewFromFile(filePath string) *Configuration {
	configuration := &Configuration{filePath: filePath, properties: make(map[string]interface{})}
	return configuration
}

// Will update the current file path to allow the configuration to be saved to a new location
func (configuration *Configuration) UpdateFilePath(filePath string) {
	configuration.filePath = filePath
}

// Loads the configuration from disk into this Configuration object
func (configuration *Configuration) ReadFromDisk() (err error) {

	// Attempt to open the file
	f, err := os.Open(configuration.filePath)
	if err == nil {
		defer f.Close()
		var b bytes.Buffer
		_, err = b.ReadFrom(f)
		if err == nil {
			err = json.Unmarshal(b.Bytes(), &configuration.properties)
		}
	}
	return
}

// Writes the current data within the Configuration to disk
func (configuration *Configuration) SaveToDisk() (err error) {

	// Marshall this to disk
	b, err := json.Marshal(configuration.properties)
	if err == nil {

		// Get access to the file
		fo, err := os.Create(configuration.filePath)
		defer fo.Close()
		if err == nil {

			// Write the bytes to the buffer
			var buffer bytes.Buffer
			_, err = buffer.Write(b)
			if err == nil {

				// Write the buffer to disk
				_, err = buffer.WriteTo(fo)
			}
		}
	}
	return
}

// Deletes the property from the store
func (configuration *Configuration) DeleteProperty(key string) {
	delete(configuration.properties, key)
}

// Returns a bool for the property key specified
func (configuration *Configuration) GetBool(key string) bool {
	val, exists := configuration.properties[key]
	if !exists {
		return false
	}
	return val.(bool)
}

// Returns an int for the property key specified
func (configuration *Configuration) GetInt(key string) int {
	val, exists := configuration.properties[key]
	if !exists {
		return -1
	}
	return val.(int)
}

// Returns a float for the property key specified
func (configuration *Configuration) GetFloat(key string) float64 {
	val, exists := configuration.properties[key]
	if !exists {
		return -1
	}
	return val.(float64)
}

// Returns a string for the property key specified
func (configuration *Configuration) GetString(key string) string {
	val, exists := configuration.properties[key]
	if !exists {
		return ""
	}
	return val.(string)
}

// Returns an array for the property key specified
func (configuration *Configuration) GetArray(key string) []interface{} {
	val, exists := configuration.properties[key]
	if !exists {
		return []interface{}(nil)
	}
	return val.([]interface{})
}

// Returns a map for the key specified
func (configuration *Configuration) GetMap(key string) map[string]interface{} {
	val, exists := configuration.properties[key]
	if !exists {
		return make(map[string]interface{})
	}
	return val.(map[string]interface{})
}

// Returns a bool for the property key specified
func (configuration *Configuration) SetBool(key string, val bool) {
	configuration.setProperty(key, val)
}

// Returns an int for the property key specified
func (configuration *Configuration) SetInt(key string, val int) {
	configuration.setProperty(key, val)
}

// Returns a float for the property key specified
func (configuration *Configuration) SetFloat(key string, val float64) {
	configuration.setProperty(key, val)
}

// Returns a string for the property key specified
func (configuration *Configuration) SetString(key string, val string) {
	configuration.setProperty(key, val)
}

// Returns an array for the property key specified
func (configuration *Configuration) SetArray(key string, val []interface{}) {
	configuration.setProperty(key, val)
}

// Returns a map for the key specified
func (configuration *Configuration) SetMap(key string, val map[string]interface{}) {
	configuration.setProperty(key, val)
}

// Just updates the map data store for the given property key
func (configuration *Configuration) setProperty(key string, value interface{}) {
	configuration.properties[key] = value
}
