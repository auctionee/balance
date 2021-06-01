package data

const PROJECT_ID = "auctionee"

type Credentials struct {
	Login string `json:"login"`
	Token string `json:"token"`
}
type Modify struct {
	Login  string `json:"login"`
	Token  string `json:"token"`
	Amount int    `json:"amount"`
}
type Balance struct {
	Login   string `json:"login"`
	Balance int    `json:"balance"`
}
