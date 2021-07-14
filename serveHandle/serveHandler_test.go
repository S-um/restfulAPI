package serveHandle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewServeHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("index html", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewServeHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User Data")
}

func TestGetUsersInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewServeHandler())
	defer ts.Close()

	resp1, err1 := http.Post(ts.URL+"/users/9", "application/json",
		strings.NewReader(`{"name":"sir", "age":21, "email":"yoonms0101@gmail.com"}`))
	assert.NoError(err1)
	assert.Equal(http.StatusCreated, resp1.StatusCode)

	resp, err := http.Get(ts.URL + "/users/41")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "id")
	fmt.Println(string(data))
}

func TestCreateUsersInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewServeHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users/9", "application/json",
		strings.NewReader(`{"name":"sir", "age":21, "email":"yoonms0101@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "id")
	fmt.Println(string(data))
}
