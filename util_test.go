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
