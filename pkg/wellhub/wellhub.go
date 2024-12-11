package wellhub

import (
	"net/http"
	"time"

	"github.com/pericles-luz/go-rest/pkg/rest"
	"github.com/pericles-luz/go-wellhub/internal/config"
	"github.com/pericles-luz/go-wellhub/pkg/entity"
)

type Wellhub struct {
	rest *rest.Rest
}

func NewWellhub(configPath string) *Wellhub {
	configWellhub := config.NewWellhub(configPath)
	data := configWellhub.GetConfig()
	data["InsecureSkipVerify"] = true
	engine := rest.NewRest(data)
	token := rest.NewToken()
	token.SetKey(configWellhub.Token)
	token.SetValidity(time.Now().UTC().Add(time.Hour).Format("2006-01-02 15:04:05"))
	if !token.IsValid() {
		return nil
	}
	engine.SetToken(token)
	return &Wellhub{rest: engine}
}

func (w *Wellhub) AddIndividual(individual *entity.Individual) error {
	if err := individual.Validate(); err != nil {
		return err
	}
	individuals := entity.NewIndividuals()
	individuals.Add(individual)
	resp, err := w.rest.PostArray(individuals.ToMap(), w.rest.GetConfig("linkAPI")+"/eligibility/v1/employees/bulk-create")
	if err != nil {
		return err
	}
	if resp.GetCode() != http.StatusCreated {
		return entity.ErrIndividualNotCreated
	}
	return nil
}

func (w *Wellhub) DeleteIndividual(individual *entity.Individual) error {
	if err := individual.Validate(); err != nil {
		return err
	}
	individuals := entity.NewIndividuals()
	individuals.Add(individual)
	resp, err := w.rest.PostArray(individuals.ToMap(), w.rest.GetConfig("linkAPI")+"/eligibility/v1/employees/bulk-delete")
	if err != nil {
		return err
	}
	if resp.GetCode() != http.StatusNoContent {
		return entity.ErrIndividualNotCreated
	}
	return nil
}

func (w *Wellhub) GetIndividuals(key string) (*entity.Individuals, error) {
	resp, err := w.rest.Get(map[string]interface{}{
		"searchTerm": key,
		"deleted":    false,
	}, w.rest.GetConfig("linkAPI")+"/eligibility/v1/employees")
	if err != nil {
		return nil, err
	}
	if resp.GetCode() != http.StatusOK {
		return nil, entity.ErrIndividualNotCreated
	}
	individuals := entity.NewIndividuals()
	individuals.FromJSON([]byte(resp.GetRaw()))
	return individuals, nil
}

func (w *Wellhub) AddIndividualWithVerification(individual *entity.Individual) error {
	search, err := w.GetIndividuals(individual.KeyId())
	if err != nil {
		return err
	}
	if len(search.List()) > 0 {
		return nil
	}
	return w.AddIndividual(individual)
}
