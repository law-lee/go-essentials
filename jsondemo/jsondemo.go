package jsondemo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Student struct {
	Name     string
	Standard int `json:"Standard"`
}

func decodeFromReader(r io.Reader) ([]*Student, error) {
	var res []*Student

	dec := json.NewDecoder(r)
	err := dec.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func decodeFromString(s string) ([]*Student, error) {
	r := bytes.NewBufferString(s)
	return decodeFromReader(r)
}

func decodeFromFile(path string) ([]*Student, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decodeFromReader(f)
}

func Run() {
	// Structs are serialized as JSON dictionaries. By default dictionary keys are the same as struct field names.
	type Person struct {
		fullName string
		Name     string
		Age      int    `json:"age"`
		City     string `json:"city"`
	}

	p := Person{
		fullName: "John Stanson",
		Name:     "John",
		Age:      37,
		City:     "SF",
	}
	d, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in compact JSON: %s\n", string(d))

	d, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in pretty-printed JSON:\n%s\n", string(d))

	// parsing arbitrary JSON documents
	var jsonStr = `{
		"name": "Jane",
		"age": 24,
		"city": "ny"
	}`

	// For arbitrary JSON documents we can decode into a map[string]interface{}, which can represent all valid JSON documents.
	var doc map[string]any
	err = json.Unmarshal([]byte(jsonStr), &doc)
	if err != nil {
		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("doc: %#v\n", doc)
	name, ok := doc["name"].(string)
	if !ok {
		log.Fatalf("doc has no key 'name' or its value is not string\n")
	}
	fmt.Printf("name: %#v\n", name)

	// decoding into anonymous structs
	var jsonBlob = []byte(`
{
  "_total": 1,
  "_links": {
	"self": "https://api.twitch.tv/kraken/channels/foo/subscriptions?direction=ASC&limit=25&offset=0",
	"next": "https://api.twitch.tv/kraken/channels/foo/subscriptions?direction=ASC&limit=25&offset=25"
  },
  "subscriptions": [
	{
	  "created_at": "2011-11-23T02:53:17Z",
	  "_id": "abcdef0000000000000000000000000000000000",
	  "_links": {
		"self": "https://api.twitch.tv/kraken/channels/foo/subscriptions/bar"
	  },
	  "user": {
		"display_name": "bar",
		"_id": 123456,
		"name": "bar",
		"created_at": "2011-06-16T18:23:11Z",
		"updated_at": "2014-10-23T02:20:51Z",
		"_links": {
		  "self": "https://api.twitch.tv/kraken/users/bar"
		}
	  }
	}
  ]
}
`)

	var js struct {
		Total int `json:"_total"`
		Links struct {
			Next string `json:"next"`
		} `json:"_links"`
		Subs []struct {
			Created string `json:"created_at"`
			User    struct {
				Name string `json:"name"`
				ID   int    `json:"_id"`
			} `json:"user"`
		} `json:"subscriptions"`
	}

	err = json.Unmarshal(jsonBlob, &js)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", js)

	// decode json from file
	r, err := decodeFromFile(filepath.Join("jsondemo", "s.json"))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(r)

	// custom json marshal
	type Event struct {
		What string
		When time.Time
	}
	e := Event{
		What: "earthquake",
		When: time.Now(),
	}
	d, err = json.Marshal(&e)
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Standard time JSON: %s\n", string(d))

	type Event2 struct {
		What string
		When customTime
	}

	e2 := Event2{
		What: "earthquake",
		When: customTime(time.Now()),
	}
	d, err = json.Marshal(&e2)
	if err != nil {
		log.Fatalf("json.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("\nCustom time JSON: %s\n", string(d))
	var decoded Event2
	err = json.Unmarshal(d, &decoded)
	if err != nil {
		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	}
	t := time.Time(decoded.When)
	fmt.Printf("Decoded custom time: %s\n", t.Format(customTimeFormat))
}

type customTime time.Time

const customTimeFormat = `"2006-01-02"`

func (ct customTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	s := t.Format(customTimeFormat)
	return []byte(s), nil
}

func (ct *customTime) UnmarshalJSON(d []byte) error {
	t, err := time.Parse(customTimeFormat, string(d))
	if err != nil {
		return err
	}
	*ct = customTime(t)
	return nil
}
