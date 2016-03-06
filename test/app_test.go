package flyingSquid_Test

import (
	"net/http"
	"testing"

	"github.com/e-r-w/flying-squid/lib"
	"github.com/jmcvetta/restclient"
)

func init() {
	fakeServer := flyingSquid.App{flyingSquid.FakeImageRepository{}, http.DefaultTransport}
	go fakeServer.CreateServer().RunOnAddr(":8080")
}

func TestInitServer(t *testing.T) {

	var expected []map[string]string
	_, err := restclient.Do(&restclient.RequestResponse{
		Url:    "http://localhost:8080/images/test",
		Result: &expected,
	})
	if err != nil {
		t.Fatal(err)
	} else {
		if expected[0]["name"] != "foo" {
			t.Fatal(expected)
		}
	}

}
