package buffer

import (
	"io"
	"strconv"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 256)
		return (*Buffer)(&b)
	},
}

// Buffer - это адаптированный буфер для записи.
type Buffer []byte

// New возвращает экземпляр буфера из пула.
func New() *Buffer {
	buffer, _ := bufPool.Get().(*Buffer)
	return buffer
}

// Free высвобождает ресурсы и возвращает буфер в пул.
func (b *Buffer) Free() {
	const maxBufSize = 16 << 10
	if cap(*b) <= maxBufSize {
		b.Reset()
		bufPool.Put(b)
	}
}

// Len возвращает текущий размер буфера.
func (b *Buffer) Len() int {
	return len(*b)
}

// Reset сбрасывает буфер.
func (b *Buffer) Reset() {
	*b = (*b)[:0]
}

// Truncate отбрасывает хвостовую часть буфера на n элементов.
//
// Паникует, если n меньше нуля или больше текущего размера буфера.
func (b *Buffer) Truncate(n int) {
	if n == 0 {
		b.Reset()
		return
	}
	if n < 0 || n > len(*b) {
		panic("buffer: truncation out of range")
	}
	*b = (*b)[:len(*b)-n]
}

// String возвращает строковое содержимое буфера.
func (b *Buffer) String() string {
	return string(*b)
}

// Bytes возвращает копию содержимого буфера.
func (b *Buffer) Bytes() []byte {
	c := make([]byte, len(*b))
	copy(c, *b)
	return c
}

// WriteTo записывает в w содержимое буфера.
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	n0, err := w.Write(*b)
	return int64(n0), err
}

// Write записывает len(p) байт из p в буфер.
func (b *Buffer) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

// ReadByte возвращает байт из буфера под номером n.
//
// Паникует, если n меньше нуля или больше текущего размера буфера.
func (b *Buffer) ReadByte(n int) byte {
	if n < 0 || n > len(*b) {
		panic("buffer: index out of range")
	}
	return (*b)[n]
}

// WriteByte записывает входной байт в буфер.
func (b *Buffer) WriteByte(c byte) {
	*b = append(*b, c)
}

// writestring кодирует входную строку и записывает её в буфер.
//
// операция перебирает каждый байт в строке в поисках символов, которые
// нуждаются в кодировке json или utf-8. если строка не нуждается
// в кодировании, то строка добавляется полностью в буфер, иначе
// операция переключится на побайтовое кодирование.
func (b *Buffer) WriteString(s string) {
	*b = append(*b, s...)
}

// WriteBool преобразует входное значение в строку и добавляет её в буфер.
func (b *Buffer) WriteBool(c bool) {
	*b = strconv.AppendBool(*b, c)
}

// WriteInt64 преобразует входное значение в строку и добавляет её в буфер.
// Если x (исключая знак) короче width, результат дополняется нулями.
func (b *Buffer) WriteInt64(x int64, width int) {
	u := uint64(x)
	if x < 0 {
		*b = append(*b, '-')
		u = uint64(-x)
	}
	b.WriteUint64(u, width)
}

// WriteUint64 преобразует входное значение в строку и добавляет её в буфер.
// Если x короче width, результат дополняется нулями.
func (b *Buffer) WriteUint64(x uint64, width int) {
	var buf [20]byte
	i := len(buf) - 1
	for x >= 10 {
		q := x / 10
		buf[i] = byte('0' + x - q*10)
		i--
		x = q
	}
	buf[i] = byte('0' + x)
	for w := len(buf) - i; w < width; w++ {
		*b = append(*b, '0')
	}
	*b = append(*b, buf[i:]...)
}

// WriteFloat64 преобразует входное значение в строку и добавляет её в буфер.
func (b *Buffer) WriteFloat64(f float64, fmt byte) {
	*b = strconv.AppendFloat(*b, f, fmt, -1, 64)
}

// WriteDuration преобразует входное значение в строку формата "72h3m0.5s" и
// добавляет её в буфер.
func (b *Buffer) WriteDuration(d time.Duration) {
	var buf [64]byte
	w := len(buf)

	u := uint64(d)
	neg := d < 0
	if neg {
		u = -u
	}

	if u < uint64(time.Second) {
		// Особый случай: если длительность меньше секунды,
		// используются меньшие единицы измерения, например 1.2ms.
		var prec int
		w--
		buf[w] = 's'
		w--
		switch {
		case u == 0:
			*b = append(*b, "0s"...)
			return
		case u < uint64(time.Microsecond):
			// запись наносекунд.
			prec = 0
			buf[w] = 'n'
		case u < uint64(time.Millisecond):
			// запись микросекунд.
			prec = 3
			w--
			// символ µ занимает 2 байта.
			copy(buf[w:], "µ")
		default:
			// запись миллисекунд.
			prec = 6
			buf[w] = 'm'
		}
		w, u = fmtFrac(buf[:w], u, prec)
		w = fmtInt(buf[:w], u)
	} else {
		w--
		buf[w] = 's'

		w, u = fmtFrac(buf[:w], u, 9)

		// u здесь целочисленные секунды.
		w = fmtInt(buf[:w], u%60)
		u /= 60

		// u здесь целочисленные минуты.
		if u > 0 {
			w--
			buf[w] = 'm'
			w = fmtInt(buf[:w], u%60)
			u /= 60

			// u здесь целочисленные часы.
			if u > 0 {
				w--
				buf[w] = 'h'
				w = fmtInt(buf[:w], u)
			}
		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	*b = append(*b, buf[w:]...)
}

// fmtFrac форматирует дробь v/10**prec (например, ".12345") в конце buf,
// опуская конечные нули. Он также опускает десятичную точку, когда дробь
// равна 0. Он возвращает индекс, с которого начинаются выходные байты, и
// значение v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	// Опустите конечные нули вплоть до десятичной точки включительно.
	w := len(buf)
	isPrint := false
	for i := 0; i < prec; i++ {
		digit := v % 10
		isPrint = isPrint || digit != 0
		if isPrint {
			w--
			buf[w] = byte(digit) + '0'
		}
		v /= 10
	}
	if isPrint {
		w--
		buf[w] = '.'
	}
	return w, v
}

// fmtInt форматирует v в конце buf.
// Он возвращает индекс, с которого начинается вывод.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

// WriteTime преобразует входное значение в строку формата layout и
// добавляет её в буфер.
func (b *Buffer) WriteTime(t time.Time, layout string) {
	switch layout {
	case time.DateTime:
		b.writeDateTimeFormat(t)
	default:
		*b = t.AppendFormat(*b, layout)
	}
}

func (b *Buffer) writeDateTimeFormat(t time.Time) {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	b.WriteUint64(uint64(year), 4)
	*b = append(*b, '-')
	b.WriteUint64(uint64(month), 2)
	*b = append(*b, '-')
	b.WriteUint64(uint64(day), 2)
	*b = append(*b, ' ')
	b.WriteUint64(uint64(hour), 2)
	*b = append(*b, ':')
	b.WriteUint64(uint64(minute), 2)
	*b = append(*b, ':')
	b.WriteUint64(uint64(second), 2)
}
