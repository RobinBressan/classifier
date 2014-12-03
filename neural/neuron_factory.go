package neural

type NeuronFactory struct {}

func (_ *NeuronFactory) Create(inputSize int, outputSize int, compute func([]float64) float64) *Neuron {

    inputs := make([]chan float64, inputSize)
    outputs := make([]chan float64, outputSize)

    for i := 0; i < inputSize; i++ {
        inputs[i] = make(chan float64, 1)
    }

    for j := 0; j < outputSize; j++ {
        outputs[j] = make(chan float64, 1)
    }

    return &Neuron{ Inputs: inputs, Outputs: outputs, Compute: compute }
}
