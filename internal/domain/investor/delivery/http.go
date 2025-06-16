package delivery

import (
	"lender-service/package/di/usecases"
	"lender-service/package/structs"
	"lender-service/package/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvestorHTTP struct{
	uc usecases.Usecases
	validate *validator.Validator
}

func (h InvestorHTTP) GetAll(c echo.Context) error {
	resp, err := h.uc.Loan.ListLoan(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h InvestorHTTP) Invest(c echo.Context) error {
	loanID := c.Param("loan_id")
	req := structs.RequestInvestment{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	req.LoanID = loanID
	resp, err := h.uc.Investor.Invest(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func NewInvestorHTTP(r *echo.Group, uc usecases.Usecases, v *validator.Validator)  {
	u := InvestorHTTP{uc : uc, validate: v}
	InRoute := r.Group("/investments")
	InRoute.POST("/:loan_id", u.Invest).Name = "investor.request"
}