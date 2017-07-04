package phpserialize_test

import (
	"github.com/elliotchance/phpserialize"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

type struct1 struct {
	Foo    int
	Bar    Struct2
	hidden bool
	Baz    string
}

type Struct2 struct {
	Qux float64
}

func testMarshal(input interface{}, output []byte, options *phpserialize.MarshalOptions) {
	if options == nil {
		options = phpserialize.DefaultMarshalOptions()
	}

	result, err := phpserialize.Marshal(input, options)
	Expect(err).ToNot(HaveOccurred())
	Expect(result).To(Equal(output))
}

func testMarshalFail(input interface{}, errorMessage string) {
	result, err := phpserialize.Marshal(input, phpserialize.DefaultMarshalOptions())
	Expect(err).To(HaveOccurred())
	Expect(result).To(BeNil())
	Expect(err.Error()).To(Equal(errorMessage))
}

// These tests have been adpated from the wonderful work at:
// https://github.com/mitsuhiko/phpserialize/blob/master/tests.py
var _ = Describe("phpserialize", func() {
	Describe("Marshal - serialize()", func() {
		DescribeTable("encode null",
			testMarshal,

			Entry("null", nil, []byte("N;"), nil),
		)

		DescribeTable("encode bool",
			testMarshal,

			Entry("true", true, []byte("b:1;"), nil),
			Entry("false", false, []byte("b:0;"), nil),
		)

		DescribeTable("encode int",
			testMarshal,

			Entry("int: 0", 0, []byte("i:0;"), nil),
			Entry("int: 5", 5, []byte("i:5;"), nil),
			Entry("int: -8", -8, []byte("i:-8;"), nil),

			Entry("int8: 20", int8(20), []byte("i:20;"), nil),
			Entry("int16: 22", int16(22), []byte("i:22;"), nil),
			Entry("int32: 27", int32(27), []byte("i:27;"), nil),
			Entry("int64: 28", int64(28), []byte("i:28;"), nil),

			Entry("uint8: 4", uint8(4), []byte("i:4;"), nil),
			Entry("uint16: 7", uint16(7), []byte("i:7;"), nil),
			Entry("uint32: 9", uint32(9), []byte("i:9;"), nil),
			Entry("uint64: 11", uint64(11), []byte("i:11;"), nil),
		)

		DescribeTable("encode float",
			testMarshal,

			Entry("float64: 3.2", 3.2, []byte("d:3.2;"), nil),
			Entry("float64: 10.0", 10.0, []byte("d:10;"), nil),
			Entry("float64: 123.456789", 123.456789, []byte("d:123.456789;"), nil),
			Entry("float64: 1.23e9", 1.23e9, []byte("d:1230000000;"), nil),
			Entry("float64: -17.23", 3.2, []byte("d:3.2;"), nil),

			Entry("float32: 4.8", float32(4.8), []byte("d:4.8;"), nil),
		)

		DescribeTable("encode string",
			testMarshal,

			Entry("string: ''", "", []byte("s:0:\"\";"), nil),
			Entry("string: 'Hello world'", "Hello world",
				[]byte("s:11:\"Hello world\";"), nil),
			Entry("string: 'Björk Guðmundsdóttir'", "Björk Guðmundsdóttir",
				[]byte("s:23:\"Bj\\xc3\\xb6rk Gu\\xc3\\xb0mundsd\\xc3\\xb3ttir\";"), nil),
		)

		DescribeTable("encode binary",
			testMarshal,

			Entry("[]byte: \\001\\002\\003", []byte{1, 2, 3},
				[]byte("s:3:\"\\x01\\x02\\x03\";"), nil),
		)

		DescribeTable("encode array (slice)",
			testMarshal,

			Entry("[]float64: [7.89]", []float64{7.89},
				[]byte("a:1:{i:0;d:7.89;}"), nil),
			Entry("[]int: [7, 8, 9]", []int{7, 8, 9},
				[]byte("a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}"), nil),
			Entry("[]interface{}: [7.2, 'foo']", []interface{}{7.2, "foo"},
				[]byte("a:2:{i:0;d:7.2;i:1;s:3:\"foo\";}"), nil),
		)

		DescribeTable("encode associative array (map)",
			testMarshal,

			Entry("map[string]int: {'foo': 10, 'bar': 20}",
				map[string]int{"foo": 10, "bar": 20},
				[]byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}"), nil),
			Entry("map[int]interface{}: {1: 10, 2: 'foo'}",
				map[int]interface{}{1: 10, 2: "foo"},
				[]byte("a:2:{i:1;i:10;i:2;s:3:\"foo\";}"), nil),
		)

		stdClassOnly := phpserialize.DefaultMarshalOptions()
		stdClassOnly.OnlyStdClass = true

		DescribeTable("encode object (struct)",
			testMarshal,

			Entry("struct1{Foo int, Bar Struct2{Qux float64}, hidden bool, Bar string}",
				struct1{10, Struct2{1.23}, true, "yay"},
				[]byte("O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"),
				nil),
			Entry("&struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}",
				&struct1{20, Struct2{7.89}, false, "yay"},
				[]byte("O:7:\"struct1\":3:{s:3:\"foo\";i:20;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:7.89;}s:3:\"baz\";s:3:\"yay\";}"),
				nil),

			// stdClassOnly
			Entry("struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}: OnlyStdClass = true",
				struct1{10, Struct2{1.23}, true, "yay"},
				[]byte("O:8:\"stdClass\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:8:\"stdClass\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"),
				stdClassOnly),
			Entry("&struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}: OnlyStdClass = true",
				&struct1{20, Struct2{7.89}, false, "yay"},
				[]byte("O:8:\"stdClass\":3:{s:3:\"foo\";i:20;s:3:\"bar\";O:8:\"stdClass\":1:{s:3:\"qux\";d:7.89;}s:3:\"baz\";s:3:\"yay\";}"),
				stdClassOnly),
		)

		DescribeTable("can not encode",
			testMarshalFail,

			Entry("uintptr", uintptr(13), "can not encode: uintptr"),
		)
	})
})
