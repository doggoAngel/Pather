package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		logo()
		fmt.Printf("Usage: %s <url> <wordlist>", os.Args[0])
		return
	}
	logo()
	var url string = os.Args[1]
	var wordlist string = os.Args[2]

	content, err := os.ReadFile(wordlist)

	if err != nil {
		fmt.Println("Error! FILE NOT FOUND")
		return
	}
	c := strings.Split(string(content), "\n")
	//check the endpoint is live
	checker(url)

	fmt.Printf("> 0%%")

	for i := 0; i < len(c); i++ {
		var result string = send(url,c[i])
		percent := int(float64(i+1)/float64(len(c))*100) 
		if(result != ""){
			fmt.Print("\r\033[K")
			fmt.Println(result)
		}
		fmt.Printf("\r> %d%%", percent)
	}

}

func send(url string, path string) string {
	resp, _ := http.Get(url + path)
	if resp.StatusCode == 200 || resp.StatusCode == 403 {
		return url  + path + " Code: " + strconv.Itoa(resp.StatusCode)
	}
	return ""
}

func checker(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error! Endpoint is not available")
		os.Exit(0)
	}

}

func logo() {
	fmt.Println(`
 ____       _   _       ____                      _               
|  _ \ __ _| |_| |__   / ___|  ___  __ _ _ __ ___| |__   ___ _ __ 
| |_) / _' | __| '_ \  \___ \ / _ \/ _' | '__/ __| '_ \ / _ \ '__|
|  __/ (_| | |_| | | |  ___) |  __/ (_| | | | (__| | | |  __/ |   
|_|   \__,_|\__|_| |_| |____/ \___|\__,_|_|  \___|_| |_|\___|_|   
	
                         by B1M0
	`)
}
