package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

func RLPEncode(data interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func RLPDecode(b []byte, out interface{}) {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(out)
	if err != nil {
		log.Fatal(err)
	}
}
