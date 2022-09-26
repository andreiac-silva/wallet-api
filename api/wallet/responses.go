package api

var EmptyResponse = map[string]string{}

type CreateResponse struct {
	ID string `json:"wallet_id"`
}
