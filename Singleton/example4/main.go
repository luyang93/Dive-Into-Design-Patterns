package main

func main() {
	for i := 0; i < 30; i++ {
		go GetInstance()
		wg.Add(1)
	}

	wg.Wait()
}
