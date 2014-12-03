package neural

type Buffer struct {
    data []float64
    indexes []bool
    threshold int
    level int
    onComplete func([] float64)
}

func (b *Buffer) set(index int, datum float64) {
    b.data[index] = datum

    if (!b.indexes[index]) {
        b.indexes[index] = true
        b.level += index
    }

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
        b.level = 0
    }
}
