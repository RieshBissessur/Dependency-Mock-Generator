package model

type Mappings struct {
	Mappings []Mapping `json:"mappings"`
	Meta     *Meta     `json:"meta"`
	Name     *string   `json:"name"`
}

type Mapping struct {
	ID       string   `json:"id"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	UUID     string   `json:"uuid"`
}

type Request struct {
	UrlPattern string `json:"urlPattern"`
	Method     string `json:"method"`
}

type Response struct {
	Status   int               `json:"status"`
	JsonBody map[string]string `json:"jsonBody"`
}

type Meta struct {
	Total int `json:"total"`
}
