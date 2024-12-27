package entity

import (
	"bytes"
	"encoding/json"
)

type JSONZipcode Zipcode

func (z JSONZipcode) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	zipcode := Zipcode(z)
	if err := json.NewEncoder(&b).Encode(zipcode); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (z *JSONZipcode) UnmarshalBinary(data []byte) (err error) {
	zipcode := Zipcode(*z)
	b := bytes.NewBuffer(data)
	decoder := json.NewDecoder(b)
	err = decoder.Decode(&zipcode)
	*z = JSONZipcode(zipcode)
	return err
}
