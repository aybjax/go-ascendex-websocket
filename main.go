package main

import (
	"aybjax_ascendex_websocket/app"
	"aybjax_ascendex_websocket/wsapp"
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan app.BestOrderBook)
	exitCh := make(chan int)

	app := wsapp.New()
	err := app.Connection()

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Connected\n")

	err = app.SubscribeToChannel("USDT_BTC")

	if err != nil {
		log.Println(err)

		return
	}

	fmt.Printf("Subscribed\n")

	go func() {
		time.Sleep(time.Second * 5)
		app.Disconnect()
		close(ch)
		time.Sleep(time.Millisecond * 1500)
		exitCh <- 0
	}()

	go app.ReadMessagesFromChannel(ch)

	for {
		select {
		case msg, ok := <-ch:
			if ok {
				fmt.Printf("received is %#v\n", msg)
			}
		case <-exitCh:
			return
		}
	}
}
