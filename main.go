package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/github"
)

func makeGist(fileBody, fileName string) string {
	if fileName == "" {
		fileName = "File one"
	}
	gistFilename := github.GistFilename(filepath.Base(fileName))
	isPublic := false
	desc := ""
	gf := github.GistFile{Content: &fileBody}
	files := map[github.GistFilename]github.GistFile{gistFilename: gf}

	newGist := &github.Gist{
		Files:       files,
		Public:      &isPublic,
		Description: &desc}

	clt := github.NewClient(nil).Gists
	gist, _, err := clt.Create(newGist)

	if err != nil {
		log.Fatal("Upload error:", err)
	}

	return *gist.HTMLURL
}

func getStringsFromReader(reader *bufio.Reader) []string {
	result := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			result = append(result, line)
		}

		if err != nil {
			break
		}
	}

	return result
}

func stdInToString() []string {
	reader := bufio.NewReader(os.Stdin)
	return getStringsFromReader(reader)
}

func fileToString(fileName string) []string {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	return getStringsFromReader(reader)
}

func main() {
	var body []string
	fileName := ""

	if len(os.Args) > 1 {
		fileName = os.Args[1]
		body = fileToString(fileName)
	} else {
		body = stdInToString()
	}

	url := makeGist(strings.Join(body, ""), fileName)
	println(url)
}
