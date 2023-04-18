package wsapp

import (
	"aybjax_ascendex_websocket/app"
	"reflect"
	"testing"
)

func TestIsTypeBBO(t *testing.T) {
	t.Run("It should be true if correct type", func(t *testing.T) {
		v := response{
			Type: "bbo",
		}

		if !v.isTypeBBO() {
			t.Error("Incorrectly identify as false")
		}
	})

	t.Run("It should be true if empty type", func(t *testing.T) {

		v := response{}

		if v.isTypeBBO() {
			t.Error("Incorrectly identify as false")
		}
	})

	t.Run("It should be true if incorrect type", func(t *testing.T) {

		v := response{
			Type: "sub",
		}

		if v.isTypeBBO() {
			t.Error("Incorrectly identify as false")
		}
	})
}

func TestGetBestOrderBook(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		v := response{
			Data: responseData{
				Bid: [2]string{"0.1", "0.2"},
				Ask: [2]string{"0.1", "0.2"},
			},
		}

		result, err := v.getBestOrderBook()
		correct := app.BestOrderBook{
			Ask: app.Order{
				Price:  0.1,
				Amount: 0.2,
			},
			Bid: app.Order{
				Price:  0.1,
				Amount: 0.2,
			},
		}

		if err != nil || !reflect.DeepEqual(result, correct) {
			t.Errorf("err=%#v; expected=%#v, actual=%#v", err, correct, result)
		}
	})
	t.Run("Int should be used parsed as float", func(t *testing.T) {
		v := response{
			Data: responseData{
				Bid: [2]string{"0.1", "0.2"},
				Ask: [2]string{"0.1", "0.2"},
			},
		}

		_, err := v.getBestOrderBook()

		if err != nil {
			t.Errorf("err=%#v", err)
		}
	})
	t.Run("Int should be used parsed as float", func(t *testing.T) {
		t.Run("0", func(t *testing.T) {
			v := response{
				Data: responseData{
					Bid: [2]string{"a", "0.2"},
					Ask: [2]string{"0.1", "0.2"},
				},
			}

			_, err := v.getBestOrderBook()

			if err == nil {
				t.Fail()
			}
		})
		t.Run("1", func(t *testing.T) {
			v := response{
				Data: responseData{
					Bid: [2]string{"0.1", "a"},
					Ask: [2]string{"0.1", "0.2"},
				},
			}

			_, err := v.getBestOrderBook()

			if err == nil {
				t.Fail()
			}
		})
		t.Run("2", func(t *testing.T) {
			v := response{
				Data: responseData{
					Bid: [2]string{"0.1", "0.2"},
					Ask: [2]string{"a", "0.2"},
				},
			}

			_, err := v.getBestOrderBook()

			if err == nil {
				t.Fail()
			}
		})
		t.Run("3", func(t *testing.T) {
			v := response{
				Data: responseData{
					Bid: [2]string{"0.1", "0.2"},
					Ask: [2]string{"0.1", "a"},
				},
			}

			_, err := v.getBestOrderBook()

			if err == nil {
				t.Fail()
			}
		})
	})
}
