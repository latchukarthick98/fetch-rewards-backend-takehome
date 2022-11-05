/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */
package controllers

import (
	"fetch-rewards-backend/datastore"

	"github.com/gin-gonic/gin"
)

/*
*	Handles the balance route for "GET" method
 */
func GetBalance(c *gin.Context) {
	// q := datastore.Tq

	// if q.GetCount() == 0 {
	// 	c.JSON(404, "No transactions recorded!")
	// 	return
	// }

	m := datastore.Summary

	// jsonStr, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// } else {
	// 	fmt.Println(string(jsonStr))
	// }

	c.JSON(200, m)

}
