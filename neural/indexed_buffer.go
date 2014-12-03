package neural

type IndexedBuffer struct {
    data []bool
}

func (b *IndexedBuffer) set(index int, datum bool) {
    b.data[index] = datum
}

func (b *IndexedBuffer) truncate() {
    for i := 0; i < len(b.data); i++ {
        b.data[i] = false
    }
}

func (b *IndexedBuffer) isFull() bool {
    for i := 0; i < len(b.data); i++ {
        if b.data[i] == false {
            return false
        }
    }

    return true
}
