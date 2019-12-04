package main

import (
	"fmt"
	"net/http/httptest"
	"os"
	"io/ioutil"

	"github.com/jmhobbs/advent-of-code/support/badge"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: badge <user-id>")
		os.Exit(1)
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("/?user_id=%s", os.Args[1]), nil)
	w := httptest.NewRecorder()
	badge.Badge(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)


	fmt.Fprintln(os.Stderr, resp.StatusCode)
	fmt.Fprintln(os.Stderr, resp.Header.Get("Content-Type"))
	fmt.Print(string(body))
}
