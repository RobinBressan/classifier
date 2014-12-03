package neural

type BufferFactory struct {}

func (b *BufferFactory) create(size int, onComplete func([] float64)) *Buffer {
    buffer := new(Buffer)
    buffer.data = make([]float64, size)
    buffer.indexes = make([]bool, size)
    buffer.onComplete = onComplete

    return buffer
}
