package utils

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

const (
	FmtInfColorPrefix = "\033[1;34m"
	FmtWrnColorPrefix = "\033[1;33m"
	FmtColorEndLine   = "\033[0m\n"
)

func PingTcpAddress(address string, timeout time.Duration) error {
	const dialTimeout = 500 * time.Millisecond

	// remove scheme prefix
	if i := strings.Index(address, "://"); i != -1 {
		address = address[i+3:]
	}

	retryCount := int(timeout / dialTimeout)
	connected := false
	for i := 0; i < retryCount; i++ {
		conn, err := net.DialTimeout("tcp", address, dialTimeout)
		if err == nil {
			connected = true
		}
		if conn != nil {
			conn.Close()
		}

		if connected {
			break
		}
	}

	if !connected {
		return fmt.Errorf("TCP ping to %s failed after %d retry attempts with %v timeout", address, retryCount, dialTimeout)
	}

	return nil
}

func WaitForFileExists(filePath string, timeoutDur time.Duration) error {
	timeoutCh := time.After(timeoutDur)

	for {
		select {
		case <-timeoutCh:
			return fmt.Errorf("file %q did not appear after %v", filePath, timeoutDur)
		default:
			if _, err := os.Stat(filePath); err == nil {
				return nil
			}
		}
	}
}

func CheckExpectedErr(t *testing.T, expectedErr, receivedErr error) {
	require.NotNil(t, receivedErr, "receivedErr")

	expectedSdkErr, ok := expectedErr.(*sdkErrors.Error)
	require.True(t, ok, "not a SDK error: %T", expectedErr)

	require.True(t, expectedSdkErr.Is(receivedErr), "receivedErr / expectedErr: %v / %v", receivedErr, expectedErr)
}
