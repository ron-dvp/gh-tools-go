package tools

type RepoObject struct {
	Full_name   string
	Description string
	Private     bool
}

type User struct {
	Login string
}

type NewRepo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Url         string `json:"html_url"`
}
