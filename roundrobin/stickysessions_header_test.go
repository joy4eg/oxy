package roundrobin

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

func TestBasicHeader(t *testing.T) {
	const headerName = "x-real-ip"
	a := testutils.NewResponder("a")
	b := testutils.NewResponder("b")

	defer a.Close()
	defer b.Close()

	fwd, err := forward.New()
	require.NoError(t, err)

	sticky := NewStickySessionHeader(headerName)
	require.NotNil(t, sticky)

	lb, err := New(fwd, EnableStickySession(sticky))
	require.NoError(t, err)

	err = lb.UpsertServer(testutils.ParseURI(a.URL))
	require.NoError(t, err)
	err = lb.UpsertServer(testutils.ParseURI(b.URL))
	require.NoError(t, err)

	proxy := httptest.NewServer(lb)
	defer proxy.Close()

	client := http.DefaultClient

	for i := 0; i < 10; i++ {
		req, err := http.NewRequest(http.MethodGet, proxy.URL, nil)
		require.NoError(t, err)
		req.Header.Set(headerName, "127.0.0."+strconv.Itoa(i))

		resp, err := client.Do(req)
		require.NoError(t, err)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		require.NoError(t, err)
		require.NotEmpty(t, body)
	}
}

func TestBasicHeaderEmpty(t *testing.T) {
	const headerName = "x-real-ip"
	a := testutils.NewResponder("a")
	b := testutils.NewResponder("b")

	defer a.Close()
	defer b.Close()

	fwd, err := forward.New()
	require.NoError(t, err)

	sticky := NewStickySessionHeader(headerName)
	require.NotNil(t, sticky)

	lb, err := New(fwd, EnableStickySession(sticky))
	require.NoError(t, err)

	err = lb.UpsertServer(testutils.ParseURI(a.URL))
	require.NoError(t, err)
	err = lb.UpsertServer(testutils.ParseURI(b.URL))
	require.NoError(t, err)

	proxy := httptest.NewServer(lb)
	defer proxy.Close()

	client := http.DefaultClient

	for i := 0; i < 10; i++ {
		req, err := http.NewRequest(http.MethodGet, proxy.URL, nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		require.NoError(t, err)
		require.NotEmpty(t, body)
	}
}

func TestBasicHeaderSame(t *testing.T) {
	const headerName = "x-real-ip"
	a := testutils.NewResponder("a")
	b := testutils.NewResponder("b")

	defer a.Close()
	defer b.Close()

	fwd, err := forward.New()
	require.NoError(t, err)

	sticky := NewStickySessionHeader(headerName)
	require.NotNil(t, sticky)

	lb, err := New(fwd, EnableStickySession(sticky))
	require.NoError(t, err)

	err = lb.UpsertServer(testutils.ParseURI(a.URL))
	require.NoError(t, err)
	err = lb.UpsertServer(testutils.ParseURI(b.URL))
	require.NoError(t, err)

	proxy := httptest.NewServer(lb)
	defer proxy.Close()

	client := http.DefaultClient

	var data string
	for i := 0; i < 10; i++ {
		req, err := http.NewRequest(http.MethodGet, proxy.URL, nil)
		req.Header.Set(headerName, "127.0.0.1")
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		if data == "" {
			data = string(body)
		} else {
			require.Equal(t, data, string(body))
		}
	}
}
