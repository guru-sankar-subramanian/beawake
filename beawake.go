package main

/*
#cgo LDFLAGS: -framework IOKit -framework CoreFoundation
#include <IOKit/pwr_mgt/IOPMLib.h>

extern IOPMAssertionID createSleepAssertion();
extern void releaseSleepAssertion(IOPMAssertionID assertionID);
*/
import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("🔌 Requesting macOS Power Assertion...")
	assertionID := C.createSleepAssertion()
	if assertionID == 0 {
		fmt.Println("Failed to create power assertion.")
		return
	}

	fmt.Println("System is now locked awake via native IOKit!")
	fmt.Println("Press Ctrl+C to release the lock and allow sleep.")

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	<-stopSignal

	fmt.Println("\nReleasing power assertion...")
	C.releaseSleepAssertion(assertionID)

	fmt.Println("Mac restored to normal sleep behavior. Goodbye!")
}
