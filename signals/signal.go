package signals

import (
	"fmt"
	"os"
	"os/signal"
)

func Run() {
	sigChan := make(chan os.Signal)

	// assign all signal notifications to the channel
	signal.Notify(sigChan)

	// blocks until you get a signal from the OS
	select {
	case sig := <-sigChan:
		fmt.Println("Received signal from OS:", sig)
	}
}
