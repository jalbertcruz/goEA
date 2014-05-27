package main

//import "fmt"
import "goEA/pea"

func testIsland1(){
	ch1 := make(chan pea.TIndividual, 1)
	ch2 := make(chan pea.TIndividual, 1)
	control := make(chan string, 1)
	cnf := pea.ConfIsland{control, []pea.TIndividual{"10101000", "10101101", "10101101", "11101101", "10101100", "00101111"}, ch1, ch2, 2, 10, 2, 500}
	go pea.PoolManager(cnf)
	control <- "start"
}

func main() {
	testIsland1()
}
