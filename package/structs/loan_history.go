package structs

import (
	"lender-service/internal/constant"
	"time"
)

// LoanHistory
type LoanHistory struct {
    ID int64 `json:"id gorm:primaryKey;autoIncrement"`
    LoanId string `json:"loan_id"`
    ApprovalPictureURL string `json:"approval_picture_url"`
    SignedAgreementURL string `json:"signed_agreement_url"`
    EmployeeId string `json:"employee_id"`
    ApprovalDate string `json:"approval_date"`
    DisbursementDate string`json:"disbursement_date"`
    State int32 `json:"state"`
    Comments string `json:"comments"`
    CreatedAt time.Time `json:"created_at"`
}

func (p RequestLoanValidator) NewLoanValidator(state int) LoanHistory {
    loan := LoanHistory{
        LoanId: p.LoanID,
        State: int32(state),
        EmployeeId: p.EmployeeID,
        Comments: "",
        CreatedAt: time.Now(),
    }
    if state == constant.Disbursed {
        loan.SignedAgreementURL = p.SignedAgreementURL
        loan.DisbursementDate = time.Now().Truncate(24 * time.Hour).String()
    } else {
        loan.ApprovalPictureURL = p.ApprovalPictureURL
        loan.ApprovalDate = time.Now().Truncate(24 * time.Hour).String()
    }
    
    return loan
}