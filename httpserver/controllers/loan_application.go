package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"loan_process/exceptions"
	"loan_process/httpserver/request"
	"loan_process/httpserver/services"
	"strconv"
)

type loanApplicationController struct {
	svc services.LoanApplicationSvc
}

func NewLoanApplicationController(svc services.LoanApplicationSvc) LoanApplicationController {
	return &loanApplicationController{
		svc: svc,
	}
}

func (l *loanApplicationController) Create(ctx *gin.Context) {
	var req request.CreateLoanApplication
	err := ctx.ShouldBind(&req)
	if err != nil {
		exceptions.ValidationError(ctx, err)
		return
	}
	err = validator.New().Struct(req)
	if err != nil {
		exceptions.ValidationError(ctx, err)
		return
	}
	response := l.svc.CreateLoanApplication(ctx, &req)
	WriteJsonResponse(ctx, response)
}

func (l *loanApplicationController) GetLoanApplications(ctx *gin.Context) {
	response := l.svc.GetLoanApplication(ctx)
	WriteJsonResponse(ctx, response)
}

func (l *loanApplicationController) ReapplyLoanApplication(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		exceptions.ValidationError(ctx, err)
		return
	}
	var req request.ReapplyLoanApplication
	err = ctx.ShouldBind(&req)
	if err != nil {
		exceptions.ValidationError(ctx, err)
		return
	}
	err = validator.New().Struct(req)
	if err != nil {
		exceptions.ValidationError(ctx, err)
		return
	}
	response := l.svc.ReapplyLoanApplication(ctx, uint(customerId), &req)
	WriteJsonResponse(ctx, response)
}
