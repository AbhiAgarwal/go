package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

// A common use for crc32 is to compare two files.
// If the Sum32 value for both files is the same, it's highly likely (though not 100% certain) that the files are the same
// If the values are different then the files are definitely not the same
func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	// The crc32 hash object implements the Writer interface, so we can write bytes to it like any other Writer
	h := crc32.NewIEEE()
	h.Write(bs)

	// Once we've written everything we want we call Sum32() to return a uint32
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("../resources/readme.txt")
	if err != nil {
		return
	}
	h2, err := getHash("../resources/writtenREADME.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
