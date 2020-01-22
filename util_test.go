package phpserialize_test

import (
	"github.com/elliotchance/phpserialize"
	"reflect"
	"testing"
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

func TestStringifyKeysOnMapWithEntriesRecursively(t *testing.T) {
	m := map[interface{}]interface{}{
		"foo": map[interface{}]interface{}{
			"foo": "bar",
			"baz": 123,
		},
		"baz": 123,
	}
	result := phpserialize.StringifyKeys(m)
	expected := map[string]interface{}{
		"foo": map[string]interface{}{
			"foo": "bar",
			"baz": 123,
		},
		"baz": 123,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

func TestStringifyKeysOnMapRecursivelyMovesThroughSlices(t *testing.T) {
	m := map[interface{}]interface{}{
		"foo": []interface{}{
			map[interface{}]interface{}{
				"foo": "bar",
				"baz": 123,
			},
			123,
		},
		"baz": 123,
	}
	result := phpserialize.StringifyKeys(m)
	expected := map[string]interface{}{
		"foo": []interface{}{
			map[string]interface{}{
				"foo": "bar",
				"baz": 123,
			},
			123,
		},
		"baz": 123,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}
