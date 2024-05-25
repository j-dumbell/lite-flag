package api

// // ToDo
// func (api *API) PostKey(r *http.Request) chix.Response {
// 	var body auth.CreateApiKeyParams
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		return chix.BadRequest("invalid JSON body")
// 	}
//
// 	// ToDo - handle permissions
// 	//	- root can create admin / readonly
// 	//	- admin can create readonly
//
// 	apiKey, err := api.authService.CreateKey()
// 	if err != nil {
// 	}
// }
