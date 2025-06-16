package structs

type RequestLoanDisbursement struct {
	IDLoan  				 string `json:"loan_id"`
	SignedAgreementUrl       string `json:"signed_agreement_url"`
	FieldValidatorEmployeeID string `json:"field_validator_employee_id"`
	ApprovalDate             string `json:"approval_date"`
}