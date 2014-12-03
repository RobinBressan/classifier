package neural

type BufferFactory struct {}

func (b *BufferFactory) create(size int, onComplete func([] float64)) *Buffer {
    buffer := new(Buffer)
    buffer.data = make([]float64, size)
    buffer.indexes = make([]bool, size)
    buffer.level = 0

    buffer.threshold = 0
    for i := 0; i < size; i++ {
        buffer.threshold += i
    }

    buffer.onComplete = onComplete

    return buffer
}
