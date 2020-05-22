package main
import(
	"fmt"
    "os"
    "strings"
    "bufio"
) 
type Person struct{
    fname []byte
    lname []byte
}

func main() {

    var textfile string
    fmt.Scanln(&textfile)
    data, err := os.Open(textfile)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    defer data.Close()
    fileScanner := bufio.NewScanner(data)
    kappa := make([]Person,0,100)
    for fileScanner.Scan() {
        sam := strings.Split(fileScanner.Text(), " ")
        jim := Person{fname: []byte(sam[0]), lname: []byte(sam[1])}
        kappa = append(kappa, jim)

    }
    for _, n:= range kappa{
        fmt.Printf("%s %s\n",n.fname,n.lname)

    }
}