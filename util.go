package phpserialize

// StringifyKeys recursively converts a map into a more sensible map with
// strings as keys.
//
// map[interface{}]interface{} is used as an unmarshalling format because PHP
// serialise() permits keys of associative arrays to be non-string. However, in
// reality this is rarely the case and so strings for keys are much more
// compatible with external code.
func StringifyKeys(m map[interface{}]interface{}) map[string]interface{} {
	return map[string]interface{}{}
}
