package log

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/uuid"
	traceable_context "github.com/shkshariq/go-util/traceable_context"
)

var byt = bytes.NewBuffer(make([]byte, 0))
var lg = NewLog(WithLevel(INFO), WithStdOut(byt), WithFilePath(true), WithColors(true))
var pxLg = lg.Log()

var testCtx = traceable_context.WithUUID(uuid.New())

func BenchmarkInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pxLg.Info(`dd`)
		}
	})
}

func BenchmarkInfoContext(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pxLg.InfoContext(testCtx, `dd`)
		}
	})
}

func BenchmarkInfoParams(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pxLg.InfoContext(context.Background(), `dd`, 1, 2, 3)
		}
	})
}
