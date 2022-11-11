package provider

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

type Engineer_map struct {
	Engineers map[string]Engineer_Api `json:"engineers"`
}
type Engineer_Api struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}
