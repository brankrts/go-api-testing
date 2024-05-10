package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)
type NewModel struct {
	Data string
}

type DistributionRM struct {
	Data      *DistributionDataRM `json:"data"`
	Succeeded bool                `json:"succeeded"`
	Error     *ErrorModel         `json:"error"`
}
type DistributionDataRM struct {
	MassPaymentId   *string    `json:"massPaymentId"`
	MassPaymentType *int       `json:"massPaymentType"`
	InvitationId    *int       `json:"invitationId"`
	Id              *int       `json:"id"`
	CreatedAt       *time.Time `json:"createdAt"`
	Amount          *float64   `json:"amount"`
	Currency        *int       `json:"currency"`
	Fee             *float64   `json:"fee"`
	ResultinBalance *float64   `json:"resultinBalance"`
	Description     *string    `json:"description"`
}

type Notification struct {
	Count            int    `json:"time"`
	NotificationName string `json:"notificationName"`
}

type ValidionForPhone struct {
	PhoneNumber   string  `json:"phoneNumber"`
	Amount        float32 `json:"amount"`
	MassPaymentId string  `json:"massPaymentId"`
}

type DistributeForPhone struct {
	PhoneNumber   string  `json:"phoneNumber"`
	Amount        float32 `json:"amount"`
	MassPaymentId string  `json:"massPaymentId"`
}

type DistributeForTc struct {
	TurkishNationalId string  `json:"turkishNationalId"`
	Amount            float32 `json:"amount"`
	MassPaymentId     string  `json:"massPaymentId"`
}

type ValidionForTC struct {
	TurkishNationalId string  `json:"turkishNationalId"`
	Amount            float32 `json:"amount"`
	MassPaymentId     string  `json:"massPaymentId"`
}
type ValidationResponse struct {
	Succeeded bool        `json:"succeeded"`
	Error     *ErrorModel `json:"error"`
}
type ErrorModel struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func generateFakeDistributionData() *DistributionDataRM {
	massPaymentId := "123456789"
	massPaymentType := 1
	invitationId := 9876
	id := 54321
	createdAt := time.Now()
	amount := 100.50
	currency := 1
	fee := 5.0
	resultinBalance := amount - fee
	description := "sahte aciklama"

	fakeData := &DistributionDataRM{
		MassPaymentId:   &massPaymentId,
		MassPaymentType: &massPaymentType,
		InvitationId:    &invitationId,
		Id:              &id,
		CreatedAt:       &createdAt,
		Amount:          &amount,
		Currency:        &currency,
		Fee:             &fee,
		ResultinBalance: &resultinBalance,
		Description:     &description,
	}

	return fakeData
}
func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/notification/setnotification", func(c *gin.Context) {
		var notification *Notification
		c.BindJSON(&notification)
		fmt.Printf("Values : %v ", gin.H{"Count": notification.Count, "NotificationName": notification.NotificationName})
		c.JSON(200, gin.H{"Count": notification.Count, "NotificationName": notification.NotificationName})
	})

	r.POST("/validation/tckn", func(c *gin.Context) {
		var validation *ValidionForTC
		c.BindJSON(&validation)
		fmt.Printf("Values : %v ", gin.H{"TC": validation.TurkishNationalId, "Amount": validation.Amount})
		c.JSON(200, ValidationResponse{
			Succeeded: true,
			Error:     nil,
		})
	})

	r.POST("/validation/phone", func(c *gin.Context) {
		var validation *ValidionForPhone
		c.BindJSON(&validation)
		fmt.Printf("Values : %v ", gin.H{"Phone": validation.PhoneNumber, "Amount": validation.Amount})
		c.JSON(200, ValidationResponse{
			Succeeded: false,
			Error: &ErrorModel{
				Message: "Hata mesaji",
				Code:    0,
			},
		})
	})

	r.POST("/distribute/phone", func(c *gin.Context) {
		var distribute *DistributeForPhone
		c.BindJSON(&distribute)
		fmt.Printf("Values : %v ", gin.H{"Phone": distribute.PhoneNumber, "Amount": distribute.Amount})
		c.JSON(200, DistributionRM{
			Data:      generateFakeDistributionData(),
			Succeeded: true,
			Error:     nil,
		})
	})

	r.POST("/distribute/tckn", func(c *gin.Context) {
		var distribute *DistributeForTc
		c.BindJSON(&distribute)
		fmt.Printf("Values : %v ", gin.H{"Phone": distribute.TurkishNationalId, "Amount": distribute.Amount})
		c.JSON(200, DistributionRM{
			Data:      generateFakeDistributionData(),
			Succeeded: true,
			Error:     nil,
		})
	})
	return r

}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
