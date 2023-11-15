package maps

// GetKeyValSlice concatenates the key and value for each key-value pair in a map and returns
// a slice of the concatenations.
// For example: input {"a" : "1", "b" : "2"} should return {"a1", "b2"}
// Input of nil or empty map should return an uninitialized slice (or nil)
func GetKeyValSlice(input map[string]string) []string {
	var result []string
	for key, value := range input {
		result = append(result, key+value)
	}
	return result
}

// DeleteFromMap deletes a list of keys from a map.
// If any key(s) are not in the map, those keys can be skipped.
// If the map is nil, do nothing. (hint: check the docs for map key deletion to figure out an easy way to "handle" this)
// For example: input {1: 100, 2: 200}, {2, 7} should alter the map to be: {1: 100}
func DeleteFromMap(input map[int]int, deleteKeys []int) {
	if input == nil {
		return
	}
	for _, key := range deleteKeys {
		delete(input, key)
	}

}

// SetAndGet sets the input key-value pair in the input map. Then it gets getKey from the map and returns the value.
// If the map is nil, SetAndGet does nothing and returns the zero value for string
// If getKey does not exist in the map, returns "NOTEXIST"
// The get for getKey -should include- what is set for setKey & setVal (do the set before the get)
func SetAndGet(input map[float64]string, setKey float64, setVal string, getKey float64) string {
	if input == nil {
		return ""
	}
	input[setKey] = setVal
	resultValue, ok := input[getKey]
	if ok {
		return resultValue
	}
	return "NOTEXIST"
}
