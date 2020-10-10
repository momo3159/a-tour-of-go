package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	
	cnt := 0
	an, an1 := 0, 1 
	
	return func() int {
		if cnt == 0 {
			cnt++
			return 0
		} else {
		    retVal := an + an1
			an = an1
			an1 = retVal
			
			return retVal
		}
		
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
