/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/04/2022
 */
package datastore

import (
	"fetch-rewards-backend/models"
)

var Tq models.TransactionQueue

var Summary map[string]int = make(map[string]int)

func InitTQ(payer string, points int, ts string) {
	Tq = make(models.TransactionQueue, 1)
	Tq.SetCount(1)
	Tq.InitItem(payer, points, ts)
	Tq.SetInitialzed(true)
}
