/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package controllers

import (
	"fetch-rewards-backend/datastore"
	"fetch-rewards-backend/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// Handler function for "POST" method for /transaction route
func InsertTransaction(c *gin.Context) {
	// tq := &datastore.Tq
	var item models.Item

	// Parse request body
	if err := c.BindJSON((&item)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid Input",
			"error": err,
		})
		return
	}

	// Validate request body
	validate := validator.New()
	err := validate.Struct(item)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(400, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}
	// fmt.Printf("Count: %d", datastore.Tq.GetCount())
	// Update the total spendable points by payer
	datastore.Summary[item.Payer] += item.Points

	// Validate if the balance is negative, if so set points by payer to 0 (ensures the payer's points are non-negative)
	if datastore.Summary[item.Payer] < 0 {
		datastore.Summary[item.Payer] = 0
		c.IndentedJSON(http.StatusCreated, item)
		return
	}
	if datastore.Tq.GetCount() == 0 {
		// Initialize the Transaction Queue if it is the first transaction
		datastore.InitTQ(item.Payer, item.Points, item.Timestamp)
	} else {
		// Insert Transaction
		datastore.Tq.Insert(item.Payer, item.Points, item.Timestamp)
	}

	c.IndentedJSON(http.StatusCreated, item)
}

func TestConrol(c *gin.Context) {
	c.JSON(200, "hi")
}
