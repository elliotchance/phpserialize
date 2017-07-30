[![Build Status](https://travis-ci.org/elliotchance/phpserialize.svg?branch=master)](https://travis-ci.org/elliotchance/phpserialize)

PHP [serialize()](http://php.net/manual/en/function.serialize.php) and
[unserialize()](http://php.net/manual/en/function.unserialize.php) for Go.

# Install / Update

```bash
go get -u github.com/elliotchance/phpserialize
```

`phpserialize` requires Go 1.8+.

# Example

```go
package main

import (
	"github.com/elliotchance/phpserialize"
	"fmt"
)

func main() {
	out, err := phpserialize.Marshal(3.2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	var in float64
	err = phpserialize.Unmarshal(out, &in)

	fmt.Println(in)
}
```
