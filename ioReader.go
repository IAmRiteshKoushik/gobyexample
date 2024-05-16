package main

import(
    "bytes"
    "io"
)

func ioReaderDriver() {
    rawMsg := []byte("rawbytes")
    handleRawBytes(bytes.NewReader(rawMsg))
}

// Why do we need to use an io.Reader instead of []byte
func handleRawBytes(r io.Reader) {

}
