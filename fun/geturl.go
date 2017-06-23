package main

import "fmt"

//import "io/ioutil"
import "net/http"
import "os"

import "io"

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :%v \n", err)
		}
		//dst := ""
		//b, err := ioutil.ReadAll(resp.Body)
		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :reading %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", n)
	}

}
