package structs

type UpdateLoanRequest struct {
	IDLoan 		string `json:"id_loan"`
	IDUser 		string `json:"id_user"`
	ImageLoan 	string `json:"image_loan"`
	Status 		int8 `json:"status"`
}