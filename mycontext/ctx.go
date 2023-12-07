package mycontext

import (
    "context"
    "crypto/rand"
    "fmt"
    "io"
    "time"
)

func NewContext(ctx context.Context) context.Context {
    requestId, ok := ctx.Value(RequestId).(string)
    if !ok {
        requestId = GenRequestId()
    }

    return context.WithValue(context.Background(), RequestId, requestId)
}

func MakeContext(requestId string) context.Context {
    return context.WithValue(context.Background(), RequestId, requestId)
}

// GetRequestId 读取requestId
func GetRequestId(ctx context.Context) string {
    if ctx == nil {
        return GenRequestId()
    }

    requestId, ok := ctx.Value(RequestId).(string)
    if !ok {
        return GenRequestId()
    }

    if len(requestId) == 0 {
        return GenRequestId()
    }

    return requestId
}

const RequestId = "requestId"

func GenRequestId() string {
    return fmt.Sprintf("%d%s", time.Now().UnixNano(), randSerialNumber(4))
}

var serialNumberBytes = []byte(`0123456789`)

func randSerialNumber(max int) string {
    if max <= 0 {
        return ""
    }
    var bytes = make([]byte, max)
    n, err := io.ReadAtLeast(rand.Reader, bytes, max)
    if max != n || err != nil {
        return ""
    }
    for i := 0; i < len(bytes); i++ {
        bytes[i] = serialNumberBytes[int(bytes[i])%len(serialNumberBytes)]
    }
    return string(bytes)
}
