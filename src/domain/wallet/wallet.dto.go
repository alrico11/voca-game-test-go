package wallet

import "github.com/gin-gonic/gin"

type WalletDepositDTO struct {
	Balance int `json:"balance" binding:"required"`
}

type WalletWithdrawalDTO struct {
	Withdrawal int `json:"withdrawal" binding:"required"`
}

type HistoryWalletQueryDTO struct {
	Filter    string `form:"filter" binding:"omitempty"`
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Sort      string `form:"sort" binding:"omitempty"`
	SortBy    string `form:"sortBy" binding:"omitempty"`
	StartDate string `form:"startDate" binding:"omitempty"`
	EndDate   string `form:"endDate" binding:"omitempty"`
}

func BindWalletDepositDTO(c *gin.Context) (*WalletDepositDTO, error) {
	var dto WalletDepositDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		return nil, err
	}
	return &dto, nil
}

func BindWalletWithdrawalDTO(c *gin.Context) (*WalletWithdrawalDTO, error) {
	var dto WalletWithdrawalDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		return nil, err
	}
	return &dto, nil
}

func BindHistoryWalletQueryDTO(c *gin.Context) (*HistoryWalletQueryDTO, error) {
	var dto HistoryWalletQueryDTO
	if err := c.ShouldBindQuery(&dto); err != nil {
		return nil, err
	}
	return &dto, nil
}
