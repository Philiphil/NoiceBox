/*
NOICESOURCE
stackoverflow, dev.com, site du 0, /g/, 01 net, dev.to, 

NOICESEARCH
Google, stackoverflow

*/


package main
import (
//"fmt"
"io/ioutil"
"net/http"
)

type NoiceSource struct{
	Source string
	Regex string
	InnerDelimiter string
	OuterDelimiter string
	CapturedLinks []string
	Frequence int
}

type NoiceSearch struct{
	Source string
	Keywords []string
	Frequence int
}

func initDevSources()(ns []NoiceSource){
	n0  := NoiceSource{"https://www.developpez.com","/actu/","","", []string{},10}
	ns = append(ns,n0)

	n1  := NoiceSource{"https://www.stackoverflow.com","/questions/","","", []string{},10}
	ns = append(ns,n0)
	return
}

func initNoiceSources()(ns []NoiceSource){
	ns = append(ns,initDevSources())
	return 
}


func main(){
	ns = initDevSources()
}

func goToUrl(url string)(string){
	response, err := http.Get(url)
    if err != nil {
         return ""
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return ""
        }
	    return string(contents)
	}
}