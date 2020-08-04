package roundrobin

import (
	"net/http"
	"net/url"

	"github.com/stathat/consistent"
)

// StickySessionHeader is a mixin for load balancers that implements layer 7 (http header) session affinity
type StickySessionHeader struct {
	name string
}

// NewStickySessionHeader creates a new StickySession based on given header
func NewStickySessionHeader(name string) SessionSticker {
	return &StickySessionHeader{name: name}
}

// GetBackend returns the backend URL stored in the sticky cookie, iff the backend is still in the valid list of servers.
func (s *StickySessionHeader) GetBackend(req *http.Request, servers []*url.URL) (*url.URL, bool, error) {
	value := req.Header.Get(s.name)
	if value == "" {
		// Not set
		return nil, false, nil
	}

	c := consistent.New()
	for _, u := range servers {
		c.Add(u.String())
	}

	server, err := c.Get(value)
	if err != nil {
		return nil, false, err
	}

	url, err := url.Parse(server)
	return url, true, err
}

// StickBackend does nothing
func (*StickySessionHeader) StickBackend(*url.URL, *http.ResponseWriter) {}
