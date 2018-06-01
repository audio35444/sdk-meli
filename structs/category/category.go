package category

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"../item"
	"../setting"
)

//https://api.mercadolibre.com/sites/{Site_id}/categories
//var reTitle = regexp.MustCompile(`<title>.*</title>`)
//body = reTitle.ReplaceAllString(str,newSubstr)
// https://api.mercadolibre.com/sites/MLA/search?category=MLA5726
var reSiteId = regexp.MustCompile(`{Site_id}`)
var reCategoryId = regexp.MustCompile(`{Category_id}`)

type Category struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Results []item.Item `json:"results"`
}

func GetCategories(setting *setting.Setting, siteId string) (categories []Category) {
	path := reSiteId.ReplaceAllString(setting.Paths.Category, siteId)
	res, err := http.Get(setting.RootEndpoint + path)
	if err == nil {
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &categories)
	}
	return
}
func (p *Category) GetCategoryDetail(setting *setting.Setting, siteId string) {
	path := reCategoryId.ReplaceAllString(reSiteId.ReplaceAllString(setting.Paths.CategoryDetail, siteId), p.Id)
	res, err := http.Get(setting.RootEndpoint + path)
	if err == nil {
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &p)
	}
}
