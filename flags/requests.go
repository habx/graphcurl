package flags

type Requests struct {
	URL       string
	UserAgent string
	Headers   map[string]string
	// Retry count
	Retry int
	// Retry delay in seconds
	RetryDelay int
}
