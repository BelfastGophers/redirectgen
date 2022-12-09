package main

import (
	_ "embed"
	"html/template"
	"os"
	"path"
	"strings"
)

//go:embed template.html
var redirectTemplate string

func main() {

	event := os.Args[1]
	desc := os.Args[2]
	meetup := os.Args[3]
	card := os.Args[4]
	shortEvent := event[:3] + event[len(event)-4:]
	dir := strings.ToLower(shortEvent)
	index := path.Join(dir, "index.html")

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(index)
	if err != nil {
		os.RemoveAll(dir)
		panic(err)
	}

	t, err := template.New("redirectTemplate").Parse(redirectTemplate)
	if err != nil {
		os.RemoveAll(dir)
		panic(err)
	}

	err = t.Execute(f, struct {
		Event       string
		Description string
		Meetup      string
		Card        string
		ShortEvent  string
	}{
		event,
		desc,
		meetup,
		card,
		shortEvent,
	})
	if err != nil {
		os.RemoveAll(dir)
		panic(err)
	}
}
