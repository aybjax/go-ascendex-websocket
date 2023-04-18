package wsapp

import (
	"aybjax_ascendex_websocket/app"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestNew(t *testing.T) {
	expected := &application{}

	actual := New()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected=%#v, actual=%#v", expected, actual)
	}
}

func TestDisconnect(t *testing.T) {
	t.Run("Should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log(fmt.Sprint(r))
				t.Fail()
			}
		}()

		app := &application{}

		app.Disconnect()
	})
	t.Run("Should panic 0", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log(fmt.Sprint(r))
				if !strings.Contains(fmt.Sprint(r), "invalid memory address or nil pointer dereference") {
					t.Fail()
				}
			}
		}()

		app := &application{
			conn: &websocket.Conn{},
		}

		app.Disconnect()
	})

	t.Run("Should panic 1", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log(fmt.Sprint(r))
				if !strings.Contains(fmt.Sprint(r), "invalid memory address or nil pointer dereference") {
					t.Fail()
				}
			} else {
				t.Fail()
			}
		}()

		app := &application{
			conn: &websocket.Conn{},
		}

		app.Disconnect()
	})
}

func TestSubscribeToChannel(t *testing.T) {
	t.Run("err if closed", func(t *testing.T) {
		app := &application{}

		err := app.SubscribeToChannel("USDT_BTC")

		if err == nil || err.Error() != "Websocket is closed" {
			t.Fail()
		}
	})

	t.Run("err if from incorrect channel name", func(t *testing.T) {
		app := &application{
			conn: &websocket.Conn{},
		}

		err := app.SubscribeToChannel("_")
		_, act_err := request{}.getBBOSubMessage("_")

		if act_err == nil {
			panic("incorrect error")
		}

		if err == nil || err.Error() != act_err.Error() {
			t.Fail()
		}
	})
	t.Run("should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log(fmt.Sprint(r))
				if !strings.Contains(fmt.Sprint(r), "invalid memory address or nil pointer dereference") {
					t.Fail()
				}
			} else {
				t.Fail()
			}
		}()

		app := &application{
			conn: &websocket.Conn{},
		}

		app.SubscribeToChannel("USDT_BTC")
	})
}

func TestReadMessagesFromChannel(t *testing.T) {
	t.Run("Should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fail()
			}
		}()

		ch := make(chan<- app.BestOrderBook)

		appl := &application{}

		appl.ReadMessagesFromChannel(ch)
	})

	t.Run("Should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log(fmt.Sprint(r))
				if !strings.Contains(fmt.Sprint(r), "invalid memory address or nil pointer dereference") {
					t.Fail()
				}
			} else {
				t.Fail()
			}
		}()

		ch := make(chan<- app.BestOrderBook)

		appl := &application{
			&websocket.Conn{},
		}

		appl.ReadMessagesFromChannel(ch)
	})
}
