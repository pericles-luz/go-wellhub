package wellhub_test

import (
	"testing"

	"github.com/pericles-luz/go-wellhub/pkg/entity"
	"github.com/pericles-luz/go-wellhub/pkg/wellhub"
	"github.com/stretchr/testify/require"
)

func TestWellhubMustAddIndividual(t *testing.T) {
	t.Skip("use only if you need it")
	wellhub := wellhub.NewWellhub("../../config/gympass.sandbox.json")
	individual := NewIndividual()
	require.NoError(t, wellhub.AddIndividual(individual))
}

func TestWellhubMustDeleteIndividual(t *testing.T) {
	t.Skip("use only if you need it")
	wellhub := wellhub.NewWellhub("../../config/gympass.sandbox.json")
	individual := NewIndividual()
	require.NoError(t, wellhub.DeleteIndividual(individual))
}

func TestWellhubMustGetIndividual(t *testing.T) {
	t.Skip("use only if you need it")
	wellhub := wellhub.NewWellhub("../../config/gympass.sandbox.json")
	individual := NewIndividual()
	individuals, err := wellhub.GetIndividuals(individual.KeyId())
	require.NoError(t, err)
	require.Len(t, individuals.List(), 1)
}

func NewIndividual() *entity.Individual {
	individual := entity.NewIndividual()
	individual.AddIdentifier(entity.IDENTIFIER_EMAIL)
	individual.Email = "teste1@teste.com"
	return individual
}
