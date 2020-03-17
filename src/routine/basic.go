package main
import(
	"fmt"
	"sync"
)
func sayHello(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Hi")
}

func main(){
  var wg sync.WaitGroup
  wg.Add(1)
  go sayHello(&wg);
  wg.Wait()

}