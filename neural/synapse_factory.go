package neural

type SynapseFactory struct {}

func (_ *SynapseFactory) Create(weight float64) *Synapse {
    return &Synapse{ Channel: make(chan float64, 1), Weight: weight }
}
