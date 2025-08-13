package main

type createUserReq struct {
	ID string `json:"id"`
}

type getUserReq struct {
	ID       string `param:"id"`
	GetScore string `query:"getScore"`
}

type putUserReq struct {
	ID    string  `param:"id"`
	Score *int    `json:"score,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type user struct {
	ID    string
	Score int
	Name  string
}
