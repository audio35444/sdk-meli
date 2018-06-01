package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"../../persistence/files"
	"../setting"
)

//items/{itemId}
//https://api.mercadolibre.com/sites/MLA/search?category=MLA5726
//https://api.mercadolibre.com/items/MLA663090075
var SAVE_MODE = "files"
var reItemId = regexp.MustCompile(`{Item_id}`)

type Item struct {
	Id                string  `json:"id"`
	Title             string  `json:"title"`
	CategoryId        string  `json:"category_id"`
	price             float64 `json:"price"`
	CurrencyId        string  `json:"currency_id"`
	AvailableQuantity int     `json:"available_quantity"`
	BuyingMode        string  `json:"buying_mode"`
	ListingTypeId     string  `json:"listing_type_id"`
	Condition         string  `json:"condition"`
	Descriptions      []struct {
		Id string `json:"id"`
	} `json:"descriptions"`
	VideoId  string `json:"video_id"`
	Pictures []struct {
		Source string `json:"source"`
	} `json:"pictures"`
}

func (p *Item) GetItemDetail(setting *setting.Setting) {
	path := reItemId.ReplaceAllString(setting.Paths.Item, p.Id)
	res, err := http.Get(setting.RootEndpoint + path)
	if err == nil {
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &p)
	}
}
func (p *Item) Save() {
	switch SAVE_MODE {
	case "files":
		files.InsertNewElement("database.file", p.Title)
	case "elasticsearch":
		fmt.Println("Not implement")

	}
}
