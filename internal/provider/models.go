package provider

type DevOps_Api struct {
	Id   string    `json:"id"`
	Devs []Dev_Api `json:"dev"`
	Ops  []Ops_Api `json:"ops"`
}

type Dev_Api struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Engineers []Engineer_Api `json:"engineers"`
}

type Ops_Api struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Engineers []Engineer_Api `json:"engineers"`
}

type Engineer_Api struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}

/*
type Engineer_Api_NoId struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
*/
