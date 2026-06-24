package main

import("fmt"
"os"
"strconv"
)


func isDivisible(a,b float64)bool{
	g := a / b
	return g == float64(int64(g))
}



func main(){
	f,_ := os.Create("primes.txt")
	
	f.Write([]byte("2" + "\n"))
	
	c := 3
	for c < 100000 {
		isPrime := true
		
			for i := c; i > 3 && isPrime; i -= 2{
				if(i==c) {
					continue;
					} else {
			
					isPrime = !isDivisible(float64(c),float64(i))
					if !isPrime{
						break;
					}
				}
			}
					
			if isPrime && !isDivisible(float64(c),2.0){ // loop only does odd primes
				fmt.Println(c)
				f.Write([]byte(strconv.Itoa(c) + "\n"))
			}
	c += 2
	}
}
