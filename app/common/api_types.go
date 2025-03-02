package common

type ApiInterfaceDesc struct {
	Document string `json:"document"`
	Input    string `json:"input"`
	Output   string `json:"output"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}
type MetaData struct {
	ApiVersion       string `json:"apiVersion"`
	ApiBrief         string `json:"api_brief"`
	ServiceNameCN    string `json:"serviceNameCN"`
	ServiceShortName string `json:"serviceShortName"`
}
type Member struct {
	Disabled         bool   `json:"disabled"`
	Document         string `json:"document"`
	Example          string `json:"example"`
	Member           string `json:"member"`
	Name             string `json:"name"`
	Required         bool   `json:"required"`
	Type             string `json:"type"`
	ValueAllowedNull bool   `json:"value_allowed_nul"`
}
type ApiObject struct {
	Document string   `json:"document"`
	Members  []Member `json:"members"`
	Type     string   `json:"type"`
}

type Api struct {
	Actions  map[string]ApiInterfaceDesc `json:"actions"`
	MetaData MetaData                    `json:"metadata"`
	Objects  map[string]ApiObject        `json:"objects"`
}
