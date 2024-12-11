package entity_test

import (
	"testing"

	"github.com/pericles-luz/go-wellhub/pkg/entity"
	"github.com/stretchr/testify/require"
)

func TestIndividualsMustAddOnlyOneIfAddingRepeatedly(t *testing.T) {
	individuals := entity.NewIndividuals()
	individual1 := NewIndividual()
	individual1.AddIdentifier(entity.IDENTIFIER_EMAIL)
	individuals.Add(individual1)
	individuals.Add(individual1)
	require.Len(t, individuals.List(), 1)
}

func TestIndividualsMustNotAddIfIndividualIsNil(t *testing.T) {
	individuals := entity.NewIndividuals()
	individuals.Add(nil)
	require.Len(t, individuals.List(), 0)
}

func TestIndividualsMustGenerateMapAsSlice(t *testing.T) {
	individuals := entity.NewIndividuals()
	individual1 := NewIndividual()
	individual1.AddIdentifier(entity.IDENTIFIER_EMAIL)
	individuals.Add(individual1)
	individual2 := NewIndividual()
	individual2.AddIdentifier(entity.IDENTIFIER_EMAIL)
	individual2.Email = "email8@teste.com"
	individuals.Add(individual2)
	require.Len(t, individuals.List(), 2)
	require.Len(t, individuals.ToMap(), 2)
}

func TestIndividualsMustGenerateFromJSON(t *testing.T) {
	individuals := entity.NewIndividuals()
	err := individuals.FromJSON([]byte(NewIndividualsJSON()))
	require.NoError(t, err)
	require.Len(t, individuals.List(), 3)
}

func NewIndividualsJSON() string {
	return `{
	"page_info":{
		"total_elements":3,
		"current_page":1,
		"total_pages":1,
		"offset_start":1,
		"offset_end":3,
		"page_size":50
	},
	"items":[{
		"id":"5c17bafb-3bd8-4513-9a53-53ceed90aece",
		"full_name":"",
		"email":"teste1@teste.com",
		"email_domain":"teste.com",
		"employee_id":null,
		"national_id":null,
		"eligible_item_id":"1f169966-0000-4a9a-bafa-8382227390d6",
		"created_at":"2024-12-10T23:37:35.835174Z",
		"deleted_at":null,
		"invitation_status":"NOT_SENT",
		"updated_at":"2024-12-10T23:37:35.835174Z",
		"additional_fields":{
			"department":null,
			"cost_center":null,
			"office_zip_code":null,
			"payroll_id":null
		},
		"attributes":{
			"eligible_to_payroll":false,
			"discount_subset_id":null
		},
		"custom_fields":null,
		"deleted":false
	},
	{
		"id":"deedf5ad-78b4-4504-81fa-7b17414c5099",
		"full_name":"",
		"email":"teste1@teste.com",
		"email_domain":"teste.com",
		"employee_id":null,
		"national_id":null,
		"eligible_item_id":"1f169966-0000-4a9a-bafa-8382227390d6",
		"created_at":"2024-12-10T23:34:26.313889Z",
		"deleted_at":"2024-12-10T23:34:30.505925Z",
		"invitation_status":"NOT_SENT",
		"updated_at":"2024-12-10T23:34:30.505925Z",
		"additional_fields":{
			"department":null,
			"cost_center":null,
			"office_zip_code":null,
			"payroll_id":null
		},
		"attributes":{
			"eligible_to_payroll":false,
			"discount_subset_id":null
		},
		"custom_fields":null,
		"deleted":true
	},
	{
		"id":"00f08cf5-475d-48a1-9f11-19aa60c05ec1",
		"full_name":"",
		"email":"teste1@teste.com",
		"email_domain":"teste.com",
		"employee_id":null,
		"national_id":null,
		"eligible_item_id":"1f169966-0000-4a9a-bafa-8382227390d6",
		"created_at":"2024-12-10T23:23:51.587584Z",
		"deleted_at":"2024-12-10T23:33:19.277303Z",
		"invitation_status":"NOT_SENT",
		"updated_at":"2024-12-10T23:33:19.277303Z",
		"additional_fields":{
			"department":null,
			"cost_center":null,
			"office_zip_code":null,
			"payroll_id":null
		},
		"attributes":{
			"eligible_to_payroll":false,
			"discount_subset_id":null
		},
		"custom_fields":null,
		"deleted":true
	}]}`
}
