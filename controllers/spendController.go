/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package controllers

import (
	"fetch-rewards-backend/datastore"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleSpend(c *gin.Context) {
	q := datastore.Tq
	old := q.GetOldestTransaction()
	fmt.Println(old.Payer)
}
