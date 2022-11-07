/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/04/2022
 */
package datastore

import (
	"fetch-rewards-backend/models"
)

// Exposes the Transaction Queue to other packages
var Tq models.TransactionQueue

// Maintains the points for each payer, helps keep track of points available
var Summary map[string]int = make(map[string]int)

// Initializes the Transaction Queue with the very first transaction
func InitTQ(payer string, points int, ts string) {
	Tq = make(models.TransactionQueue, 1)
	Tq.SetCount(1)
	Tq.InitItem(payer, points, ts)
	Tq.SetInitialzed(true)
}

func Cleanup() {
	for k := range Summary {
		delete(Summary, k)
	}
	Tq.Clear()
	Tq.SetCount(0)
	Tq.SetInitialzed(false)
}
