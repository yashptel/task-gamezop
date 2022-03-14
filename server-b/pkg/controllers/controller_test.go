package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashptel/server-b/pkg/config"
)

var srv *httptest.Server

func TestMain(m *testing.M) {
	config.LoadTestEnv()
	code := m.Run()
	os.Exit(code)
}

func runTestServer() {
	router := NewRouter(context.Background())
	srv = httptest.NewServer(router)
}

func Test_NewRouter(t *testing.T) {
	router := NewRouter(context.Background())
	require.NotNil(t, router)
}

func Test_Health(t *testing.T) {
	runTestServer()
	defer srv.Close()

	req, err := http.NewRequest("GET", srv.URL+"/api/health", nil)
	require.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	// negative test
	req, err = http.NewRequest("GET", srv.URL+"/api/dkljf", nil)
	require.NoError(t, err)

	client = &http.Client{}
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, 404, resp.StatusCode)
}

func Test_User(t *testing.T) {
	runTestServer()
	defer srv.Close()

	// negative test
	req, err := http.NewRequest("GET", srv.URL+"/api/reward", nil)
	require.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)

	// positive test
	req, err = http.NewRequest("GET", srv.URL+"/api/reward?id=jkh", nil)
	require.NoError(t, err)

	client = &http.Client{}
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

}
