package changoroutine

type adder interface {
	AddCounter(int64)
}

func TryChanGoroutine(intChan chan<- int, add adder) {

	// go func() {
	// 	for z := 0; z <= len(intChan); z++ {
	// 		intChan <- z
	// 	}
	// 	close(intChan)
	// }()

	// go add.AddCounter(64)

	go func() {
		for z := 0; z <= len(intChan); z++ {
			intChan <- z

		}
		close(intChan)
	}()
	add.AddCounter(64)
}
