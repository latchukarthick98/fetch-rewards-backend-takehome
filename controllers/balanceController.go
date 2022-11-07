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

	// Return the Summary map which holds the balances by payers so far
	m := datastore.Summary

	c.JSON(200, m)

}
