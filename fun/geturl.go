package main

import "fmt"

//import "io/ioutil"
import "net/http"
import "os"
import "strings"
import "io"

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :%v \n", err)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :reading %s:%v\n", url, err)
			os.Exit(1)
		}
		/*var ptr *int
		  ptr = &n*/
		//fmt.Printf("%s", b)
		fmt.Printf("%x", ptr)
	}

}
