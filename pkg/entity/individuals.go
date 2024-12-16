package entity

import (
	"encoding/json"
)

type Individuals struct {
	Individuals []*Individual `json:"items"`
}

func NewIndividuals() *Individuals {
	return &Individuals{}
}

func (i *Individuals) Add(individual *Individual) {
	if individual == nil {
		return
	}
	if err := individual.Validate(); err != nil {
		return
	}
	for _, i := range i.Individuals {
		if i.Equals(individual) {
			return
		}
	}
	i.Individuals = append(i.Individuals, individual)
}

func (i *Individuals) ToJSON() ([]byte, error) {
	return json.Marshal(i.Individuals)
}

func (i *Individuals) FromJSON(data []byte) error {
	return json.Unmarshal(data, &i)
}

func (i *Individuals) ToMap() []map[string]interface{} {
	var individuals []map[string]interface{}
	for _, individual := range i.Individuals {
		if individual == nil {
			continue
		}
		if err := individual.Validate(); err != nil {
			continue
		}
		individuals = append(individuals, individual.ToMap())
	}
	return individuals
}

func (i *Individuals) List() []*Individual {
	return i.Individuals
}
