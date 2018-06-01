package site

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../setting"
)

type Site struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//agregando el parametro antes del nombre del method
func GetSites(settingNow *setting.Setting) (sites []Site) {
	// fmt.Println(*settingNow.Paths.Site)
	res, err := http.Get(settingNow.RootEndpoint + settingNow.Paths.Site)
	if err == nil {
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &sites)
	}
	return
}
