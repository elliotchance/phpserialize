package phpserialize_test

import (
	"github.com/elliotchance/phpserialize"
	"reflect"
	"testing"
)

type struct1 struct {
	Foo    int
	Bar    Struct2
	hidden bool
	Baz    string
}

type structTag struct {
	Foo     Struct2 `php:"bar"`
	Bar     int     `php:"foo"`
	hidden  bool
	Balu    string   `php:"baz"`
	Ignored string   `php:"-"`
	Nilptr  *Struct2 `php:",omitnilptr"`
}

type Struct2 struct {
	Qux float64
}

type Struct3 struct {
	ObjectArray []Struct2
	IntArray    []int64
	FloatArray  []float64
	StringArray []string
}

type marshalTest struct {
	input   interface{}
	output  []byte
	options *phpserialize.MarshalOptions
}

func getStdClassOnly() *phpserialize.MarshalOptions {
	stdClassOnly := phpserialize.DefaultMarshalOptions()
	stdClassOnly.OnlyStdClass = true

	return stdClassOnly
}

// These tests have been adapted from the wonderful work at:
// https://github.com/mitsuhiko/phpserialize/blob/master/tests.py
var marshalTests = map[string]marshalTest{
	// encode null
	"null": {nil, []byte("N;"), nil},

	// encode bool
	"true":  {true, []byte("b:1;"), nil},
	"false": {false, []byte("b:0;"), nil},

	// encode int
	"int: 0":  {0, []byte("i:0;"), nil},
	"int: 5":  {5, []byte("i:5;"), nil},
	"int: -8": {-8, []byte("i:-8;"), nil},

	"int8: 20":  {int8(20), []byte("i:20;"), nil},
	"int16: 22": {int16(22), []byte("i:22;"), nil},
	"int32: 27": {int32(27), []byte("i:27;"), nil},
	"int64: 28": {int64(28), []byte("i:28;"), nil},

	"uint: 3":    {uint(3), []byte("i:3;"), nil},
	"uint8: 4":   {uint8(4), []byte("i:4;"), nil},
	"uint16: 7":  {uint16(7), []byte("i:7;"), nil},
	"uint32: 9":  {uint32(9), []byte("i:9;"), nil},
	"uint64: 11": {uint64(11), []byte("i:11;"), nil},

	// encode float
	"float64: 3.2":        {3.2, []byte("d:3.2;"), nil},
	"float64: 10.0":       {10.0, []byte("d:10;"), nil},
	"float64: 123.456789": {123.456789, []byte("d:123.456789;"), nil},
	"float64: 1.23e9":     {1.23e9, []byte("d:1230000000;"), nil},
	"float64: -17.23":     {3.2, []byte("d:3.2;"), nil},

	"float32: 4.8": {float32(4.8), []byte("d:4.8;"), nil},

	// encode string
	"string: ''": {"", []byte("s:0:\"\";"), nil},
	"string: 'Hello world'": {
		"Hello world",
		[]byte("s:11:\"Hello world\";"),
		nil,
	},
	"string: 'Björk Guðmundsdóttir'": {
		"Björk Guðmundsdóttir",
		[]byte("s:23:\"Björk Guðmundsdóttir\";"),
		nil,
	},

	// encode binary
	"[]byte: \\001\\002\\003": {
		[]byte{1, 2, 3},
		[]byte("s:3:\"\\x01\\x02\\x03\";"),
		nil,
	},

	// encode array (slice)
	"[]float64: [7.89]": {
		[]float64{7.89},
		[]byte("a:1:{i:0;d:7.89;}"),
		nil,
	},
	"[]int: [7, 8, 9]": {
		[]int{7, 8, 9},
		[]byte("a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}"),
		nil,
	},
	"[]interface{}: [7.2, 'foo']": {
		[]interface{}{7.2, "foo"},
		[]byte("a:2:{i:0;d:7.2;i:1;s:3:\"foo\";}"),
		nil,
	},

	// encode associative array (map)
	"map[string]int: {'foo': 10, 'bar': 20}": {
		map[string]int{"foo": 10, "bar": 20},
		[]byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}"),
		nil,
	},
	"map[int]interface{}: {1: 10, 2: 'foo'}": {
		map[int]interface{}{1: 10, 2: "foo"},
		[]byte("a:2:{i:1;i:10;i:2;s:3:\"foo\";}"),
		nil,
	},

	// encode object (struct)
	"struct1{Foo int, Bar Struct2{Qux float64}, hidden bool, Bar string}": {
		struct1{10, Struct2{1.23}, true, "yay"},
		[]byte("O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"),
		nil,
	},
	"&struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}": {
		&struct1{20, Struct2{7.89}, false, "yay"},
		[]byte("O:7:\"struct1\":3:{s:3:\"foo\";i:20;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:7.89;}s:3:\"baz\";s:3:\"yay\";}"),
		nil,
	},

	// encode object with array of objects
	"struct3{ObjectArray Struct2{Qux float64}, IntArray {1, 2}, FloatArray {1.0, 2.0}, StringArray {'a', 'b'}}": {
		Struct3{[]Struct2{{1.1}, {2.2}}, []int64{1, 2}, []float64{1.0, 2.0}, []string{"a", "b"}},
		[]byte("O:7:\"Struct3\":4:{s:11:\"objectArray\";a:2:{i:0;O:7:\"Struct2\":1:{s:3:\"qux\";d:1.1;}i:1;O:7:\"Struct2\":1:{s:3:\"qux\";d:2.2;}}s:8:\"intArray\";a:2:{i:0;i:1;i:1;i:2;}s:10:\"floatArray\";a:2:{i:0;d:1;i:1;d:2;}s:11:\"stringArray\";a:2:{i:0;s:1:\"a\";i:1;s:1:\"b\";}}"),
		nil,
	},

	// encode object (struct with tags)
	"structTag{Bar int, Foo Struct2{Qux float64}, hidden bool, Balu string, Nilptr <nil>}": {
		structTag{Struct2{1.23}, 10, true, "yay", "", nil},
		[]byte("O:9:\"structTag\":3:{s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"foo\";i:10;s:3:\"baz\";s:3:\"yay\";}"),
		nil,
	},

	// stdClassOnly
	"struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}: OnlyStdClass = true": {
		struct1{10, Struct2{1.23}, true, "yay"},
		[]byte("O:8:\"stdClass\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:8:\"stdClass\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"),
		getStdClassOnly(),
	},
	"&struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}: OnlyStdClass = true": {
		&struct1{20, Struct2{7.89}, false, "yay"},
		[]byte("O:8:\"stdClass\":3:{s:3:\"foo\";i:20;s:3:\"bar\";O:8:\"stdClass\":1:{s:3:\"qux\";d:7.89;}s:3:\"baz\";s:3:\"yay\";}"),
		getStdClassOnly(),
	},
}

func TestMarshal(t *testing.T) {
	for testName, test := range marshalTests {
		t.Run(testName, func(t *testing.T) {
			if test.options == nil {
				test.options = phpserialize.DefaultMarshalOptions()
			}

			result, err := phpserialize.Marshal(test.input, test.options)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected '%v', got '%v'", string(test.output),
					string(result))
			}
		})
	}
}

func TestMarshalFail(t *testing.T) {
	options := phpserialize.DefaultMarshalOptions()
	result, err := phpserialize.Marshal(uintptr(13), options)
	if err == nil {
		t.Error("expected error to occur")
	}
	if result != nil {
		t.Error("result was not nil")
	}
	if err.Error() != "can not encode: uintptr" {
		t.Error(err.Error())
	}
}

func TestMarshalEscape(t *testing.T) {
	for testName, test := range escapeTests {
		t.Run(testName, func(t *testing.T) {
			options := phpserialize.DefaultMarshalOptions()
			result, err := phpserialize.Marshal(test.Unserialized, options)
			expectErrorToNotHaveOccurred(t, err)

			if test.Serialized != string(result) {
				t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", test.Serialized, result)
			}
		})
	}
}
