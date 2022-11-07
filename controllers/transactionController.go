/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package controllers

import (
	"fetch-rewards-backend/datastore"
	"fetch-rewards-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertTransaction(c *gin.Context) {
	// tq := &datastore.Tq
	var item models.Item
	if err := c.BindJSON(&item); err != nil {
		return
	}
	fmt.Printf("Count: %d", datastore.Tq.GetCount())
	// Update the total spendable points by payer
	datastore.Summary[item.Payer] += item.Points
	if datastore.Summary[item.Payer] < 0 {
		datastore.Summary[item.Payer] = 0
		c.IndentedJSON(http.StatusCreated, item)
		return
	}
	if datastore.Tq.GetCount() == 0 {
		datastore.InitTQ(item.Payer, item.Points, item.Timestamp)
	} else {
		datastore.Tq.Insert(item.Payer, item.Points, item.Timestamp)
	}

	c.IndentedJSON(http.StatusCreated, item)
}

func TestConrol(c *gin.Context) {
	c.JSON(200, "hi")
}
