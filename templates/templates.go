package templatedemo

import (
	"log"
	"math"
	"os"
	"text/template"
)

type DataWithMethod struct {
	Field string
}

func (d DataWithMethod) Methods() string {
	return "method value\n"
}

// To help write more readable templates We can add - at the beginning or end of action as seen in {{end -}}.
// {{range .Tweets}}{{end}} evaluates inner part for every element of []string slice Tweets
// and sets current context . within the inner part to elements of Tweets slice.
// {{index .RecentTweets 0}} is equivalent to RecentTweets[0] in Go code.
// {{.TweetCount}} means printing the value of TweetCount in current context.
// {{range .RecentTweets}} changes variable scope and we don’t have access to data outside.
// If we need to access data from upper scope, we can define variables like {{ $tweetCount := len .RecentTweets }}.
func Run() {
	var tmplStr = `User {{.User}} has {{.TotalTweets}} tweets.
{{- $tweetCount := len .RecentTweets }}
Recent tweets:
{{range $idx, $tweet := .RecentTweets}}Tweet {{$idx}} of {{$tweetCount}}: '{{.}}'
{{end -}}
Most recent tweet: '{{index .RecentTweets 0}}'
`

	t := template.New("tweets")
	t, err := t.Parse(tmplStr)
	if err != nil {
		log.Fatalf("template.Parse() failed with '%s'\n", err)
	}
	data := struct {
		User         string
		TotalTweets  int
		RecentTweets []string
	}{
		User:         "kjk",
		TotalTweets:  124,
		RecentTweets: []string{"hello", "there"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	tmplStr1 := "Data: {{.}}\n"
	t1 := template.Must(template.New("simple").Parse(tmplStr1))
	execWithData := func(data interface{}) {
		err := t1.Execute(os.Stdout, data)
		if err != nil {
			log.Fatalf("t.Execute() failed with '%s'\n", err)
		}
	}

	execWithData(5)
	execWithData("foo")
	st := struct {
		Number int
		Str    string
	}{
		Number: 3,
		Str:    "hello",
	}
	execWithData(st)

	var tmpl2 = `Data from a field: '{{ .Field }}'
	Data from a method: '{{ .Methods }}'`
	t2 := template.Must(template.New("withmethods").Parse(tmpl2))

	d1 := DataWithMethod{
		Field: "field value",
	}
	err = t2.Execute(os.Stdout, d1)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func IfAction() {
	const tmplStr = `{{range . -}}
{{if .IsNew}}'{{.Name}}' is new{{else}}'{{.Name}}' is not new{{end}}
{{end}}`

	t := template.Must(template.New("if").Parse(tmplStr))

	data := []struct {
		Name  string
		IsNew bool
	}{
		{"Bridge", false},
		{"Electric battery", true},
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func FalseValue() {
	const tmplStr = `{{range . -}}
{{printf "%- 16s" .Name}} is: {{if .Value}}true{{else}}false{{end}}
{{end}}`

	t := template.Must(template.New("if").Parse(tmplStr))

	var nilPtr *string = nil
	var nilSlice []float32
	emptySlice := []int{}

	data := []struct {
		Name  string
		Value interface{}
	}{
		{"bool false", false},
		{"bool true", true},
		{"integer 0", 0},
		{"integer 1", 1},
		{"float32 0", float32(0)},
		{"float64 NaN", math.NaN},
		{"empty string", ""},
		{"non-empty string", "haha"},
		{"nil slice", nilSlice},
		{"empty slice", emptySlice},
		{"non-empty slice", []int{3}},
		{"nil pointer", nilPtr},
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func AvoidEmptyValue() {
	type UserTweets struct {
		User   string
		Tweets []string
	}

	const tmplStr = `
	{{- if not .Tweets -}}
	User '{{.User}}' has no tweets.
	{{ else -}}
	User '{{.User}}' has {{ len .Tweets }} tweets:
	{{ range .Tweets -}}
	  '{{ . }}'
	{{ end }}
	{{- end}}`

	t := template.Must(template.New("if").Parse(tmplStr))

	data := UserTweets{
		User: "kjk",
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	data = UserTweets{
		User:   "masa",
		Tweets: []string{"tweet one", "tweet two"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func RangeSlice() {
	const tmplStr = `Elements of arrays or slice: {{ range . }}{{ . }} {{end}}
`
	t := template.Must(template.New("range").Parse(tmplStr))

	array := [...]int{3, 8}
	err := t.Execute(os.Stdout, array)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	slice := []int{12, 5}
	err = t.Execute(os.Stdout, slice)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func RangeMap() {
	const tmplStr = `Elements of map:
{{ range $k, $v := . }}{{ $k }}: {{ $v }}
{{end}}`

	t := template.Must(template.New("range").Parse(tmplStr))

	data := map[string]int{
		"one":  1,
		"five": 5,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func RangeChannel() {
	const tmplStr = `Elements of a channel: {{ range . }}{{ . }} {{end}}
`

	t := template.Must(template.New("range").Parse(tmplStr))

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	err := t.Execute(os.Stdout, ch)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func FuncLogic() {
	const tmplStr = `Or:  {{ if or .True .False }}true{{ else }}false{{ end }}
And: {{ if and .True .False }}true{{ else }}false{{ end }}
Not: {{ if not .False }}true{{ else }}false{{ end }}
`

	t := template.Must(template.New("and_or_not").Parse(tmplStr))

	data := struct {
		True  bool
		False bool
	}{True: true, False: false}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func FuncIndex() {
	const tmplStr = `Slice[0]: {{ index .Slice 0 }}
SliceNested[1][0]: {{ index .SliceNested 1 0 }}
Map["key"]: {{ index .Map "key" }}
`

	t := template.Must(template.New("index").Parse(tmplStr))

	data := struct {
		Slice       []string
		SliceNested [][]int
		Map         map[string]int
	}{
		Slice: []string{"first", "second"},
		SliceNested: [][]int{
			{3, 1},
			{2, 3},
		},
		Map: map[string]int{
			"key": 5,
		},
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func FuncLen() {
	const tmplStr = `len(nil)       : {{ len .SliceNil }}
len(emptySlice): {{ len .SliceEmpty }}
len(slice)     : {{ len .Slice }}
len(map)       : {{ len .Map }}
`

	t := template.Must(template.New("len").Parse(tmplStr))

	data := struct {
		SliceNil   []int
		SliceEmpty []string
		Slice      []bool
		Map        map[int]bool
	}{
		SliceNil:   nil,
		SliceEmpty: []string{},
		Slice:      []bool{true, true, false},
		Map:        map[int]bool{5: true, 3: false},
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func FuncPrint() {
	const tmplStr = `print:   {{ print .Str .Num }}
println: {{ println .Str .Num }}
printf:  {{ printf "%s %#v %d" .Str .Str .Num }}
`

	t := template.Must(template.New("print").Parse(tmplStr))

	data := struct {
		Str string
		Num int
	}{
		Str: "str",
		Num: 8,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}

func FuncJS() {
	const tmplStr = `js escape  : {{ js .JS }}
html escape: {{ html .HTML }}
url escape : {{ urlquery .URL }}
`

	t := template.Must(template.New("print").Parse(tmplStr))

	data := struct {
		JS   string
		HTML string
		URL  string
	}{
		JS:   `function me(s) { return "foo"; }`,
		HTML: `<div>text</div>`,
		URL:  `http://www.programming-books.io`,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}
