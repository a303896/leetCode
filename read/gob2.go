package read

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type string
	City string
	Country string
}

type VCard struct {
	FirstName string
	LastName string
	Addresses []*Address
	Remark string
}

var content string
var vcard VCard

func Encode() {
	pa := &Address{"private", "Aartselaar","Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}
	file, _ := os.OpenFile(DIR + "vcard.gob", os.O_CREATE | os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Fatal("encode error:", err)
	}
}

func Decode() {
	file, _ := os.Open(DIR + "vcard.gob")
	dec := gob.NewDecoder(file)
	err := dec.Decode(&vcard)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(vcard)
}