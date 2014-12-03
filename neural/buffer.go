package neural

type Buffer struct {
    data []float64
    indexes []bool
    onComplete func([] float64)
}

func (b *Buffer) set(index int, datum float64) {
    b.data[index] = datum
    b.indexes[index] = true

    for _, i := range b.indexes {
        if !i {
            return;
        }
    }

    b.onComplete(b.data)
    b.Truncate()
}

func (b *Buffer) Truncate() {
    for i := 0; i < len(b.indexes); i++ {
        b.indexes[i] = false
    }
}
