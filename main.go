package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

func makeGist(fileBody string) string {
	isPublic := false
	desc := ""
	gf := github.GistFile{Content: &fileBody}
	files := map[github.GistFilename]github.GistFile{"file_one": gf}

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

func getStringsFromStdin() []string {
	reader := bufio.NewReader(os.Stdin)

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

func main() {
	body := getStringsFromStdin()
	log.Printf(makeGist(strings.Join(body, "\n")))
}
