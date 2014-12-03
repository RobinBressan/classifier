package main

import (
    "fmt"
    "./neural"
)

// func generateSynapses(size int) []*neural.Synapse {
//     synapses := make([]*neural.Synapse, size)
//     factory := new(neural.SynapseFactory)

//     for i := 0; i < size; i++ {
//         synapses[i] = factory.Create(1.0)
//     }

//     return synapses
// }

func main() {
    // factory := new(neural.NeuronFactory)

    compute := func (value float64) float64 {
        return value
    }

    synapseFactory := new(neural.SynapseFactory)
    neuronFactory := new(neural.NeuronFactory)

    generateSynapses := func (size int) []*neural.Synapse {
        synapses := make([]*neural.Synapse, size)

        for i := 0; i < size; i++ {
            synapses[i] = synapseFactory.Create(1.0)
        }

        return synapses
    }

    // Id 0
    starterNeuron1 := neuronFactory.Create(generateSynapses(1), generateSynapses(3), compute)

    // Id 1
    starterNeuron2 := neuronFactory.Create(generateSynapses(1), generateSynapses(3), compute)

    hiddenNeuron1Inputs := make([]*neural.Synapse, 2)
    hiddenNeuron1Inputs[0] = starterNeuron1.Outputs[0];
    hiddenNeuron1Inputs[1] = starterNeuron2.Outputs[0];

    // Id 2
    hiddenNeuron1 := neuronFactory.Create(hiddenNeuron1Inputs, generateSynapses(1), compute)

    hiddenNeuron2Inputs := make([]*neural.Synapse, 2)
    hiddenNeuron2Inputs[0] = starterNeuron1.Outputs[1];
    hiddenNeuron2Inputs[1] = starterNeuron2.Outputs[1];

    // Id 3
    hiddenNeuron2 := neuronFactory.Create(hiddenNeuron2Inputs, generateSynapses(1), compute)

    hiddenNeuron3Inputs := make([]*neural.Synapse, 2)
    hiddenNeuron3Inputs[0] = starterNeuron1.Outputs[2];
    hiddenNeuron3Inputs[1] = starterNeuron2.Outputs[2];

    // Id 4
    hiddenNeuron3 := neuronFactory.Create(hiddenNeuron3Inputs, generateSynapses(1), compute)

    terminalNeuronInputs := make([]*neural.Synapse, 3)
    terminalNeuronInputs[0] = hiddenNeuron1.Outputs[0];
    terminalNeuronInputs[1] = hiddenNeuron2.Outputs[0];
    terminalNeuronInputs[2] = hiddenNeuron3.Outputs[0];

    // Id 5
    terminalNeuron := neuronFactory.Create(terminalNeuronInputs, generateSynapses(1), compute)

    go starterNeuron1.Run()
    go starterNeuron2.Run()

    go hiddenNeuron1.Run()
    go hiddenNeuron2.Run()
    go hiddenNeuron3.Run()

    go terminalNeuron.Run()
    // go terminalNeuron2.Run()
    fmt.Println("Inject data")

    // go func() {
    starterNeuron1.Inputs[0].Channel <- 1.0
    starterNeuron2.Inputs[0].Channel <- 3.0
    // starterNeuron2.Inputs[0] <- 1.0
    // }()

    for i, output := range terminalNeuron.Outputs {
        fmt.Println("Output", i, <- output.Channel)
    }

    return
}
