/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */
package models

import (
	"container/heap"
	"fmt"
	"time"
)

// An Item is a single transaction.
type Item struct {
	Payer     string `json:"payer"`     // Payer name.
	Timestamp string `json:"timestamp"` // Timestamp in UNIX epoch
	Points    int    `json:"points"`    // Reward points
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

/** A TransactionQueue implements heap.Interface and holds Items.
*	It's a variant of priority queue, where we use timestamps instead of priorities
 */
type TransactionQueue []*Item

var isInitialized bool = false
var count int = 0

func (tq TransactionQueue) IsInitialized() bool {
	return isInitialized
}

func (tq TransactionQueue) GetCount() int {
	return count
}

func (tq TransactionQueue) SetCount(n int) {
	count = n
}

func (tq TransactionQueue) SetInitialzed(value bool) {
	isInitialized = value
}

func (tq TransactionQueue) Len() int { return len(tq) }

func (tq TransactionQueue) Less(i, j int) bool {
	// Condition to maintain the queue property, oldest timestamp first
	ts1, _ := time.Parse(time.RFC3339, tq[i].Timestamp)
	ts2, _ := time.Parse(time.RFC3339, tq[j].Timestamp)
	return ts1.Unix() < ts2.Unix()
}

func (tq TransactionQueue) Swap(i, j int) {
	tq[i], tq[j] = tq[j], tq[i]
	tq[i].index = i
	tq[j].index = j
}

func (tq *TransactionQueue) Push(x any) {
	n := len(*tq)
	item := x.(*Item)
	item.index = n
	*tq = append(*tq, item)
}

func (tq *TransactionQueue) Pop() any {
	old := *tq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*tq = old[0 : n-1]
	count -= 1
	return item
}

// update modifies the points of an Item (payer) in the queue.
func (tq *TransactionQueue) Update(item *Item, points int) {
	item.Points = points
	heap.Fix(tq, item.index)
}

// Add an item into transaction queue
func (tq *TransactionQueue) Insert(payer string, points int, ts string) {
	item := &Item{
		Payer:     payer,
		Points:    points,
		Timestamp: ts,
	}
	heap.Push(tq, item)
	count += 1
}

// Get the root of heap. In this case its the oldest item.
func (t1 TransactionQueue) GetOldestTransaction() *Item {
	// Constant time access O(1)
	if t1.GetCount() == 0 {
		return nil
	}
	// println(t1[0].Payer)
	return t1[0]
}

func (tq TransactionQueue) InitItem(payer string, points int, ts string) {
	item := &Item{
		Payer:     payer,
		Points:    points,
		Timestamp: ts,
		index:     0,
	}
	tq[0] = item
	heap.Init(&tq)
	// return item
}

func (tq *TransactionQueue) PopTransaction() *Item {
	return heap.Pop(tq).(*Item)
}

func (tq TransactionQueue) PrintTQ() {
	for tq.Len() > 0 {
		item := heap.Pop(&tq).(*Item)
		fmt.Printf("%s -> %s -> %d \n", item.Timestamp, item.Payer, item.Points)
	}
}

func (tq TransactionQueue) ToMap() map[string]int {
	m := make(map[string]int)
	for _, v := range tq {
		m[v.Payer] += v.Points
	}
	return m
}
