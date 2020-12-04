package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
The gob serialization and deserialization need the Encoder and Decoder.
The gob.NewEncoder function creates the Encoder with the underlying Writer.
Each call of the Encode method itself is the self-describing binary format.
This means each serialized struct is preceded by its description.

To decode the data from the serialized form, the Decoder must be created by
calling the gob.NewDecoder with the underlying Reader. The Decode then accepts
the pointer to the structure where the data should be deserialized.
*/

// User basic model
type User struct {
	FirstName string
	LastName  string
	Age       int
	Active    bool
}

func (u User) String() string {
	return fmt.Sprintf(`{"FirstName":%s, "LastName":%s, "Age":%d, "Active":%v}`, u.FirstName, u.LastName, u.Age, u.Active)
}

// SimpleUser basic model
type SimpleUser struct {
	FirstName string
	LastName  string
}

func (u SimpleUser) String() string {
	return fmt.Sprintf(`{"FirstName":%s, "LastName":%s}`, u.FirstName, u.LastName)
}

func main() {
	var buff bytes.Buffer

	// Encode value
	enc := gob.NewEncoder(&buff)
	user := User{
		"Radomir",
		"Sohlich",
		30,
		true,
	}
	enc.Encode(user)
	fmt.Printf("%X\n", buff.Bytes())

	// Decode value
	out := User{}
	dec := gob.NewDecoder(&buff)
	dec.Decode(&out)
	fmt.Println(out.String())

	enc.Encode(user)
	out2 := SimpleUser{}
	dec.Decode(&out2)
	fmt.Println(out2.String())
}
