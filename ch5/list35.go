package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))

	_, _ = fmt.Fprintf(&b, " World!")

	_, _ = io.Copy(os.Stdout, &b)
}
