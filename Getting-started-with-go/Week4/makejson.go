package main
import(
	"fmt"
    "encoding/json"
    "os"
) 

func main() {
    var name string
    var address string
    fmt.Print("Name? ")

    fmt.Scanln(&name)
    fmt.Print("Address? ")
    fmt.Scanln(&address)
    idMap := make(map[string]string)
    idMap["name"] = name
    idMap["address"] = address
    barr, err := json.Marshal(idMap)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    fmt.Println(string(barr))

    // var kaim map[string]string

    // srr := json.Unmarshal(barr, &kaim)
    // if srr != nil {
    //     // handle error
    //     fmt.Println(srr)
    //     os.Exit(2)
    // }
    // fmt.Println(kaim)
}