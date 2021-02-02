package oauth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")

	os.Exit(m.Run())
}

func TestOauthConstants(t *testing.T) {
	assert.EqualValues(t, "X-Public", headerXPublic)
	assert.EqualValues(t, "X-Client-Id", headerXClientId)
	assert.EqualValues(t, "X-Caller-Id", headerXCallerId)
	assert.EqualValues(t, "access_token", paramAccessToken)
}

func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}

func TestIsPublicNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.False(t, IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t, IsPublic(&request))
}

func TestGetCallerIdNilRequest(t *testing.T) {
	assert.EqualValues(t, 0, GetCallerId(nil))
}

func TestGetCallerInvalidCallerFormat(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add("X-Caller-Id", "someText")
	assert.EqualValues(t, 0, GetCallerId(&request))
}

func TestGetCallerNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add("X-Caller-Id", "134")
	assert.EqualValues(t, 134, GetCallerId(&request))
}

func TestGetClientIdNilRequest(t *testing.T) {
	assert.EqualValues(t, 0, GetClientId(nil))
}

func TestGetClientInvalidCallerFormat(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add("X-Client-Id", "someText")
	assert.EqualValues(t, 0, GetClientId(&request))
}

func TestGetClientNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add("X-Client-Id", "134")
	assert.EqualValues(t, 134, GetClientId(&request))
}

func TestCleanRequest(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add("X-Client-Id", "123")
	request.Header.Add("X-Caller-Id", "134")
	cleanRequest(&request)
	assert.EqualValues(t, "", request.Header.Get("X-Client-Id"))
	assert.EqualValues(t, "", request.Header.Get("X-Caller-Id"))
}
