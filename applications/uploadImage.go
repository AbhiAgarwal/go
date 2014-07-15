package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"text/template"
)

var uploadTemplate, _ = template.ParseFiles("./templates/upload.html")
var errorTemplate, _ = template.ParseFiles("./templates/error.html")

func errorHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				w.WriteHeader(500)
				errorTemplate.Execute(w, e)
			}
		}()
		fn(w, r)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func uploadPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		uploadTemplate.Execute(w, nil)
		return
	}

	f, _, err := r.FormFile("image")
	check(err)
	defer f.Close()

	// Create a temporary file
	t, err := ioutil.TempFile(".", "image-")
	check(err)
	defer t.Close()

	// Copy the file into the temporary file
	_, err = io.Copy(t, f)
	check(err)

	http.Redirect(w, r, "/view?id="+t.Name()[6:], 302)
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, "image-"+r.FormValue("id"))
}

func main() {
	http.HandleFunc("/", errorHandler(uploadPage))
	http.HandleFunc("/view", errorHandler(view))
	http.ListenAndServe(":8080", nil)
}
