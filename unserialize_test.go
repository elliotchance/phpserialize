package phpserialize_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/elliotchance/phpserialize"
)

func expectErrorToNotHaveOccurred(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func expectErrorToEqual(t *testing.T, err1, err2 error) {
	if err1 == nil {
		t.Error("err1 is nil")
	}

	if err2 == nil {
		t.Error("err2 is nil")
	}

	if err1.Error() != err2.Error() {
		t.Errorf("Expected '%s' to be '%s'", err1, err2)
	}
}

func TestUnmarshalInt(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        int
		expectedError error
	}{
		"0":              {[]byte("i:0;"), 0, nil},
		"5":              {[]byte("i:5;"), 5, nil},
		"-8":             {[]byte("i:-8;"), -8, nil},
		"1000000":        {[]byte("i:1000000;"), 1000000, nil},
		"not an integer": {[]byte("N;"), 0, errors.New("not an integer")},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			t.Run("int", func(t *testing.T) {
				var result int
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != test.output {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int8", func(t *testing.T) {
				var result int8
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int8(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int16", func(t *testing.T) {
				var result int16
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int16(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int32", func(t *testing.T) {
				var result int32
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int32(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int64", func(t *testing.T) {
				var result int64
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int64(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint", func(t *testing.T) {
				var result uint
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint8", func(t *testing.T) {
				var result uint8
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint8(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint16", func(t *testing.T) {
				var result uint16
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint16(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint32", func(t *testing.T) {
				var result uint32
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint32(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint64", func(t *testing.T) {
				var result uint64
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint64(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})
		})
	}
}

func TestUnmarshalFloat(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        float64
		expectedError error
	}{
		"3.2":         {[]byte("d:3.2;"), 3.2, nil},
		"10.0":        {[]byte("d:10;"), 10.0, nil},
		"123.456789":  {[]byte("d:123.456789;"), 123.456789, nil},
		"1.23e9":      {[]byte("d:1230000000;"), 1.23e9, nil},
		"-17.23":      {[]byte("d:3.2;"), 3.2, nil},
		"not a float": {[]byte("N;"), 0.0, errors.New("not a float")},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			t.Run("float32", func(t *testing.T) {
				var result float32
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != float32(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("float64", func(t *testing.T) {
				var result float64
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != test.output {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})
		})
	}
}

func TestUnmarshalString(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        string
		expectedError error
	}{
		"''":            {[]byte("s:0:\"\";"), "", nil},
		"'Hello world'": {[]byte("s:11:\"Hello world\";"), "Hello world", nil},
		"'Björk Guðmundsdóttir'": {
			[]byte(`s:23:"Björk Guðmundsdóttir";`),
			"Björk Guðmundsdóttir",
			nil,
		},
		"not a string": {[]byte("N;"), "", errors.New("not a string")},
		"Backslash":    {[]byte("s:1:\"\\\";"), "\\", nil},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var result string
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if result != test.output {
					t.Errorf("Expected '%v', got '%v'", test.output, result)
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

func TestUnmarshalBinary(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        []byte
		expectedError error
	}{
		"[]byte: \\x01\\x02\\x03": {
			[]byte("s:3:\"\x01\x02\x03\";"),
			[]byte{1, 2, 3},
			nil,
		},
		"not a string": {[]byte("N;"), []byte{}, errors.New("not a string")},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var result []byte
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if string(result) != string(test.output) {
					t.Errorf("Expected '%v', got '%v'", result, test.output)
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

func TestUnmarshalArray(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        []interface{}
		expectedError error
	}{
		"[]interface{}: [7.89]": {
			[]byte("a:1:{i:0;d:7.89;}"),
			[]interface{}{7.89},
			nil,
		},
		"[]interface{}: [7, 8, 9]": {
			[]byte("a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}"),
			[]interface{}{int64(7), int64(8), int64(9)},
			nil,
		},
		"[]interface{}: [7.2, 'foo']": {
			[]byte("a:2:{i:0;d:7.2;i:1;s:3:\"foo\";}"),
			[]interface{}{7.2, "foo"},
			nil,
		},
		"[]interface{}: [null]": {
			[]byte("a:1:{i:0;N;}"),
			[]interface{}{nil},
			nil,
		},
		"[]interface{}: [true, false]": {
			[]byte("a:2:{i:0;b:1;i:1;b:0;}"),
			[]interface{}{true, false},
			nil,
		},
		"[]interface{}: [1, 2, 'foo']": {
			[]byte(`a:3:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";}`),
			[]interface{}{int64(1), int64(2), "foo"},
			nil,
		},
		"[]interface{}: [1, 2, 'foo', '中文']": {
			[]byte(`a:4:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;s:6:"中文";}`),
			[]interface{}{int64(1), int64(2), "foo", "中文"},
			nil,
		},
		"[]interface{}: [1, 2, 'foo', '中文', ['a' => 'a']]": {
			[]byte(`a:5:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;s:6:"中文";i:4;a:1:{s:1:"a";s:1:"a";}}`),
			[]interface{}{int64(1), int64(2), "foo", "中文", map[interface{}]interface{}{"a": "a"}},
			nil,
		},
		"[]interface{}: [1, 2, 'foo', ['a' => 'a']]": {
			[]byte(`a:4:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;a:1:{s:1:"a";s:1:"a";}}`),
			[]interface{}{int64(1), int64(2), "foo", map[interface{}]interface{}{"a": "a"}},
			nil,
		},
		"[]interface{}: [1, 2, 'foo', ['a' => 'a'], ['a' => 'a']]": {
			[]byte(`a:5:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;a:1:{s:1:"a";s:1:"a";}i:4;a:1:{s:1:"a";s:1:"a";}}`),
			[]interface{}{int64(1), int64(2), "foo", map[interface{}]interface{}{"a": "a"}, map[interface{}]interface{}{"a": "a"}},
			nil,
		},
		"[]interface{}: [1, 2, 'foo', '中文', ['a' => 'a'], ['a' => 'a']]": {
			[]byte(`a:6:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;s:6:"中文";i:4;a:1:{s:1:"a";s:1:"a";}i:5;a:1:{s:1:"a";s:1:"a";}}`),
			[]interface{}{int64(1), int64(2), "foo", "中文", map[interface{}]interface{}{"a": "a"}, map[interface{}]interface{}{"a": "a"}},
			nil,
		},
		"[]interface{}: [['id'=> '1'], ['id'=> '2']]": {
			[]byte(`a:2:{i:0;a:1:{s:2:"id";s:1:"1";}i:1;a:1:{s:2:"id";s:1:"2";}}`),
			[]interface{}{map[interface{}]interface{}{"id": "1"}, map[interface{}]interface{}{"id": "2"}},
			nil,
		},
		"[]interface{}: [['id'=> '1', 'name' => '1'], ['id'=> '2', 'name' => '2'], ['id'=> '3', 'name' => '3']]": {
			[]byte(`a:3:{i:0;a:2:{s:2:"id";s:1:"1";s:4:"name";s:1:"1";}i:1;a:2:{s:2:"id";s:1:"2";s:4:"name";s:1:"2";}i:2;a:2:{s:2:"id";s:1:"3";s:4:"name";s:1:"3";}}`),
			[]interface{}{map[interface{}]interface{}{"id": "1", "name": "1"}, map[interface{}]interface{}{"id": "2", "name": "2"}, map[interface{}]interface{}{"id": "3", "name": "3"}},
			nil,
		},
		"cannot decode map as slice": {
			[]byte("a:2:{i:0;b:1;i:5;b:0;}"),
			[]interface{}{},
			errors.New("cannot decode map as slice"),
		},
		"not an array": {
			[]byte("N;"),
			[]interface{}{},
			errors.New("not an array"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var result []interface{}
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if len(result) != len(test.output) {
					t.Errorf("Expected %v, got %v", len(result), len(test.output))
				}

				for k, _ := range result {
					if !reflect.DeepEqual(result[k], test.output[k]) {
						t.Errorf("Expected %v (%s), got %v (%s) for #%d",
							result[k], reflect.TypeOf(result[k]).Name(),
							test.output[k], reflect.TypeOf(test.output[k]).Name(), k)
					}
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

func TestUnmarshalAssociativeArray(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        map[interface{}]interface{}
		expectedError error
	}{
		"map[interface{}]interface{}: {'foo': 10, 'bar': 20}": {
			[]byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}"),
			map[interface{}]interface{}{"foo": int64(10), "bar": int64(20)},
			nil,
		},
		"map[interface{}]interface{}: {1: 10, 2: 'foo'}": {
			[]byte("a:2:{i:1;i:10;i:2;s:3:\"foo\";}"),
			map[interface{}]interface{}{int64(1): int64(10), int64(2): "foo"},
			nil,
		},
		"map[interface{}]interface{}: {'foo': 10, 'bar': 20, 'foobar': {'foo': 10, 'bar': 20}}": {
			[]byte(`a:3:{s:3:"foo";i:10;s:3:"bar";i:20;s:6:"foobar";a:2:{s:3:"foo";i:10;s:3:"bar";i:20;}}`),
			map[interface{}]interface{}{"foo": int64(10), "bar": int64(20), "foobar": map[interface{}]interface{}{"foo": int64(10), "bar": int64(20)}},
			nil,
		},
		"map[interface{}]interface{}: {'foo': 10, 'bar': 20, 'foobar': {'foo': 10, 'bar': 20}, 'foobar1': {'foo': 10, 'bar': 20}}": {
			[]byte(`a:4:{s:3:"foo";i:10;s:3:"bar";i:20;s:6:"foobar";a:2:{s:3:"foo";i:10;s:3:"bar";i:20;}s:7:"foobar1";a:2:{s:3:"foo";i:10;s:3:"bar";i:20;}}`),
			map[interface{}]interface{}{"foo": int64(10), "bar": int64(20), "foobar": map[interface{}]interface{}{"foo": int64(10), "bar": int64(20)}, "foobar1": map[interface{}]interface{}{"foo": int64(10), "bar": int64(20)}},
			nil,
		},
		"not an array": {
			[]byte("N;"),
			map[interface{}]interface{}{},
			errors.New("not an array"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			result := make(map[interface{}]interface{})
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if !reflect.DeepEqual(result, test.output) {
					t.Errorf("Expected %v, got %v", result, test.output)
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

var inputNull = []byte("N;")
var inputBoolFalse = []byte("b:0;")
var inputBoolTrue = []byte("b:1;")

func TestUnmarshalWithNull(t *testing.T) {
	result := interface{}(nil)
	err := phpserialize.Unmarshal(inputNull, &result)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestUnmarshalNilWithNull(t *testing.T) {
	err := phpserialize.UnmarshalNil(inputNull)
	if err != nil {
		t.Error(err)
	}
}

func TestBadUnmarshalNilWithNull(t *testing.T) {
	err := phpserialize.UnmarshalNil(inputBoolFalse)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestUnmarshalWithBooleanTrue(t *testing.T) {
	var result bool
	err := phpserialize.Unmarshal(inputBoolTrue, &result)

	if result != true {
		t.Errorf("expected true")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshalObject(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
	var result struct1
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if result.Foo != 10 {
		t.Errorf("Expected %v, got %v", 10, result.Foo)
	}

	if result.Bar.Qux != 1.23 {
		t.Errorf("Expected %v, got %v", 1.23, result.Bar.Qux)
	}

	if result.Baz != "yay" {
		t.Errorf("Expected %v, got %v", "yay", result.Baz)
	}
}

func TestUnmarshalObjectWithTags(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
	var result structTag
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if result.Bar != 10 {
		t.Errorf("Expected %v, got %v", 10, result.Bar)
	}

	if result.Foo.Qux != 1.23 {
		t.Errorf("Expected %v, got %v", 1.23, result.Foo.Qux)
	}

	if result.Balu != "yay" {
		t.Errorf("Expected %v, got %v", "yay", result.Balu)
	}

	if result.Ignored != "" {
		t.Errorf("Expected %v, got %v", "yay", result.Ignored)
	}
}

func TestUnmarshalObjectIntoMap(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
	var result map[interface{}]interface{}
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	expected := map[interface{}]interface{}{
		"baz": "yay",
		"foo": int64(10),
		"bar": map[interface{}]interface{}{
			"qux": 1.23,
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

func TestUnmarshalObjectIntoMapContainingArray(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}s:3:\"baz\";s:3:\"yay\";}"
	var result map[interface{}]interface{}
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	expected := map[interface{}]interface{}{
		"baz": "yay",
		"foo": int64(10),
		"bar": []interface{}{int64(7), int64(8), int64(9)},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

func TestUnmarshalArrayThatContainsObject(t *testing.T) {
	data := "a:3:{i:0;O:7:\"struct1\":2:{s:3:\"foo\";i:10;s:3:\"baz\";s:3:\"yay\";}i:1;i:8;i:2;i:9;}"
	var result []interface{}
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	expected := []interface{}{
		map[interface{}]interface{}{
			"baz": "yay",
			"foo": int64(10),
		},
		int64(8),
		int64(9),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

// https://github.com/elliotchance/phpserialize/issues/7
func TestUnmarshalArrayThatContainsInteger(t *testing.T) {
	data := `a:3:{s:4:"name";s:2:"tw";s:3:"age";i:123;s:4:"wife";a:1:{s:1:"x";s:1:"y";}}`
	var result map[interface{}]interface{}
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	expected := map[interface{}]interface{}{
		"wife": map[interface{}]interface{}{
			"x": "y",
		},
		"name": "tw",
		"age":  int64(123),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

func TestUnmarshalObjectThatContainsArray(t *testing.T) {
	data := "O:7:\"Struct3\":4:{s:11:\"objectArray\";a:2:{i:0;O:7:\"Struct2\":1:{s:3:\"qux\";d:1.1;}i:1;O:7:\"Struct2\":1:{s:3:\"qux\";d:2.2;}}s:8:\"intArray\";a:2:{i:0;i:1;i:1;i:2;}s:10:\"floatArray\";a:2:{i:0;d:1;i:1;d:2;}s:11:\"stringArray\";a:2:{i:0;s:1:\"a\";i:1;s:1:\"b\";}}"
	var result Struct3
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if len(result.ObjectArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.ObjectArray))
	}
	if len(result.IntArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.IntArray))
	}
	if len(result.FloatArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.FloatArray))
	}
	if len(result.StringArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.StringArray))
	}
}

// https://github.com/elliotchance/phpserialize/issues/1
func TestUnmarshalMultibyte(t *testing.T) {
	data := `a:3:{i:0;a:2:{i:0;s:6:"白色";i:1;s:6:"黑色";}i:1;a:3:{i:0;s:3:"大";i:1;s:3:"中";i:2;s:3:"小";}i:2;a:2:{i:0;s:3:"女";i:1;s:3:"男";}}`
	var result map[interface{}]interface{}
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	expected := map[interface{}]interface{}{
		int64(0): []interface{}{"白色", "黑色"},
		int64(1): []interface{}{"大", "中", "小"},
		int64(2): []interface{}{"女", "男"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", expected, result)
	}
}

var escapeTests = map[string]struct {
	Unserialized, Serialized string
}{
	"SingleQuote": {
		"foo'bar", `s:8:"foo\'bar";`,
	},
	"DoubleQuote": {
		"foo\"bar", `s:7:"foo"bar";`,
	},
	"Backslash": {
		"foo\\bar", `s:7:"foo\bar";`,
	},
	"Dollar": {
		"foo$bar", `s:7:"foo$bar";`,
	},
	"NewLine": {
		"foo\nbar", "s:7:\"foo\nbar\";",
	},
	"HorizontalTab": {
		"foo\tbar", "s:7:\"foo\tbar\";",
	},
	"CarriageReturn": {
		"foo\rbar", "s:7:\"foo\rbar\";",
	},
}

func TestUnmarshalEscape(t *testing.T) {
	for testName, test := range escapeTests {
		t.Run(testName, func(t *testing.T) {
			var result string
			err := phpserialize.Unmarshal([]byte(test.Serialized), &result)
			expectErrorToNotHaveOccurred(t, err)

			if test.Unserialized != result {
				t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", test.Unserialized, result)
			}
		})
	}
}
