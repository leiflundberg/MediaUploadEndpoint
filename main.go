package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received from: " + r.Host)
	r.ParseMultipartForm(1000 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error retrieving the file")
		return
	}
	defer file.Close()
	filename := handler.Filename
	log.Println("Filename: " + filename)
	log.Println("Size in bytes: " + strconv.FormatInt(handler.Size, 10))
	extension := filename[strings.LastIndex(filename, ".")+1:]
	log.Println("Extension: ." + extension)
	tempFile, err := os.CreateTemp("uploads", "upload-*."+extension)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error creating temporary file")
	}
	defer tempFile.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error reading received file")
	}
	tempFile.Write(fileBytes)
	log.Println("File successfully uploaded")
	fmt.Fprintf(w, "Successfully uploaded file: "+filename+"\n")
}

func router(port string) {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(port, nil)
}

func main() {
	log.Println("Starting app...")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	port := ":8080"
	log.Println("Media upload endpoint is listening on " + port)
	router(port)
}
