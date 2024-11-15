package webapi

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

/*
apiFileShare: Allows a user to share the file in network by downloading and then uploading the
the file to the warehouse and this will allow multiple versions of the file to be shared.

Request:    GET /file/share?hash=[hash]&node=[node ID]
Response:   200 with the content

	206 with partial content
	400 if the parameters are invalid
	404 if the file was not found or other error on transfer initiate
	502 if unable to find or connect to the remote peer in time
*/
func (api *WebapiInstance) apiFileShare(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error

	// validate hashes (must be blake3) and other input
	fileHash, valid1 := DecodeBlake3Hash(r.Form.Get("hash"))
	nodeID, valid2 := DecodeBlake3Hash(r.Form.Get("node"))

	fmt.Println("hash & node valid? !!!!")
	if !valid1 || (!valid2) {
		http.Error(w, "", http.StatusBadRequest)
		fmt.Println("hash & node NOT valid? !!!!")
		return
	}

	// Download file since the NodeID and Hash is valid
	filePath := r.Form.Get("path")
	if filePath == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Download

	info := &downloadInfo{backend: api.Backend, api: api, id: uuid.New(), created: time.Now(), hash: fileHash, nodeID: nodeID}

	api.Backend.LogError("Download.DownloadStart", "output %v", downloadInfo{backend: api.Backend, api: api, id: uuid.New(), created: time.Now(), hash: fileHash, nodeID: nodeID})

	// Creates the path to warehouse based on the hash provided
	pathFull, err := api.Backend.UserWarehouse.CreateFilePath(fileHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create the file immediately
	if info.initDiskFile(pathFull) != nil {
		EncodeJSON(api.Backend, w, r, apiResponseDownloadStatus{APIStatus: DownloadResponseFileInvalid})
		return
	}

	// add the download to the list
	api.downloadAdd(info)

	// start the download!
	go info.Start()

	api.Backend.LogError("Download.DownloadStart", "output %v", apiResponseDownloadStatus{APIStatus: DownloadResponseSuccess, ID: info.id, DownloadStatus: DownloadWaitMetadata})

}
