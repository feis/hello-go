package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	resp, _ := http.Get("https://www.ptt.cc/bbs/movie")
	body, _ := io.ReadAll(resp.Body)
	result := strings.ReplaceAll(string(body), "//images.ptt.cc/", "https://images.ptt.cc/")
	os.WriteFile("movie.html", []byte(result), 0644)
}
