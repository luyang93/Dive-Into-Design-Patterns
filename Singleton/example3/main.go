package main

func main() {
	for i := 0; i < 100; i++ {
		go GetInstance()
		wg.Add(1)
		go GetInstance2()
		wg.Add(1)
	}

	wg.Wait()
}
