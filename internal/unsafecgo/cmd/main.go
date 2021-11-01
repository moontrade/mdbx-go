package main

import (
	"github.com/moontrade/mdbx-go/internal/unsafecgo"
)

func main() {
	//cgo.CGO()
	unsafecgo.NonBlocking((*byte)(nil), 0, 0)
}
