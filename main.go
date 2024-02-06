package main

func addsOne(mychan chan interface{}) {
	mychan <- 1
}

func main() {

	mychan := make(chan interface{})

	go addsOne(mychan)

	for i := 0; i < 1; i++ {
		<-mychan
	}
	println("fim")

}
