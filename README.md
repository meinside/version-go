# version-go

Get version strings from runtime.

## Usage

```go
package main

import (
	"fmt"

	"github.com/meinside/version-go"
)

func main() {
	fmt.Printf("full build version of this application: %s\n", version.Full())

	fmt.Printf("build version with some flags: %s\n", version.Build(version.Revision|version.Time|version.Modification))
}
```

## License

MIT

