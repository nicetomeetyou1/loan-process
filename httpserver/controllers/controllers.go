package controllers

import (
	"github.com/gin-gonic/gin"
)

type LoanApplicationController interface {
	Create(ctx *gin.Context)
	GetLoanApplications(ctx *gin.Context)
	ReapplyLoanApplication(ctx *gin.Context)
}

type CustomerController interface {
	GetCustomers(ctx *gin.Context)
	GetDetailCustomer(ctx *gin.Context)
	GetCustomerLoanApplications(ctx *gin.Context)
}

type PaymentInstallmentController interface {
	GetInstallmentByCustomerAndLoanRequest(ctx *gin.Context)
}
