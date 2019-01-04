/*NOICESOURCE
	Source string
	Regex string
	InnerDelimiter string
	OuterDelimiter string
	CapturedLinks
	Frequentie int
Noice source : stackoverflow, dev.com, site du 0, /g/, 01 net, dev.to, 

NOICESEARCH
	link string
	keywords string[]
	frequentie
Google, stackoverflow

*/


package main
import (
"fmt"
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


func main(){
	fmt.Println(goToUrl("http://theo.sorr.io"))
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