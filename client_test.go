package whatsonchain

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	t.Parallel()

	client := NewClient(NetworkTest, nil, nil)

	if len(client.UserAgent()) == 0 {
		t.Fatal("missing user agent")
	}
}

// TestNewClient_CustomHTTPClient test new client
func TestNewClient_CustomHTTPClient(t *testing.T) {
	t.Parallel()

	h := http.DefaultClient
	client := NewClient(NetworkTest, nil, h)
	assert.NotNil(t, client)
	assert.Equal(t, h, client.HTTPClient())
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client := NewClient(NetworkTest, nil, nil)
	fmt.Println(client.UserAgent())
	// Output:go-whatsonchain: v0.9.2
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClient(NetworkTest, nil, nil)
	}
}

// TestClientDefaultOptions tests setting ClientDefaultOptions()
func TestClientDefaultOptions(t *testing.T) {
	t.Parallel()

	options := ClientDefaultOptions()

	if options.UserAgent != defaultUserAgent {
		t.Fatalf("expected value: %s got: %s", defaultUserAgent, options.UserAgent)
	}

	if options.BackOffExponentFactor != 2.0 {
		t.Fatalf("expected value: %f got: %f", 2.0, options.BackOffExponentFactor)
	}

	if options.BackOffInitialTimeout != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffInitialTimeout)
	}

	if options.BackOffMaximumJitterInterval != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	}

	if options.BackOffMaxTimeout != 10*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 10*time.Millisecond, options.BackOffMaxTimeout)
	}

	if options.DialerKeepAlive != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.DialerKeepAlive)
	}

	if options.DialerTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.DialerTimeout)
	}

	if options.RequestRetryCount != 2 {
		t.Fatalf("expected value: %v got: %v", 2, options.RequestRetryCount)
	}

	if options.RequestTimeout != 10*time.Second {
		t.Fatalf("expected value: %v got: %v", 10*time.Second, options.RequestTimeout)
	}

	if options.TransportExpectContinueTimeout != 3*time.Second {
		t.Fatalf("expected value: %v got: %v", 3*time.Second, options.TransportExpectContinueTimeout)
	}

	if options.TransportIdleTimeout != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.TransportIdleTimeout)
	}

	if options.TransportMaxIdleConnections != 10 {
		t.Fatalf("expected value: %v got: %v", 10, options.TransportMaxIdleConnections)
	}

	if options.TransportTLSHandshakeTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.TransportTLSHandshakeTimeout)
	}
}

// TestClientDefaultOptions_NoRetry will set 0 retry counts
func TestClientDefaultOptions_NoRetry(t *testing.T) {
	options := ClientDefaultOptions()
	options.RequestRetryCount = 0
	client := NewClient(NetworkTest, options, nil)

	if client.UserAgent() != defaultUserAgent {
		t.Errorf("user agent mismatch")
	}
}
