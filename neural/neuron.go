package neural

import (
    "reflect"
    "fmt"
)

type Neuron struct {
    Inputs []chan float64
    Outputs []chan float64
    Compute func([]float64) float64
}

func (n Neuron) dispatch(value float64) {
    for _, output := range n.Outputs {
        output <- value
    }
}

func (n Neuron) Run() {
    cases := make([]reflect.SelectCase, len(n.Inputs))
    for i, input := range n.Inputs {
        cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(input)}
    }

    values := make([]float64, len(n.Inputs))
    indexes := IndexedBuffer{ data: make([]bool, len(n.Inputs)) }

    fmt.Println(++count, " neurons running...")
    for {
        fmt.Println("Waiting")
        chosen, value, _ := reflect.Select(cases)
        values[chosen] = value.Float()
        indexes.set(chosen, true)
        fmt.Println("Received", values[chosen])
        if indexes.isFull() {
            n.dispatch(n.Compute(values));
            indexes.truncate()
        }
    }
}
