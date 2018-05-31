package main

import (
	"net/http"
	"html/template"
	"bufio"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/file_form_template.gohtml"))
}

func main() {
	http.HandleFunc("/", fileFormHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func fileFormHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=`UTF-8`")

	// We initialize an empty map that will be the container for
	// our word occurrences.
	wordsMap := map[string]int{}

	if req.Method == http.MethodPost {
		// Request struct provides handy FormFile() method. It will
		// return file attached to the form, as well, as set of
		// headers with file stats.
		f, _, err := req.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// File implements io.Reader interface, so it is possible to make
		// new scanner object easily. Scanner allows for advanced data
		// manipulation so it will fit perfect our needs.
		// We want to read every token (word) from submitted file.
		scanner := bufio.NewScanner(f)

		// Scanner by default uses ScanLines() SplitFunc, so this is not
		// the behavior we expect in that particular case. To accomplish
		// what we need we can setup ScanWords() SplitFunc instead.
		// It will treat every word as a token, rather than every line.
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			wordsMap[scanner.Text()]++
		}
	}

	tpl.ExecuteTemplate(w,"file_form_template.gohtml", wordsMap)
}