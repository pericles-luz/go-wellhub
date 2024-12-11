package entity

import "errors"

const (
	IDENTIFIER_EMAIL = 1 << iota
	IDENTIFIER_NATIONAL_ID
	IDENTIFIER_EMPLOYEE_ID
)

var (
	ErrEmailRequired        = errors.New("email is required")
	ErrNationalIDRequired   = errors.New("national id is required")
	ErrEmployeeIDRequired   = errors.New("employee id is required")
	ErrIdentifierRequired   = errors.New("at least one identifier is required")
	ErrIndividualNotCreated = errors.New("individual not created")
)

type Individual struct {
	identifiers     uint8
	Email           string `json:"email"`
	NationalID      string `json:"national_id"`
	FullName        string `json:"full_name"`
	EmployeeID      string `json:"employee_id"`
	AdditionalField struct {
		Department    string `json:"department"`
		CostCenter    string `json:"cost_center"`
		OfficeZipCode string `json:"office_zip_code"`
		PayrollID     string `json:"payroll_id"`
	} `json:"additional_fields"`
	Attributes struct {
		EligibleToPayroll bool `json:"eligible_to_payroll"`
	} `json:"attributes"`
}

// NewIndividual is a constructor for Individual

func NewIndividual() *Individual {
	return &Individual{}
}

func (i *Individual) AddIdentifier(identifier uint8) {
	i.identifiers |= identifier
}

func (i *Individual) HasIdentifier(identifier uint8) bool {
	return i.identifiers&identifier == identifier
}

func (i *Individual) Validate() error {
	if i.identifiers == 0 {
		return ErrIdentifierRequired
	}
	if i.HasIdentifier(IDENTIFIER_EMAIL) && i.Email == "" {
		return ErrEmailRequired
	}
	if i.HasIdentifier(IDENTIFIER_NATIONAL_ID) && i.NationalID == "" {
		return ErrNationalIDRequired
	}
	if i.HasIdentifier(IDENTIFIER_EMPLOYEE_ID) && i.EmployeeID == "" {
		return ErrEmployeeIDRequired
	}
	return nil
}

func (i *Individual) Equals(other *Individual) bool {
	if i.identifiers == 0 || other.identifiers == 0 {
		return false
	}
	if i.identifiers != other.identifiers {
		return false
	}
	if i.HasIdentifier(IDENTIFIER_EMAIL) && i.Email != other.Email {
		return false
	}
	if i.HasIdentifier(IDENTIFIER_NATIONAL_ID) && i.NationalID != other.NationalID {
		return false
	}
	if i.HasIdentifier(IDENTIFIER_EMPLOYEE_ID) && i.EmployeeID != other.EmployeeID {
		return false
	}
	return true
}

func (i *Individual) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":       i.Email,
		"national_id": i.NationalID,
		"full_name":   i.FullName,
		"employee_id": i.EmployeeID,
		"additional_fields": map[string]interface{}{
			"department":      i.AdditionalField.Department,
			"cost_center":     i.AdditionalField.CostCenter,
			"office_zip_code": i.AdditionalField.OfficeZipCode,
			"payroll_id":      i.AdditionalField.PayrollID,
		},
		"attributes": map[string]interface{}{
			"eligible_to_payroll": i.Attributes.EligibleToPayroll,
		},
	}
}

func (i *Individual) KeyId() string {
	if i.HasIdentifier(IDENTIFIER_EMAIL) {
		return i.Email
	}
	if i.HasIdentifier(IDENTIFIER_NATIONAL_ID) {
		return i.NationalID
	}
	if i.HasIdentifier(IDENTIFIER_EMPLOYEE_ID) {
		return i.EmployeeID
	}
	return ""
}
