package _Log

func Write(log string) {
	go func() { writeChan <- log }()
}

func Exit(){
	go func() { exitChan <- 0}()
}

