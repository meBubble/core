package KotlinServer

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var KotlinServerUrl = "http://127.0.0.1:8080/test" // Changed from ws:// to http://

func RequestFile(filehash string) {
	fmt.Println()
	fmt.Println("starting request...")
	response, err := http.Get(KotlinServerUrl) // + "?filehash=" + filehash)
	if err != nil {
		fmt.Println("http.Get error:", err)

	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("close Error:", err)
			}
		}(response.Body)

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll:", err)

		} else {

			//fileContent := hex.EncodeToString(body)
			//fmt.Println(fileContent)
			fmt.Println()
			fmt.Printf("Received: %s", body)
			fmt.Println()
		}
	}
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the request query parameters
	params := r.URL.Query()

	filehash := params.Get("filehash")

	log.Println("File hash received: ", filehash)

	// process and provide the file content based on the 'filehash'
	// ...
}
