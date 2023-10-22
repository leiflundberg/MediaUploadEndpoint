package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(1000 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	filename := handler.Filename
	fmt.Printf("Uploaded File: %+v\n", filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	extension := filename[strings.LastIndex(filename, ".")+1:]
	fmt.Println("Extension: ." + extension)
	tempFile, err := os.CreateTemp("uploads", "upload-*."+extension)
	if err != nil {
		fmt.Println("Permission denied here? inside CreateTemp")
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Permission denied here? inside ReadAll")
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Media upload endpoint is listening...")
	setupRoutes()
}
