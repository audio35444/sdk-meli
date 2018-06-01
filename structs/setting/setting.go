package setting

type Setting struct {
	RootEndpoint string `json:"root-endpoint"`
	Paths        struct {
		Site           string `json:"site"`
		Category       string `json:"category"`
		CategoryDetail string `json:"categoryDetail"`
		Item           string `json:"item"`
	}
}

//devuelve la URL de sitios
func (p Setting) GetEndpointSite() string {
	return p.RootEndpoint + p.Paths.Site
}
