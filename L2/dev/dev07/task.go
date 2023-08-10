package main

func Channel(channels ...<-chan interface{}) <-chan interface{} {
	signal := make(chan interface{})
	for _, ch := range channels {
		go func(signal chan interface{}, done <-chan interface{}) {
			select {
			case <-signal:
				return
			case <-done:
				close(signal)
				return
			}
		}(signal, ch)
	}
	return signal
}
