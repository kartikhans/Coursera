package main
import(
	"fmt"
	"sort"
	"os"
	st "strconv"
) 

func main() {
    var age2 string
    sli := make([]int,0,6)
    for {
    	fmt.Scanln(&age2)
    	// if(len(sli)==cap(sli)){
    	// 	kappa = make([]int,len(sli),2*len(sli))
    	// 	kappa = sli
    	// 	sli = kappa
    	// }
    	if(age2!="X"){
    		f, err := st.Atoi(age2)
    		if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    		sli = append(sli, f)
    		sort.Ints(sli)
    		fmt.Println(sli)
    	} else {
    		break
    	}
    }
}