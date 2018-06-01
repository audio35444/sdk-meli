package resultEs

type ResultEs struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	Id      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string `json:"result"`
}
