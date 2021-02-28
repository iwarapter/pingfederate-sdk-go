package pingfederate_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleServerSideTimout(t *testing.T) {

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond) //<- Any value > 20ms
		_, _ = io.WriteString(w, "")
		w.WriteHeader(http.StatusOK)
	})

	backend := httptest.NewServer(http.TimeoutHandler(handlerFunc, 20*time.Millisecond, "server timeout"))
	defer backend.Close()

	svc := serverSettings.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(backend.URL).WithDebug(true))
	_, _, err := svc.GetServerSettingsWithContext(context.Background())
	require.NotNil(t, err)
	assert.Equal(t, "503 Service Unavailable server timeout", err.Error())
}

func TestContextDeadlineExceeded(t *testing.T) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond) //<- Any value > 20ms
	})

	backend := httptest.NewServer(handlerFunc)
	defer backend.Close()

	svc := serverSettings.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(backend.URL).WithDebug(true))
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Millisecond)

	_, _, err := svc.GetServerSettingsWithContext(ctx)
	require.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("Get \"%s/serverSettings\": context deadline exceeded", backend.URL), err.Error())
}
