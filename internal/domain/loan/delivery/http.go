package delivery

import (
	"lender-service/package/di/usecases"
	"lender-service/package/structs"
	"lender-service/package/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoanHTTP struct{
	uc usecases.Usecases
	validate *validator.Validator
}
func (h LoanHTTP) UploadDocument(c echo.Context) error {
	ctx := c.Request().Context()
	if err := h.uc.Loan.UploadDocument(ctx); err != nil {
		return err
	}
	return nil
}

func (h LoanHTTP) LoanRequest(c echo.Context)  error {
	ctx := c.Request().Context()
	req := &structs.LoanRequest{}
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}	
	res, err := h.uc.Loan.LoanRequest(ctx, req)
	if  err != nil{
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h LoanHTTP) LoanValidation(c echo.Context)  error {
	ctx := c.Request().Context()
	loanID := c.Param("loan_id");
	req := structs.RequestLoanValidator{}
	req.LoanID = loanID
	errBind := c.Bind(&req)
	if errBind != nil {
		return errBind
	}
	resp, err := h.uc.Loan.ValidationLoan(ctx, req); 
	if err != nil{
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h LoanHTTP) GetAll(c echo.Context) error {
	resp, err := h.uc.Loan.ListLoan(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h LoanHTTP) LoanDetail(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.uc.Loan.LoanDetail(ctx, c.Param("id_loan"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h LoanHTTP) LoanDisbursement(c echo.Context)  error {
	ctx := c.Request().Context()
	loanID := c.Param("loan_id");
	req := structs.RequestLoanValidator{}
	req.LoanID = loanID
	errBind := c.Bind(&req)
	if errBind != nil {
		return errBind
	}
	resp, err := h.uc.Loan.LoanDisbursement(ctx, req); 
	if err != nil{
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func NewLoanHTTP(r *echo.Group, uc usecases.Usecases, v *validator.Validator)  {
	u := LoanHTTP{uc : uc, validate: v}
	loanRoute := r.Group("/loan")
	loanRoute.POST("", u.LoanRequest).Name = "loan.request"
	loanRoute.GET("/:id_loan", u.LoanDetail).Name = "loan.detail"
	loanRoute.GET("", u.GetAll).Name = "loan.all"
	loanRoute.POST("/:id_loan/approval", u.LoanValidation).Name = "loan.approval"
	loanRoute.POST("/:id_loan/disbursement", u.LoanDisbursement).Name = "loan.disbursement"
}