package entity

import (
	"bytes"
	"encoding/gob"
)

type GOBZipcode Zipcode

func (z GOBZipcode) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	zipcode := Zipcode(z)
	if err := gob.NewEncoder(&b).Encode(zipcode); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (z *GOBZipcode) UnmarshalBinary(data []byte) (err error) {
	zipcode := Zipcode(*z)
	b := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(b)
	err = decoder.Decode(&zipcode)
	*z = GOBZipcode(zipcode)
	return err
}
