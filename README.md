# md5-Hash 
gomd5hash is a simple Golang package designed to make it easy to create MD5 hashes. It provides a straightforward function to generate MD5 hashes for any given input text.

### Installation
To use gomd5hash, you need to have Golang installed on your machine. You can install the package using the following go get command:
```
go get github.com/yourusername/gomd5hash
```

### Usage
Import the package in your Golang code:
```
import (
    "fmt"
    "github.com/codeArtisanry/gomd5hash"
)
```

---

To create an MD5 hash, use the CreateMd5Hash function:
```
text := "Hello, gomd5hash!"
hash := gomd5hash.CreateMd5Hash(text)
fmt.Println("MD5 Hash:", hash)
```

---
### Example
Here is a simple example demonstrating how to use gomd5hash:

```
package main

import (
	"fmt"
	"github.com/yourusername/gomd5hash"
)

func main() {
	text := "Hello, gomd5hash!"
	hash := gomd5hash.CreateMd5Hash(text)
	fmt.Println("Input Text:", text)
	fmt.Println("MD5 Hash:", hash)
}
```
Output:

`
Input Text: Hello, gomd5hash!
MD5 Hash: 79054025255fb1a26e4bc422aef54eb4
`

### Contributing
If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request. Contributions are welcome!

