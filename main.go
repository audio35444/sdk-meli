package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"./esDriver"
	"./structs/category"
	"./structs/setting"
)

func main() {
	data, err := ioutil.ReadFile("./settings.json")
	// esDriver.NewIndex("items")
	var setting setting.Setting
	if err == nil {
		err1 := json.Unmarshal(data, &setting)
		if err1 == nil {
			resultCategory := category.GetCategories(&setting, "MLA")
			// for _, objC := range resultCategory {
			// 	fmt.Println("Id:", objC.Id, "Name:", objC.Name)
			// }
			resultCategory[0].GetCategoryDetail(&setting, "MLA")
			for _, item := range resultCategory[0].Results {
				item.GetItemDetail(&setting)
				esDriver.AddDocToIndexWithId(item, "items", item.Id)
				// 	item.Save()
				fmt.Println(item.Id, item.Title)
			}
			data, err := esDriver.GetDocs("items")
			if err == nil {
				for _, item := range data.Hits.Hits {
					fmt.Println(item)
				}
			}

			// fmt.Println(len(resultCategory[0].Results[0].Descriptions))
			// resultCategory[0].Results[0].GetItemDetail(&setting)
			// fmt.Println(len(resultCategory[0].Results[0].Descriptions))
			// result := site.GetSites(&setting)
			// fmt.Println(result)
		}
	}
}

// type Setting struct {
// 	RootEndpoint string `json:"root-endpoint"`
// 	Paths        struct {
// 		Site     string `json:"site"`
// 		Category string `json:"category"`
// 		Item     string `json:"item"`
// 	}
// }
