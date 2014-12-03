package neural

import (
    "fmt"
    "reflect"
)

type Neuron struct {
    Id      int
    Inputs  []*Synapse
    Outputs []*Synapse
    Compute func(float64) float64
}

/**
 * Compute the output result of the Neuron
 */
func (n Neuron) Activate(values[] float64) {
    inputValue := 0.0

    for i, value := range values {
        inputValue += value*n.Inputs[i].Weight
    }

    result := n.Compute(inputValue)

    fmt.Println("Neuron", n.Id, "returns", result, "for", values)

    for _, output := range n.Outputs {
        output.Channel <- result
    }
}

/**
 * Start the Neuron worker
 * It will listen from all input synapses and then send a result on all output synapses
 */
func (n Neuron) Run() {
    cases := make([]reflect.SelectCase, len(n.Inputs))
    for i, input := range n.Inputs {
        cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(input.Channel)}
    }

    bufferFactory := new(BufferFactory)
    buffer := bufferFactory.create(len(n.Inputs), func (values []float64) {
        // When all synapse have been activated we can activate the neuron
        n.Activate(values)
    })

    for {
        chosen, value, _ := reflect.Select(cases)
        fmt.Println("Neuron", n.Id - len(n.Inputs) - chosen, "send", value.Float(), "to Neuron", n.Id)
        buffer.set(chosen, value.Float())
    }
}
