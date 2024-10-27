package KotlinServer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var KotlinServerUrl = "http://127.0.0.1:8080/test" // Changed from ws:// to http://

func RequestFile() {
	fmt.Println("starting request...")

	resp, err := http.Get("http://localhost:8080/file/path")

	if err != nil {
		fmt.Println("http.Get error:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error:", err)
		return
	}

	// Get file name from headers if available
	contentDisposition := resp.Header.Get("Content-Disposition")
	var fileName string
	if contentDisposition != "" {
		// Parse filename from content-disposition
		parts := strings.Split(contentDisposition, ";")
		for _, part := range parts {
			if strings.Contains(part, "filename") {
				fileName = strings.Split(part, "=")[1]
				break
			}
		}
	} else {
		fileName = "unknown" // default to 'unknown' if filename not provided
	}

	fileSize := len(bodyBytes) // size of file

	// now you can use bodyBytes variable that contains your file in memory.
	fmt.Printf("File downloaded successfully. File name: %s, size: %d bytes\n", fileName, fileSize)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the request query parameters
	params := r.URL.Query()

	filehash := params.Get("filehash")

	log.Println("File hash received: ", filehash)

	// process and provide the file content based on the 'filehash'
	// ...
}
