# gullwidth

Create fullwidth text from ASCII input. For ＡＥＳＴＨＥＴＩＣ purposes ＯＮＬＹ.

## Usage

### CLI

#### Building

```bash
git clone https://github.com/clukawski/gullwidth
cd gullwidth/cmd/gullwidth
go build
```

#### Creating fullwidth text

```bash
./gullwidth your text here
ｙｏｕｒ　ｔｅｘｔ　ｈｅｒｅ
```

### Library

See `cmd/gullwidth/main.go`:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clukawski/gullwidth"
)

func main() {
	input := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	fullwidth, err := gullwidth.Fullwidth(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullwidth)
}
```
