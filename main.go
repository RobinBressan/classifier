package main

import (
    "fmt"
    "./neural"
)

func generateChannels(size int) []chan float64 {
    channels := make([]chan float64, size)

    for i := 0; i < size; i++ {
        channels[i] = make(chan float64, 1)
    }

    return channels
}

func main() {
    // factory := new(neural.NeuronFactory)

    compute := func (values []float64) float64 {
        result := float64(0)

        for _, value := range values {
            result += value
        }

        return result
    }

    inputs := generateChannels(2)

    starterNeuron1Inputs := make([]chan float64, 1)
    starterNeuron1Inputs[0] = inputs[0];

    starterNeuron1 := neural.Neuron{ Inputs: starterNeuron1Inputs, Outputs: generateChannels(3), Compute: compute }

    starterNeuron2Inputs := make([]chan float64, 1)
    starterNeuron2Inputs[0] = inputs[1];

    starterNeuron2 := neural.Neuron{ Inputs: starterNeuron2Inputs, Outputs: generateChannels(3), Compute: compute }

    hiddenNeuron1Inputs := make([]chan float64, 2)
    hiddenNeuron1Inputs[0] = starterNeuron1.Outputs[0];
    hiddenNeuron1Inputs[1] = starterNeuron2.Outputs[0];

    hiddenNeuron1 := neural.Neuron{ Inputs: hiddenNeuron1Inputs, Outputs: generateChannels(1), Compute: compute }

    hiddenNeuron2Inputs := make([]chan float64, 2)
    hiddenNeuron2Inputs[0] = starterNeuron1.Outputs[1];
    hiddenNeuron2Inputs[1] = starterNeuron2.Outputs[1];

    hiddenNeuron2 := neural.Neuron{ Inputs: hiddenNeuron2Inputs, Outputs: generateChannels(1), Compute: compute }

    hiddenNeuron3Inputs := make([]chan float64, 2)
    hiddenNeuron3Inputs[0] = starterNeuron1.Outputs[2];
    hiddenNeuron3Inputs[1] = starterNeuron2.Outputs[2];

    hiddenNeuron3 := neural.Neuron{ Inputs: hiddenNeuron3Inputs, Outputs: generateChannels(1), Compute: compute }

    terminalNeuronInputs := make([]chan float64, 3)
    terminalNeuronInputs[0] = hiddenNeuron1.Outputs[0];
    terminalNeuronInputs[1] = hiddenNeuron2.Outputs[0];
    terminalNeuronInputs[2] = hiddenNeuron3.Outputs[0];

    terminalNeuron := neural.Neuron{ Inputs: terminalNeuronInputs, Outputs: generateChannels(1), Compute: compute }

    go starterNeuron1.Run()
    go starterNeuron2.Run()

    go hiddenNeuron1.Run()
    go hiddenNeuron2.Run()
    go hiddenNeuron3.Run()

    go terminalNeuron.Run()
    // go terminalNeuron2.Run()
    fmt.Println("Inject data")

    // go func() {
    inputs[0] <- 1.0
    inputs[1] <- 3.0
    // starterNeuron2.Inputs[0] <- 1.0
    // }()

    for i, output := range terminalNeuron.Outputs {
        fmt.Println("Output", i, <- output)
    }

    return
}
