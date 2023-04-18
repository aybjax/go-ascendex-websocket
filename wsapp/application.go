package wsapp

import (
	"aybjax_ascendex_websocket/app"
	"errors"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

const _WEBSOCKET_CLOSED_MESSAGE = -1
const (
	_SCHEME = "wss"
	_HOST   = "ascendex.com:443"
	_PATH   = "/api/pro/v1/stream"
)
const _VERBOSE = true

type application struct {
	conn *websocket.Conn
}

func New() app.APIClient {
	return &application{}
}

func (a *application) Connection() error {
	if a.conn != nil {
		return errors.New("Websocket is already connected")
	}

	u := url.URL{Scheme: _SCHEME, Host: _HOST, Path: _PATH}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		return err
	}

	a.conn = c

	return nil
}

func (a *application) Disconnect() {
	if a.conn == nil {
		return
	}

	a.conn.Close()
	a.conn = nil

	return
}

func (a *application) SubscribeToChannel(symbol string) error {
	if a.conn == nil {
		return errors.New("Websocket is closed")
	}

	ch, err := request{}.getBBOSubMessage(symbol)

	if err != nil {
		return err
	}

	return a.conn.WriteJSON(ch)
}

func (a *application) ReadMessagesFromChannel(ch chan<- app.BestOrderBook) {
	for {
		var response response
		conn := a.conn

		if conn == nil {
			if _VERBOSE {
				log.Println("Websocket is closed")
			}

			return
		}

		err := conn.ReadJSON(&response)

		if err != nil {
			if _VERBOSE {
				log.Println("webSocket message error:", err)
			}

			return
		}

		if !response.isTypeBBO() {
			continue
		}

		bbo, err := response.getBestOrderBook()

		if err != nil {
			if _VERBOSE {
				log.Println("webSocket response error:", err)
			}

			return
		}

		ch <- bbo
	}
}

func (a *application) WriteMessagesToChannel() {
	if a.conn == nil {
		if _VERBOSE {
			log.Println("Websocket is closed")
		}

		return
	}

	a.conn.WriteJSON(request{}.getPingMessage())
}
