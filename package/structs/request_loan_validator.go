package structs

type RequestLoanValidator struct {
	LoanID  				 string `json:"loan_id"`
	ApprovalPictureURL       string `json:"approval_picture_url,omitempty"`
	EmployeeID string `json:"field_validator_employee_id"`
	SignedAgreementURL string `json:"signed_agreement_url,omitempty"`
}