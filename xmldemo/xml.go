package xmldemo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

// https://www.programming-books.io/essential/go/easy-generation-of-xml-struct-definition-8b225356aaa9406990d279868a1cd548
func Run() {
	var xmlStr = `
<people>
	<person age="34">
		<first-name>John</first-name>
		<address>
			<city>San Francisco</city>
			<state>CA</state>
		</address>
	</person>

	<person age="23">
		<first-name>Julia</first-name>
	</person>
</people>`

	// This is useful when you need to know if a value was present in XML or not. If we used string for City field,
	// we wouldn’t know if empty string means:
	//
	//	city element was not present in XML
	//	element was present but had empty value (<city></city>)
	type Address struct {
		City  *string `xml:"city"`
		State string  `xml:"state"`
	}
	// In xml struct tags element is the default. To switch to attribute, add ,attr to struct xml tag as done for Age.
	type Person struct {
		Age       int     `xml:"age,attr"`
		FirstName string  `xml:"first-name"`
		Address   Address `xml:"address"`
	}
	type People struct {
		Person []Person `xml:"person"`
	}

	var people People
	data := []byte(xmlStr)
	err := xml.Unmarshal(data, &people)
	if err != nil {
		log.Fatalf("xml.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("%#v\n\n", people)
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

type Person struct {
	Age       int     `xml:"age,attr"`
	FirstName string  `xml:"first-name"`
	Address   Address `xml:"address"`
}

type People struct {
	XMLName        xml.Name `xml:"people"`
	Person         []Person `xml:"person"`
	noteSerialized int
}

func Run2() {
	people := People{
		Person: []Person{
			Person{
				Age:       34,
				FirstName: "John",
				Address: Address{
					City:  "San Francisco",
					State: "CA",
				},
			},
		},
		noteSerialized: 8,
	}
	d, err := xml.Marshal(&people)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("Compact XML: %s\n\n", string(d))

	d, err = xml.MarshalIndent(&people, "", "  ")
	if err != nil {
		log.Fatalf("xml.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Pretty printed XML:\n%s\n", string(d))
}

func Run3() {
	var xmlStr = `
<people>
	<person age="34">
		<first-name>John</first-name>
		<address>
			<city>San Francisco</city>
			<state>CA</state>
		</address>
	</person>
	<person age="23">
		<address>
			<city>Austin</city>
			<state>TX</state>
		</address>
	</person>
</people>`

	type Address struct {
		City  string `xml:"city"`
		State string `xml:"state"`
	}

	r := bytes.NewBufferString(xmlStr)
	decoder := xml.NewDecoder(r)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF is a successful end
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {

		case xml.StartElement:
			if v.Name.Local == "address" {
				var address Address
				err = decoder.DecodeElement(&address, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement() failed with '%s'\n", err)
					break
				}
				fmt.Printf("%+#v\n", address)
			}
		}
	}
}

func decodeFromReader(r io.Reader) (*People, error) {
	var people People
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&people)
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func decodeFromString(s string) (*People, error) {
	r := bytes.NewBufferString(s)
	return decodeFromReader(r)
}

func decodeFromFile(path string) (*People, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decodeFromReader(f)
}
