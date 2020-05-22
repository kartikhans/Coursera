package main
import(
	"fmt"
	s "strings"
) 

func main() {
    var age2 string
    fmt.Scanln(&age2)
    age2 = s.TrimSpace(age2)
    age2 = s.ToLower(age2)
    if(s.HasPrefix(age2,"i") && s.Contains(age2,"a") && s.HasSuffix(age2,"n")){
    	fmt.Printf("Found")
    } else{
    	fmt.Printf("Not Found")
    }
}