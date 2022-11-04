/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */
package models

import (
	"container/heap"
)

// An Item is a single transaction.
type Item struct {
	payer     string // Payer name.
	timestamp int64  // Timestamp in UNIX epoch
	points    int    // Reward points
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

/** A TransactionQueue implements heap.Interface and holds Items.
*	It's a variant of priority queue, where we use timestamps instead of priorities
 */
type TransactionQueue []*Item

func (tq TransactionQueue) Len() int { return len(tq) }

func (tq TransactionQueue) Less(i, j int) bool {
	// Condition to maintain the queue property, oldest timestamp first
	return tq[i].timestamp < tq[j].timestamp
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
	return item
}

// update modifies the points of an Item (payer) in the queue.
func (tq *TransactionQueue) update(item *Item, points int) {
	item.points = points
	heap.Fix(tq, item.index)
}

// Get the root of heap. In this case its the oldest item.
func (t1 TransactionQueue) root() *Item {
	return t1[0]
}
