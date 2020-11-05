package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
The Time function implements the interface for Binary, Gob and JSON serialization.
The JSON format is considered to be very universal, so an example on how the value
is serialized to JSON is shown. Notoe that the Time function serializes the value
in the manner of RFC 3339 which proposes a so-called internet data/timeformat.

Another very universal way to serialize/keep the time is to use the epoch time.
The epoch time is independent if timezones because it is defined by second/nanosecond
elapsed since an absolute point in time. Finally, it is represented as a number so
there is no reason to serialize and deserialize the value.
*/
func main() {
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2017, 11, 20, 11, 20, 10, 0, eur)

	// json.Marshaler interface
	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized as RFC 3339:", string(b))
	t2 := time.Time{}
	t2.UnmarshalJSON(b)
	fmt.Println("Deserialized from RFC 3339:", t2)

	// Serialize as epoch
	epoch := t.Unix()
	fmt.Println("Serialized as Epoch:", epoch)

	// Deserialize epoch
	jsonStr := fmt.Sprintf("{\"created\":%d}", epoch)
	data := struct {
		Created int64 `json:"created"`
	}{}
	json.Unmarshal([]byte(jsonStr), &data)
	deserialied := time.Unix(data.Created, 0)
	fmt.Println("Deserialized from Epoch:", deserialied)
}
