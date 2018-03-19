package phpserialize_test

import (
	"testing"
	"github.com/elliotchance/phpserialize"
	"reflect"
)

func TestStringifyKeysOnEmptyMap(t *testing.T) {
	m := map[interface{}]interface{}{}
	result := phpserialize.StringifyKeys(m)
	expected := map[string]interface{}{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

func TestStringifyKeysOnMapWithEntries(t *testing.T) {
	m := map[interface{}]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	result := phpserialize.StringifyKeys(m)
	expected := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}
