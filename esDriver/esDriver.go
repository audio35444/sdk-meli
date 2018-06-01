package esDriver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"../structs/item"
	"../structs/resultEs"
	"./search"
	"./setting"
)

//el pretty que s emanda por get le indica a elasticsearch que devuelva en formato json
//curl -X GET "http://localhost:9200/_cat/indices?v"  consultar indices
//curl -X PUT "localhost:9200/client?pretty"
//curl -X GET "localhost:9200/client/_doc/1?pretty" consultar el doc 1 del index client
//curl -X DELETE "localhost:9200/client?pretty" eliminar el index client
//curl -X POST "localhost:9200/client/_doc/1/_update?pretty" -H 'Content-Type: application/json' -d'  update sobre el doc 1 index client
// {
//   "script" : "ctx._source.age += 5"
// }
// '

//curl -X POST "localhost:9200/client/_doc?pretty" -H 'Content-Type: application/json' -d agregar un doc nuevo id=autogenerado al index client
//{
//   "name": "John Doe"
//}
//curl -X GET "localhost:9200/{index_name}/_search?pretty"
// curl -X GET "localhost:9200/{index_name}/_search?Q=*&sort={attribute_name}:(asc|dsc)&pretty"

var obj setting.Setting

//curl -X PUT "localhost:9200/client/_doc/1?pretty" -H 'Content-Type: application/json' -d' agregar un doc nuevo id=1 al index client
// {
//   "name": "John Doe"
// }
// '
type SP struct {
	Script string `json:"script"`
}

// func main() {
// 	obj.LoadSetting()
// 	// newIndex()
// 	// getIndices()
// 	// deleteIndex()
// 	// getIndices()
// 	getDocs()
// 	// addDocToIndex()
// 	// addDocToIndexWithId()
// 	// pruebaUpdate()
// 	// getDocFromIndex()
// 	// deleteDoc()
// 	// getDocFromIndex()
// }
func GetDocs(indexName string) (*search.Search, *error) {
	obj.LoadSetting()
	var objSearch search.Search
	data, err := genericRequest("GET",
		obj.GetEndpointDocSearch(indexName),
		nil,
		false)
	if *err == nil {
		json.Unmarshal(*data, &objSearch)
		// fmt.Println(string(*data))
		return &objSearch, nil
		// for _, item := range objSearch.Hits.Hits {
		// 	fmt.Println(item)
		// }
	} else {
		return nil, err
		// fmt.Println(error(*err).Error())
	}
}
func DeleteDoc(indexName string, docId string) (*resultEs.ResultEs, *error) {
	obj.LoadSetting()
	var objResult resultEs.ResultEs
	data, err := genericRequest("DELETE",
		obj.GetEndpointDocDelete(indexName, docId),
		nil,
		false)
	if *err == nil {
		// fmt.Println(data)
		json.Unmarshal(*data, &objResult)
		return &objResult, nil
	} else {
		return nil, err
	}

	// if *err == nil {
	// 	fmt.Println(string(*data))
	// } else {
	// 	fmt.Println(error(*err).Error())
	// }
}
func DeleteIndex() {
	if len(os.Args) >= 2 {
		data, err := genericRequest("DELETE",
			obj.GetEndpointNewIndex(os.Args[1]),
			nil,
			false)
		if *err == nil {
			fmt.Println(string(*data))
		} else {
			fmt.Println(error(*err).Error())
		}
	}
}

type Prueba struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PruebaUpdate() {
	var script SP
	// client := &http.Client{}
	script.Script = "ctx._source.age += 5"
	result, _ := json.Marshal(script)
	data, err := genericRequest("POST",
		obj.GetEndpointDocUpdate("nuevoindice", "BBxKp2MBedOyFXZO2eFz"),
		strings.NewReader(string(result)),
		true)
	if *err == nil {
		fmt.Println(string(*data))
	} else {
		fmt.Println(error(*err).Error())
	}
}
func AddDocToIndex(element interface{}, indexName string) (*resultEs.ResultEs, *error) {
	obj.LoadSetting()
	// var prueba Prueba
	// prueba.Name = "Juan Emmanuel"
	// prueba.Age = 25
	var objResult resultEs.ResultEs
	result, _ := json.Marshal(element)
	data, err := genericRequest("POST",
		obj.GetEndpointDocIndex(indexName),
		strings.NewReader(string(result)),
		true)
	if err == nil {
		json.Unmarshal(*data, &objResult)
		return &objResult, nil
	} else {
		return nil, err
	}

	// if *err == nil {
	// 	return *data
	// } else {
	// 	fmt.Println(error(*err).Error())
	// }
}
func AddDocToIndexWithId(element interface{}, indexName string, docId string) (dataResult *[]byte, errResult *error) {
	obj.LoadSetting()
	result, _ := json.Marshal(element)
	dataResult, errResult = genericRequest("PUT",
		obj.GetEndpointDocIndexWithId(indexName, docId),
		strings.NewReader(string(result)),
		true)
	return
	// if *err == nil {
	// 	fmt.Println(string(*data))
	// } else {
	// 	fmt.Println(error(*err).Error())
	// }
}
func GetDocFromIndex(indexName string, docId string) (*item.Item, *error) {
	obj.LoadSetting()
	data, err := genericRequest("GET",
		obj.GetEndpointDocFromIndex(indexName, docId),
		nil,
		false)
	if *err == nil {
		type Source struct {
			Source item.Item `json:"_source"`
		}
		var objItem Source
		// fmt.Println(string(*data))
		json.Unmarshal(*data, &objItem)
		// fmt.Println(string(*data))
		return &objItem.Source, nil
		// fmt.Println(string(*data))
	} else {
		return nil, err
		// fmt.Println(error(*err).Error())
	}
}
func NewIndex(indexName string) {
	data, err := genericRequest("PUT",
		obj.GetEndpointNewIndex(indexName),
		nil,
		false)
	if *err == nil {
		fmt.Println(string(*data))
	} else {
		fmt.Println(error(*err).Error())
	}
}
func GetIndices() {
	data, err := genericRequest("GET",
		obj.GetEndpointIndices(),
		nil,
		false)
	if *err == nil {
		fmt.Println(string(*data))
	} else {
		fmt.Println(error(*err).Error())
	}
}
func genericRequest(method string, fullPath string, body io.Reader, isJson bool) (dataResult *[]byte, errResult *error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, fullPath, body)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	if isJson {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		data, err1 := ioutil.ReadAll(res.Body)
		dataResult = &data
		errResult = &err1
		return
	} else {
		errResult = &err
		return
	}
}
func genericShow(res *http.Response, err error) {
	if err == nil {
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println("\n---------------- Elasticsearch Indices ----------------")
		fmt.Println(string(data))
		fmt.Println("--------------------------------")
	} else {
		fmt.Println(err.Error())
	}
}
