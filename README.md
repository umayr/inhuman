## Inhuman [![Build Status](https://travis-ci.org/umayr/inhuman.svg?branch=master)](https://travis-ci.org/umayr/inhuman)
> tiny package to dehumanize humanized timestamps. 

### Install

```
  λ go get -u github.com/umayr/inhuman
```

### Usage

```go
package main

import "github.com/umayr/inhuman"

func main() {
	inhuman.Parse("1h 3m") // 3780000000000, nil

	// try it with some spaces
	inhuman.Parse("1 h 3 m") // 3780000000000, nil

	// what about no spaces?
	inhuman.Parse("1h3m") // 3780000000000, nil

	// supported formats
	inhuman.Parse("1h 3m 17s")                 // 3797000000000, nil
	inhuman.Parse("1hr 3min 17sec")            // 3797000000000, nil
	inhuman.Parse("1hrs 3mins 17secs")         // 3797000000000, nil
	inhuman.Parse("1hour 3minute 17second")    // 3797000000000, nil
	inhuman.Parse("1hours 3minutes 17seconds") // 3797000000000, nil
}

```

### License

[MIT](https://github.com/umayr/inhuman/blob/master/LICENSE)
