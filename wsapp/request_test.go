package wsapp

import (
	"reflect"
	"testing"
)

func TestGetBBOSubMessage(t *testing.T) {
	t.Run("It should not fail", func(t *testing.T) {
		expected := &request{
			Op: "sub",
			Ch: "bbo:BTC/USDT",
		}

		r, err := request{}.getBBOSubMessage("USDT_BTC")

		if err != nil || !reflect.DeepEqual(expected, r) {
			t.Errorf("err=%#v, expected=%#v, actual=%#v", err, expected, r)
		}
	})

	t.Run("It should fail", func(t *testing.T) {
		_, err := request{}.getBBOSubMessage("USDT/BTC")

		if err == nil {
			t.Fail()
		}
	})
}

func TestGetPingMessage(t *testing.T) {
	expected := &request{
		Op: "ping",
	}

	actual := request{}.getPingMessage()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected=%#v, actual=%#v", expected, actual)
	}
}
