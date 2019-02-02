package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example. V7"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	// b, _ := ioutil.ReadAll(r.Body)
	// spew.Dump(string(b))

	r.ParseForm()
	spew.Dump(r.Form)

	// fmt.Fprintf(w, "Hello from:  "+title+"\n")
	if len(r.Form["Body"]) > 0 {
		fmt.Fprintf(w, "You said:  "+r.Form["Body"][0]+"\n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
