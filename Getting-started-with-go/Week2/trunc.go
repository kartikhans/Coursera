package main
import(
	"fmt"
) 

func main() {
    var age float64
    var age2 int
    fmt.Scanln(&age)
    age2 = int(age)
    fmt.Print(age2)
}