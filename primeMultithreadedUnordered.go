/* creeated Jan 1, 2025
 * made parallel June 24,2026 */
package main

import("fmt"
"os"
"strconv"

"sync"
)

const (
	UPPER_LIMIT=100000
	THREADS=16
	BUF=10
)

func isDivisible(a,b float64)bool{
	g := a / b
	return g == float64(int64(g))
}

func testOne(wg *sync.WaitGroup, f *os.File, mu *sync.Mutex, ch chan int){
	for i_ := range ch{
		c := i_
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

		if isPrime{
			mu.Lock()
			fmt.Println(c)
			f.Write([]byte(strconv.Itoa(c) + "\n"))
			mu.Unlock()
		}

	}
	wg.Done()
}


func main(){
	f,_ := os.Create("primes_multi.txt")
	
	f.Write([]byte("2" + "\n"))
	fmt.Println("2")
	
	cha := make([](chan int), THREADS)
	for i := range cha {
		cha[i] = make(chan int, BUF)
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(THREADS)
	
	for i := 0; i < THREADS; i += 1{
		go testOne(&wg, f, &mu, cha[i])
	}
	


	i := 0
	c := 3
	for c < UPPER_LIMIT{
		cha[i % THREADS] <- c
		i += 1
	c += 2 // loop only does odd for primes
	}

	//destruct
	for _,c := range cha{
		close(c)
	}
	wg.Wait()
}
