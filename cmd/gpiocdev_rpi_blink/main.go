package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/warthog618/go-gpiocdev"
	"github.com/warthog618/go-gpiocdev/device/rpi"
)

func main() {
	fmt.Println(rpi.J8p7)
	fmt.Println(rpi.Pin("J8p7"))
	fmt.Println(rpi.Pin("GPIO08"))
	fmt.Println(rpi.Pin("8"))

	// toggle a GPIO as output

	offset := 22
	chip := "gpiochip0"
	v := 0
	l, err := gpiocdev.RequestLine(chip, offset, gpiocdev.AsOutput(v))
	if err != nil {
		panic(err)
	}
	// revert line to input on the way out.
	defer func() {
		l.Reconfigure(gpiocdev.AsInput)
		fmt.Printf("Input pin %s:%d\n", chip, offset)
		l.Close()
	}()
	values := map[int]string{0: "inactive", 1: "active"}
	fmt.Printf("Set pin %s:%d %s\n", chip, offset, values[v])

	// capture exit signals to ensure pin is reverted to input on exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	for {
		select {
		case <-time.After(500 * time.Millisecond):
			v ^= 1
			l.SetValue(v)
			fmt.Printf("Set pin %s:%d %s\n", chip, offset, values[v])
		case <-quit:
			return
		}
	}

}
