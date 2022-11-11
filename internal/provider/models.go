package provider

type Dev struct {
	Engineers []Engineer_Api `json:"engineer_map"`
	Id        string         `json:"id"`
	Name      string         `json:"name"`
}

type Ops_Api struct {
	Engineers []Engineer_Api `json:"engineer_map"`
	Id        string         `json:"id"`
	Name      string         `json:"name"`
}

type Engineer_map struct {
	Engineers map[string]Engineer_Api `json:"engineer_map"`
}
type Engineer_Api struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}
