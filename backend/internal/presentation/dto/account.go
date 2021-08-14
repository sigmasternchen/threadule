package dto

type AddAccountResponse struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

type AddAccountResolveParam struct {
	Pin string `json:"pin"`
}
