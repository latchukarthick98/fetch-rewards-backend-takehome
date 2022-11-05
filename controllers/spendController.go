/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"fetch-rewards-backend/datastore"
	"fetch-rewards-backend/models"
)

type reqBody struct {
	Points int `json:"points"`
}

var q models.TransactionQueue = datastore.Tq

func HandleSpend(c *gin.Context) {
	var body reqBody

	if err := c.BindJSON((&body)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	spend_points := body.Points
	m := datastore.Summary
	m1 := make(map[string]int)
	// datastore.Tq.PrintTQ()
	fmt.Printf("Count: %d , Len: %d\n", q.GetCount(), datastore.Tq.Len())
	// for q.Len() > 0 {
	// 	item := q.PopTransaction()
	// 	fmt.Printf("%s -> %s -> %d (Len: %d) \n", item.Timestamp, item.Payer, item.Points, q.Len())
	// }
	for datastore.Tq.Len() > 0 && spend_points > 0 {
		old := datastore.Tq.GetOldestTransaction()
		fmt.Printf("Count: %d , Len: %d\n", datastore.Tq.GetCount(), datastore.Tq.Len())
		fmt.Printf("Top: %s -> %d -> %s", old.Payer, old.Points, old.Timestamp)
		if (spend_points - old.Points) >= 0 {
			// diff := old.Points - spend_points
			spend_points -= old.Points
			m[old.Payer] -= old.Points
			m1[old.Payer] -= old.Points

			fmt.Printf("Bal (%s): %d, %d \n", old.Payer, spend_points, old.Points)
			datastore.Tq.PopTransaction()
		} else {
			comp := old.Points - spend_points
			m1[old.Payer] -= spend_points
			spend_points = 0
			m[old.Payer] = comp
			old.Points = comp
			// q.Insert(old.Payer, old.Points, old.Timestamp)
			datastore.Tq.Update(old, comp)
			fmt.Printf("More case: Bal: %d, %d \n", spend_points, old.Points)
		}

	}

	c.IndentedJSON(http.StatusCreated, m1)
}
