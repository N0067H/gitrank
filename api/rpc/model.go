package rpc

type ResponseUser struct {
	Login              string `json:"login"`
	TotalContributions uint32 `json:"totalContributions"`
}
