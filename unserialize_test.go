package phpserialize_test

import (
	"errors"
	"github.com/elliotchance/phpserialize"
	"testing"
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

//func decodeFloat(input []byte, output float64, expectedError error) {
//	{
//		var result float32
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			expectErrorToNotHaveOccurred(t, err)
//			Expect(result).To(Equal(float32(output)))
//		} else {
//			expectErrorToNotHaveOccurred(t, err)
//			expectErrorToEqual(t, err, test.expectedError)
//		}
//	}
//
//	{
//		var result float64
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			expectErrorToNotHaveOccurred(t, err)
//			if result != test.output {
//						t.Errorf("Expected '%v', got '%v'", result, test.output)
//					}
//		} else {
//			expectErrorToNotHaveOccurred(t, err)
//			expectErrorToEqual(t, err, test.expectedError)
//		}
//	}
//}
//
//func decodeString(input []byte, output string, expectedError error) {
//	var result string
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		expectErrorToNotHaveOccurred(t, err)
//		if result != test.output {
//						t.Errorf("Expected '%v', got '%v'", result, test.output)
//					}
//	} else {
//		expectErrorToNotHaveOccurred(t, err)
//		expectErrorToEqual(t, err, test.expectedError)
//	}
//}
//
//func decodeBinary(input []byte, output []byte, expectedError error) {
//	var result []byte
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		expectErrorToNotHaveOccurred(t, err)
//		if result != test.output {
//						t.Errorf("Expected '%v', got '%v'", result, test.output)
//					}
//	} else {
//		expectErrorToNotHaveOccurred(t, err)
//		expectErrorToEqual(t, err, test.expectedError)
//	}
//}
//
//func decodeArray(input []byte, output []interface{}, expectedError error) {
//	var result []interface{}
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		expectErrorToNotHaveOccurred(t, err)
//		Expect(len(result)).To(Equal(len(output)))
//		for k, _ := range result {
//			// Ginkgo has a safety feature when comparing two nil
//			// values for equality. You have to use the nil
//			// assertions. I this case it's not important.
//			if result[k] != nil || output[k] != nil {
//				Expect(result[k]).To(BeEquivalentTo(output[k]))
//			}
//		}
//	} else {
//		expectErrorToNotHaveOccurred(t, err)
//		expectErrorToEqual(t, err, test.expectedError)
//	}
//}
//
//func decodeAssociativeArray(input []byte, output map[interface{}]interface{}, expectedError error) {
//	result := make(map[interface{}]interface{})
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		expectErrorToNotHaveOccurred(t, err)
//		Expect(reflect.DeepEqual(result, output)).To(BeTrue())
//	} else {
//		expectErrorToNotHaveOccurred(t, err)
//		expectErrorToEqual(t, err, test.expectedError)
//	}
//}

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

//
//var _ = Describe("phpserialize", func() {
//	Describe("Unmarshal - unserialize()", func() {
//		DescribeTable("decode int",
//			decodeInt,
//
//			Entry("0", []byte("i:0;"), 0, nil),
//			Entry("5", []byte("i:5;"), 5, nil),
//			Entry("-8", []byte("i:-8;"), -8, nil),
//			Entry("1000000", []byte("i:1000000;"), 1000000, nil),
//
//			Entry("not an integer", []byte("N;"), 0, errors.New("not an integer")),
//		)
//
//		DescribeTable("decode float",
//			decodeFloat,
//
//			Entry("3.2", []byte("d:3.2;"), 3.2, nil),
//			Entry("10.0", []byte("d:10;"), 10.0, nil),
//			Entry("123.456789", []byte("d:123.456789;"), 123.456789, nil),
//			Entry("1.23e9", []byte("d:1230000000;"), 1.23e9, nil),
//			Entry("-17.23", []byte("d:3.2;"), 3.2, nil),
//
//			Entry("not a float", []byte("N;"), 0.0, errors.New("not a float")),
//		)
//
//		DescribeTable("decode string",
//			decodeString,
//
//			Entry("''", []byte("s:0:\"\";"), "", nil),
//			Entry("'Hello world'", []byte("s:11:\"Hello world\";"),
//				"Hello world", nil),
//			Entry("'Björk Guðmundsdóttir'",
//				[]byte("s:23:\"Bj\\xc3\\xb6rk Gu\\xc3\\xb0mundsd\\xc3\\xb3ttir\";"),
//				"Björk Guðmundsdóttir", nil),
//
//			Entry("not a string", []byte("N;"), "", errors.New("not a string")),
//		)
//
//		DescribeTable("decode binary",
//			decodeBinary,
//
//			Entry("[]byte: \\001\\002\\003", []byte("s:3:\"\\x01\\x02\\x03\";"),
//				[]byte{1, 2, 3}, nil),
//
//			Entry("not a string", []byte("N;"), []byte{}, errors.New("not a string")),
//		)
//
//		DescribeTable("decode array (slice)",
//			decodeArray,
//
//			Entry("[]interface{}: [7.89]", []byte("a:1:{i:0;d:7.89;}"),
//				[]interface{}{7.89}, nil),
//			Entry("[]interface{}: [7, 8, 9]",
//				[]byte("a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}"),
//				[]interface{}{7, 8, 9}, nil),
//			Entry("[]interface{}: [7.2, 'foo']",
//				[]byte("a:2:{i:0;d:7.2;i:1;s:3:\"foo\";}"),
//				[]interface{}{7.2, "foo"}, nil),
//			Entry("[]interface{}: [null]",
//				[]byte("a:1:{i:0;N;}"),
//				[]interface{}{nil}, nil),
//			Entry("[]interface{}: [true, false]",
//				[]byte("a:2:{i:0;b:1;i:1;b:0;}"),
//				[]interface{}{true, false}, nil),
//
//			Entry("cannot decode map as slice",
//				[]byte("a:2:{i:0;b:1;i:5;b:0;}"),
//				[]interface{}{},
//				errors.New("cannot decode map as slice")),
//			Entry("not an array", []byte("N;"), []interface{}{},
//				errors.New("not an array")),
//		)
//
//		DescribeTable("decode associative array (map)",
//			decodeAssociativeArray,
//
//			Entry("map[interface{}]interface{}: {'foo': 10, 'bar': 20}",
//				[]byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}"),
//				map[interface{}]interface{}{"foo": int64(10), "bar": int64(20)},
//				nil),
//			Entry("map[interface{}]interface{}: {1: 10, 2: 'foo'}",
//				[]byte("a:2:{i:1;i:10;i:2;s:3:\"foo\";}"),
//				map[interface{}]interface{}{int64(1): int64(10), int64(2): "foo"},
//				nil),
//
//			Entry("not an array", []byte("N;"),
//				map[interface{}]interface{}{},
//				errors.New("not an array")),
//		)
//
//		Describe("decode object", func() {
//			It("struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}", func() {
//				data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
//				var result struct1
//				err := phpserialize.Unmarshal([]byte(data), &result)
//				expectErrorToNotHaveOccurred(t, err)
//
//				Expect(result.Foo).To(Equal(10))
//				Expect(result.Bar.Qux).To(Equal(1.23))
//				Expect(result.Baz).To(Equal("yay"))
//			})
//		})
//	})
//})
