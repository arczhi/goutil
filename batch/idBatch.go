package batch

import (
	"fmt"
	"math/rand"
)

/*
bookmark method
total : total number of data
batchNum batch number of data
startId primary key id of the data
*/

type IdBatch struct {
	handlePanic
	total    int
	batchNum int
	startId  int
}

func NewIdBatch(total int, batchNum int, startId int) *IdBatch {
	return &IdBatch{
		total:    total,
		batchNum: batchNum,
		startId:  startId,
	}
}

func (b *IdBatch) Total() int {
	return b.total
}

func (b *IdBatch) BatchNum() int {
	return b.batchNum
}

func (b *IdBatch) StartId() int {
	return b.startId
}

func (b *IdBatch) Exec(fn func(roundIndex int) error) (err error) {
	defer b.handlePanicError(&err)
	round := b.calculateLoopRound()
	for i := 0; i < round; i++ {
		cursorId := b.startId + rand.Intn(1000) //游标ID，当前遍历到的数据的主键ID
		fmt.Println("total:", b.total, "batchNum:", b.batchNum, "round:", round, "startId:", b.startId, "current_round:", i+1, "cursorId:", cursorId)
		if err := fn(i); err != nil {
			return err
		}
	}
	return err
}

func (b *IdBatch) calculateLoopRound() int {
	if b.batchNum <= 0 {
		return 0
	}
	round := b.total / b.batchNum
	if b.total%b.batchNum != 0 {
		round++
	}
	return round
}
