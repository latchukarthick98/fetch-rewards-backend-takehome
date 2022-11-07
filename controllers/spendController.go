/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"fetch-rewards-backend/datastore"
)

type reqBody struct {
	Points int `json:"points" validate:"required"`
}

type resultItem struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

// Calculate the total points available for a user
func calculateTotalPoints(m map[string]int) int {
	total := 0
	for _, val := range m {
		total += val
	}
	return total
}

// Handler function for "POST" method for /spend route
func HandleSpend(c *gin.Context) {
	var body reqBody

	validate := validator.New()

	// Parse Request Body
	if err := c.BindJSON((&body)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid Input",
			"error": err,
		})
		return
	}

	// Validate Request Body
	err := validate.Struct(body)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	spend_points := body.Points
	m := datastore.Summary
	m1 := make(map[string]int)
	// datastore.Tq.PrintTQ()
	available_points := calculateTotalPoints(m)

	// Determine if the transaction is possible. Drop request if no sufficent funds (points)
	if spend_points > available_points {
		c.JSON(200, gin.H{
			"msg":     "INSUFFICIENT_BALANCE",
			"balance": available_points,
		})
		return
	}

	// Check if Transactions are available before deducting the points
	for datastore.Tq.Len() > 0 && spend_points > 0 {
		old := datastore.Tq.GetOldestTransaction()

		// Use the whole transaction
		if (spend_points - old.Points) >= 0 {

			spend_points -= old.Points
			m[old.Payer] -= old.Points
			m1[old.Payer] -= old.Points

			datastore.Tq.PopTransaction()
		} else {
			// Use partial points from a transaction (only what is necessary)
			comp := old.Points - spend_points
			m1[old.Payer] -= spend_points
			spend_points = 0
			m[old.Payer] = comp
			old.Points = comp
			datastore.Tq.Update(old, comp)

		}

	}

	result := []resultItem{}

	// Generate the response with payer and points as key-value pairs
	for key, val := range m1 {
		result = append(result, resultItem{
			Payer:  key,
			Points: val,
		})
	}

	c.JSON(http.StatusOK, result)
}
