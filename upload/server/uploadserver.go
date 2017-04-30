package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", handleUpload)
	http.ListenAndServe(":9000", nil)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received upload request")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.Create("../test/" + handler.Filename)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("failed"))
		return
	}
	fmt.Fprintf(w, "%v", handler.Header)
	defer f.Close()
	io.Copy(f, file)
}
