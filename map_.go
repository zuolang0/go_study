package main
import "fmt"
func main() {
    var countryCapitalMap map[string]string
    countryCapitalMap=make(map[string]string)

    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"
    for country := range countryCapitalMap{
        fmt.Println("Capital of",country,"is",countryCapitalMap[country])
    }
    capital,ok := countryCapitalMap["United states"]

    if (ok) {
        fmt.Println("Capital of United states is ",capital)
    }else{
        fmt.Println("Capital of United states is not present")
    }

    delete(countryCapitalMap,"France")
    fmt.Println("Entry for France is deleted")
    fmt.Println("删除之后的map")
    for country := range countryCapitalMap{
        fmt.Println("Capital of",country,"is",countryCapitalMap[country])
    }
}