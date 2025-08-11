package main

type createUserReq struct {
	ID string `json:"id"`
}

type getUserReq struct {
	getScore bool `query:"getScore"`
}

type user struct {
	ID    string
	Score int
}
