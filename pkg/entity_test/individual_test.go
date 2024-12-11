package entity_test

import (
	"testing"

	"github.com/pericles-luz/go-wellhub/pkg/entity"
	"github.com/stretchr/testify/require"
)

func TestIndividualMustNotValidateIfKeyIsEmailAndHasNoId(t *testing.T) {
	individual := entity.NewIndividual()
	individual.AddIdentifier(entity.IDENTIFIER_EMAIL)
	require.ErrorIs(t, individual.Validate(), entity.ErrEmailRequired)
}

func TestIndividualMustNotValidateIfKeyIsNationalIDAndHasNoId(t *testing.T) {
	individual := entity.NewIndividual()
	individual.AddIdentifier(entity.IDENTIFIER_NATIONAL_ID)
	require.ErrorIs(t, individual.Validate(), entity.ErrNationalIDRequired)
}

func TestIndividualMustNotValidateIfKeyIsEmployeeIDAndHasNoId(t *testing.T) {
	individual := entity.NewIndividual()
	individual.AddIdentifier(entity.IDENTIFIER_EMPLOYEE_ID)
	require.ErrorIs(t, individual.Validate(), entity.ErrEmployeeIDRequired)
}

func TestIndividualMustNotBeEqualIfKeyIsEmailAndEmailIsDiferent(t *testing.T) {
	individual1 := entity.NewIndividual()
	individual1.Email = "email1@teste.com"
	individual1.AddIdentifier(entity.IDENTIFIER_EMAIL)

	individual2 := entity.NewIndividual()
	individual2.Email = "email2@teste.com"
	individual2.AddIdentifier(entity.IDENTIFIER_EMAIL)

	require.False(t, individual1.Equals(individual2))
}

func TestIndividualMustNotBeEqualIfKeyIsNationalIDAndNationalIDIsDiferent(t *testing.T) {
	individual1 := entity.NewIndividual()
	individual1.NationalID = "123456789"
	individual1.AddIdentifier(entity.IDENTIFIER_NATIONAL_ID)

	individual2 := entity.NewIndividual()
	individual2.NationalID = "987654321"
	individual2.AddIdentifier(entity.IDENTIFIER_NATIONAL_ID)

	require.False(t, individual1.Equals(individual2))
}

func TestIndividualMustNotBeEqualIfKeyIsEmployeeIDAndEmployeeIDIsDiferent(t *testing.T) {
	individual1 := entity.NewIndividual()
	individual1.EmployeeID = "123456789"
	individual1.AddIdentifier(entity.IDENTIFIER_EMPLOYEE_ID)

	individual2 := entity.NewIndividual()
	individual2.EmployeeID = "987654321"
	individual2.AddIdentifier(entity.IDENTIFIER_EMPLOYEE_ID)

	require.False(t, individual1.Equals(individual2))
}

func TestIndividualMustGenerateMap(t *testing.T) {
	individual := NewIndividual()
	asMap := individual.ToMap()
	require.Equal(t, asMap["email"], individual.Email)
	require.Equal(t, asMap["national_id"], individual.NationalID)
	require.Equal(t, asMap["full_name"], individual.FullName)
	require.Equal(t, asMap["employee_id"], individual.EmployeeID)
	require.Equal(t, asMap["additional_fields"].(map[string]interface{})["department"], individual.AdditionalField.Department)
	require.Equal(t, asMap["additional_fields"].(map[string]interface{})["cost_center"], individual.AdditionalField.CostCenter)
	require.Equal(t, asMap["additional_fields"].(map[string]interface{})["office_zip_code"], individual.AdditionalField.OfficeZipCode)
	require.Equal(t, asMap["additional_fields"].(map[string]interface{})["payroll_id"], individual.AdditionalField.PayrollID)
	require.Equal(t, asMap["attributes"].(map[string]interface{})["eligible_to_payroll"], individual.Attributes.EligibleToPayroll)
}

func NewIndividual() *entity.Individual {
	individual := entity.NewIndividual()
	individual.Email = "teste1@teste.com"
	individual.NationalID = "123456789"
	individual.EmployeeID = "987654321"
	individual.AdditionalField.Department = "TI"
	individual.AdditionalField.CostCenter = "CC"
	individual.AdditionalField.OfficeZipCode = "12345678"
	individual.AdditionalField.PayrollID = "123456"
	individual.Attributes.EligibleToPayroll = true
	return individual
}
