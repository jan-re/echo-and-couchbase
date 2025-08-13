package main

type book struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Author    string  `json:"author,omitempty"`
	PageCount int     `json:"pageCount,omitempty"`
	Genre     string  `json:"genre,omitempty"`
	Themes    []theme `json:"themes,omitempty"`
}

type theme struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// TODO Obsolete below. Delete
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
