package webapi

import "net/http"

/*
Adds node ID as follower

Request:    POST /profile/AddFollower?followerID=[Follower NodeID]
Response:   200 with JSON structure apiBlockchainBlockStatus
*/
func (api *WebapiInstance) apiAddFollower(w http.ResponseWriter, r *http.Request) {
	var input apiProfileData
	if err := DecodeJSON(w, r, &input); err != nil {
		return
	}

	var fields []uint16

	for n := range input.Fields {
		fields = append(fields, input.Fields[n].Type)
	}

	newHeight, newVersion, status := api.Backend.UserBlockchain.ProfileDelete(fields)

	EncodeJSON(api.Backend, w, r, apiBlockchainBlockStatus{Status: status, Height: newHeight, Version: newVersion})
}
