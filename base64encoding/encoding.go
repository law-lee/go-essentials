package base64encoding

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func Run() {
	// hex encode and decode
	d := []byte{0x01, 0xff, 0x3a, 0xcd}
	s := hex.EncodeToString(d)
	fmt.Printf("Hex: %s\n", s)

	d2, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalf("hex.DecodeString() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, d2) {
		log.Fatalf("decoded version is different than original")
	}

	// hex encoding with fmt.Sprintf
	d = []byte{0x01, 0xff, 0x3a, 0xcd}
	s = fmt.Sprintf("%x", d)
	fmt.Printf("Hex: %s\n", s)

	var decoded []byte
	_, err = fmt.Sscanf(s, "%x", &decoded)
	if err != nil {
		log.Fatalf("fmt.Sscanf() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, decoded) {
		log.Fatalf("decoded version is different than original")
	}

	n := 3824
	fmt.Printf("%d in hex is 0x%x\n", n, n)

	// hex encoding to writer, decoding from reader
	d = []byte{0x01, 0xff, 0x3a, 0xcd}

	writer := &bytes.Buffer{}
	hexWriter := hex.NewEncoder(writer)

	_, err = hexWriter.Write(d)
	if err != nil {
		log.Fatalf("hexWriter.Write() failed with '%s'\n", err)
	}

	encoded := writer.Bytes()
	fmt.Printf("Hex: %s\n", string(encoded))

	reader := bytes.NewBuffer(encoded)
	hexReader := hex.NewDecoder(reader)

	decoded, err = io.ReadAll(hexReader)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	if !bytes.Equal(d, decoded) {
		log.Fatalf("decoded version is different than original")
	}

	// base64 encode and decode
	d = []byte{0x01, 0xff, 0x3a, 0xcd}
	s = base64.StdEncoding.EncodeToString(d)
	fmt.Printf("base64: %s\n", s)

	d2, err = base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("hex.DecodeString() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, d2) {
		log.Fatalf("decoded version is different than original")
	}

	// url safe base64
	d = []byte{0x01, 0xff, 0x3a, 0xcd}
	s = base64.URLEncoding.EncodeToString(d)
	fmt.Printf("base64: %s\n", s)

	d2, err = base64.URLEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("hex.DecodeString() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, d2) {
		log.Fatalf("decoded version is different than original")
	}
}
