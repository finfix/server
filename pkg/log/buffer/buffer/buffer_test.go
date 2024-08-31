package buffer_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"server/pkg/log/buffer/buffer"
)

func TestBuffer_new(t *testing.T) {
	n := testing.AllocsPerRun(5, func() {
		b := buffer.New()
		b.Free()
	})
	require.Empty(t, n)
}

func TestBuffer(t *testing.T) {
	b := buffer.New()
	t.Cleanup(b.Free)

	b.WriteByte('1')
	_, _ = b.Write([]byte(" 2 "))
	b.WriteString("3 ")
	b.WriteBool(true)
	b.WriteByte(' ')
	b.WriteFloat64(1e-07, 'g')
	b.WriteByte(' ')
	b.WriteInt64(-10, 0)
	b.WriteByte(' ')
	b.WriteUint64(123, 4)
	b.WriteByte(' ')
	b.WriteDuration(time.Second)
	b.WriteByte(' ')
	b.WriteTime(time.Time{}, time.RFC3339)
	b.WriteByte(' ')
	b.WriteTime(time.Time{}, time.DateTime)

	want := "1 2 3 true 1e-07 -10 0123 1s 0001-01-01T00:00:00Z 0001-01-01 00:00:00"

	require.Equal(t, len(want), b.Len())
	require.Equal(t, want[4], b.ReadByte(4))
	require.Equal(t, want, b.String())

	b.Truncate(20)
	require.Equal(t, len(want)-20, b.Len())
	require.Equal(t, want[:len(want)-20], b.String())

	var buf bytes.Buffer
	n, err := b.WriteTo(&buf)
	require.NoError(t, err)
	require.EqualValues(t, b.Len(), n)
	require.Equal(t, b.String(), buf.String())

	b.Truncate(0)
	require.Empty(t, b.Len())

	require.Panics(t, func() { b.ReadByte(-1) })
	require.Panics(t, func() { b.Truncate(-1) })
}

func TestBuffer_WriteDuration(t *testing.T) {
	b := buffer.New()
	t.Cleanup(b.Free)

	testCases := [...]time.Duration{
		0,
		-time.Second,
		100 * time.Nanosecond,
		100 * time.Microsecond,
		100 * time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
		time.Hour + 3*time.Minute,
		time.Hour + 5*time.Minute + 10*time.Second,
		time.Hour + time.Minute + time.Second + time.Millisecond,
		time.Hour + time.Minute + time.Second + time.Millisecond + time.Microsecond,
		time.Hour + time.Minute + time.Second + time.Millisecond + time.Microsecond + time.Nanosecond,
	}

	for _, d := range &testCases {
		b.WriteDuration(d)
		require.Equal(t, d.String(), b.String())
		b.Reset()
	}
}
