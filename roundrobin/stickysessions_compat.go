package roundrobin

// NewStickySession creates a new StickySession based on cookie name
//
// Deprecated: Use NewStickySessionCookie instead
func NewStickySession(cookieName string) SessionSticker {
	return NewStickySessionCookie(cookieName)
}

// NewStickySessionWithOptions creates a new StickySession based on cookie with given options
//
// Deprecated: Use NewStickySessionCookieWithOptions instead
func NewStickySessionWithOptions(cookieName string, options CookieOptions) SessionSticker {
	return NewStickySessionCookieWithOptions(cookieName, options)
}
