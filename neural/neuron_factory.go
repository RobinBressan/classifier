package neural

type NeuronFactory struct {
    Id int
}

func (f *NeuronFactory) Create(inputs []*Synapse, outputs []*Synapse, compute func(float64) float64) *Neuron {
    f.Id++
    return &Neuron{ Id: f.Id, Inputs: inputs, Outputs: outputs, Compute: compute }
}
