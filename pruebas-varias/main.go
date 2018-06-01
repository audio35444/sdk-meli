package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var cant = 10

func main() {
	waitMethod()
	notWaitMethod()
}
func testAsync1(arr *[]string) {
	http.Get("https://api.mercadolibre.com/sites/")
	mutex.Lock()
	*arr = append(*arr, "test1")
	mutex.Unlock()
	wg.Done()
}
func testAsync2(arr *[]string) {
	http.Get("https://api.mercadolibre.com/sites/")
	mutex.Lock()
	*arr = append(*arr, "test2")
	mutex.Unlock()
	wg.Done()
}
func waitMethod() {
	t := time.Now()
	var arr []string
	wg.Add(cant)
	for i := 0; i < cant; i++ {
		if i%2 == 0 {
			go testAsync1(&arr)
		} else {
			go testAsync2(&arr)
		}
	}
	wg.Wait()
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Println("waitMethod", time.Since(t))
}
func notWaitMethod() {
	t := time.Now()
	var arr []string
	// wg.Add(cant)
	for i := 0; i < cant; i++ {
		if i%2 == 0 {
			testAsync3(&arr)
		} else {
			testAsync4(&arr)
		}
	}
	// wg.Wait()
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Println("notWaitMethod", time.Since(t))
}
func testAsync3(arr *[]string) {
	http.Get("https://api.mercadolibre.com/sites/")
	*arr = append(*arr, "test3")
}
func testAsync4(arr *[]string) {
	http.Get("https://api.mercadolibre.com/sites/")
	*arr = append(*arr, "test4")
}

// func prueba3() {
// 	var arr []string
// 	arr = append(arr, "")
// 	arr = append(arr, "")
//
// 	wg.Add(10)
//
// 	go func(arrIndex int) {
// 		// s[10] = "hola"
// 		arr[arrIndex] = "valor1"
// 		wg.Done()
// 	}(0)
// 	go func(arrIndex int) {
// 		// s[10] = "hola"
// 		arr[arrIndex] = "valor2"
// 		wg.Done()
// 	}(1)
// 	wg.Wait()
// 	fmt.Println(arr)
// }
// func prueba2() {
// 	var func1 string
// 	var func2 string
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	// objTest[3]=""
// 	go func() {
// 		// s[10] = "hola"
// 		func1 = "valor1"
// 		wg.Done()
// 	}()
// 	go func() {
// 		// s[10] = "hola"
// 		func2 = "valor2"
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println(func1, func2)
// }
// func prueba1() {
// 	var v int
// 	var objTest []string
// 	objTest = append(objTest, "d")
//
// 	objTest = append(objTest, "a")
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	// objTest[3]=""
// 	go func(s *[]string) {
// 		// s[10] = "hola"
// 		v = 1
// 		*s = append(*s, "en func1")
// 		wg.Done()
// 	}(&objTest)
// 	// objTest[4]=""
// 	go func(s *[]string) {
// 		*s = append(*s, "en func2")
// 		fmt.Println(v)
// 		wg.Done()
// 	}(&objTest)
// 	wg.Wait()
// 	fmt.Println(v)
// 	fmt.Println(objTest)
// }
