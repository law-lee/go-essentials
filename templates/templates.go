package templatedemo

import (
	"html/template"
	"log"
	"os"
)

// To help write more readable templates We can add - at the beginning or end of action as seen in {{end -}}.
// {{range .Tweets}}{{end}} evaluates inner part for every element of []string slice Tweets 
// and sets current context . within the inner part to elements of Tweets slice.
// {{index .RecentTweets 0}} is equivalent to RecentTweets[0] in Go code.
// {{.TweetCount}} means printing the value of TweetCount in current context.
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
}
