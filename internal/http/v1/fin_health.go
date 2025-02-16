package v1

import (
	"encoding/json"
	"fmt"
	utility "github.com/wachrusz/Back-End-API/pkg/util"
	"net/http"
)

type DeltaResponse struct {
	Message    string  `json:"message"`
	Delta      float64 `json:"delta"`
	StatusCode int     `json:"status_code"`
}

// ExpenditureDeltaHandler calculates the expenditure delta for an authenticated user.
//
// @Summary Calculate expenditure delta
// @Description This endpoint allows authenticated users to calculate the expenditure delta, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} DeltaResponse "Successfully calculated expenditure delta"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating expenditure delta"
// @Security JWT
// @Router /fin_health/expenses/delta [get]
func (h *MyHandler) ExpenditureDeltaHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Getting expenditure delta...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("auth err"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.ExpenditureDelta(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := DeltaResponse{
		Message:    "Expenditure delta calculated successfully",
		Delta:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

type PropensityResponse struct {
	Message    string  `json:"message"`
	Propensity float64 `json:"expense_propensity"`
	StatusCode int     `json:"status_code"`
}

// ExpensePropensity calculates the expense propensity for an authenticated user.
//
// @Summary Calculate expense propensity
// @Description This endpoint allows authenticated users to calculate the expense propensity, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} PropensityResponse "Successfully calculated expense propensity"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating expenditure delta"
// @Security JWT
// @Router /fin_health/expenses/propensity [get]
func (h *MyHandler) ExpensePropensity(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Getting expense propensity...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("auth err"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.ExpensePropensity(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := PropensityResponse{
		Message:    "expense propensity calculated successfully",
		Propensity: result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

type RatioResponse struct {
	Message    string  `json:"message"`
	Ratio      float64 `json:"ratio"`
	StatusCode int     `json:"status_code"`
}

// LiquidFundRatioHandler calculates the liquid fund ratio for an authenticated user.
//
// @Summary Calculate liquid fund ratio
// @Description This endpoint allows authenticated users to calculate the liquid fund ratio, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} RatioResponse "Successfully calculated liquid fund ratio"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating liquid fund ratio"
// @Security JWT
// @Router /fin_health/savings/ratio/liquid [get]
func (h *MyHandler) LiquidFundRatioHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating liquid fund ratio...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.LiquidFundRatio(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := RatioResponse{
		Message:    "Liquid fund ratio calculated successfully",
		Ratio:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// IlliquidFundRatioHandler calculates the illiquid fund ratio for an authenticated user.
//
// @Summary Calculate illiquid fund ratio
// @Description This endpoint allows authenticated users to calculate the illiquid fund ratio, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} RatioResponse "Successfully calculated illiquid fund ratio"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating illiquid fund ratio"
// @Security JWT
// @Router /fin_health/savings/ratio/illiquid [get]
func (h *MyHandler) IlliquidFundRatioHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating illiquid fund ratio...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.IlliquidFundRatio(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := RatioResponse{
		Message:    "Illiquid fund ratio calculated successfully",
		Ratio:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// SavingToIncomeRatioHandler calculates the savings to income ratio for an authenticated user.
//
// @Summary Calculate savings to income ratio
// @Description This endpoint allows authenticated users to calculate the savings to income ratio, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} RatioResponse "Successfully calculated savings to income ratio"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating savings to income ratio"
// @Security JWT
// @Router /fin_health/savings/ratio/savings_to_income [get]
func (h *MyHandler) SavingToIncomeRatioHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating savings to income ratio...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.SavingsToIncomeRatio(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := RatioResponse{
		Message:    "Savings to income ratio calculated successfully",
		Ratio:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// SavingsDelta calculates the savings delta for an authenticated user.
//
// @Summary Calculate savings delta
// @Description This endpoint allows authenticated users to calculate the savings delta, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} DeltaResponse "Successfully calculated savings delta"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating savings delta"
// @Security JWT
// @Router /fin_health/savings/delta [get]
func (h *MyHandler) SavingsDelta(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating savings delta...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.SavingDelta(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := DeltaResponse{
		Message:    "Savings delta calculated successfully",
		Delta:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// InvestmentsToSavingsRatioHandler calculates the monthly investments to savings ratio for an authenticated user.
//
// @Summary Calculate monthly investments to savings ratio
// @Description This endpoint allows authenticated users to calculate the monthly investments to savings ratio, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} RatioResponse "Successfully calculated investments to savings"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating investments to savings ratio"
// @Security JWT
// @Router /fin_health/investments/ratio/investments_to_savings [get]
func (h *MyHandler) InvestmentsToSavingsRatioHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating savings to income ratio...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.InvestmentsToSavingsRatio(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := RatioResponse{
		Message:    "Monthly investments to savings ratio calculated successfully",
		Ratio:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// InvestmentsToFundRatioHandler calculates the monthly investments to fund ratio for an authenticated user.
//
// @Summary Calculate monthly investments to fund ratio
// @Description This endpoint allows authenticated users to calculate the monthly investments to fund ratio, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} RatioResponse "Successfully calculated investments to fund"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating investments to fund ratio"
// @Security JWT
// @Router /fin_health/investments/ratio/investments_to_fund [get]
func (h *MyHandler) InvestmentsToFundRatioHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating fund to income ratio...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.InvestmentsToFundRatio(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := RatioResponse{
		Message:    "Monthly investments to fund ratio calculated successfully",
		Ratio:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// LoansToAssetsRatioHandler calculates the monthly loans to assets ratio for an authenticated user.
//
// @Summary Calculate monthly loans to assets ratio
// @Description This endpoint allows authenticated users to calculate the monthly loans to assets ratio, providing insight into their financial health.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} RatioResponse "Successfully calculated loans to assets ratio"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating loans to assets ratio"
// @Security JWT
// @Router /fin_health/loans/ratio/loans_to_assets [get]
func (h *MyHandler) LoansToAssetsRatioHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating loans to assets ratio...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.LoansToAssetsRatio(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := RatioResponse{
		Message:    "Monthly loans to assets ratio calculated successfully",
		Ratio:      result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// LoansPropensityHandler calculates the loans propensity for an authenticated user.
//
// @Summary Calculate loans propensity
// @Description This endpoint allows authenticated users to calculate their loans propensity, providing insight into their financial health and spending behavior.
// @Tags Financial Health
// @Accept  json
// @Produce  json
// @Success 200 {object} PropensityResponse "Successfully calculated loans propensity"
// @Failure 401 {object} jsonresponse.ErrorResponse "User not authenticated"
// @Failure 500 {object} jsonresponse.ErrorResponse "Server error while calculating loans propensity"
// @Security JWT
// @Router /fin_health/loans/propensity [get]
func (h *MyHandler) LoansPropensityHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Debug("Calculating loans propensity...")

	user, ok := utility.GetUserIDFromContext(r.Context())
	if !ok {
		h.errResp(w, fmt.Errorf("authentication error"), http.StatusUnauthorized)
		return
	}
	result, err := h.s.FinHealth.LoansPropensity(user)
	if err != nil {
		h.errResp(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := PropensityResponse{
		Message:    "Loans propensity calculated successfully",
		Propensity: result,
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
