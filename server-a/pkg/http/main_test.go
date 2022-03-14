package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashptel/go-api-template/pkg/config"
)

func TestMain(m *testing.M) {
	config.LoadTestEnv()
	code := m.Run()
	os.Exit(code)
}
func Test_RunHttpServer(t *testing.T) {
	go RunHttpServer(context.Background())
	time.Sleep(time.Second)

	addr := fmt.Sprintf("http://localhost:%s", config.GetConfig().Port)

	req, err := http.NewRequest("GET", addr+"/api/health", nil)
	require.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
