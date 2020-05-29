package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var dropPath string = "/tmp"

func main() {
	if value, exists := os.LookupEnv("DROP_PATH"); exists {
		dropPath = value
	}

	http.HandleFunc("/upload", upload)

	http.ListenAndServe(":8080", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "POST" {
		r.ParseMultipartForm(128 << 20)

		file, handler, err := r.FormFile("uploadFile")

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer file.Close()

		f, err := os.OpenFile(fmt.Sprintf("%v/%v", dropPath, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer f.Close()

		io.Copy(f, file)

		w.Write([]byte(fmt.Sprintf("Uploaded file: %v\n", handler.Filename)))
	}
}
