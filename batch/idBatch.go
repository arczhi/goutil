package batch

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
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

// experiment method
func (b *IdBatch) GoExec(fn func(roundIndex int) error) error {
	round := b.calculateLoopRound()
	var wg sync.WaitGroup
	var errorChan = make(chan error, round)
	for i := 0; i < round; i++ {
		wg.Add(1)
		go func(i int) {
			var err error
			defer func() {
				b.handlePanicError(&err)
				if err != nil {
					errorChan <- err
				}
			}()
			cursorId := b.startId + rand.Intn(1000) //游标ID，当前遍历到的数据的主键ID
			fmt.Println("total:", b.total, "batchNum:", b.batchNum, "round:", round, "startId:", b.startId, "current_round:", i+1, "cursorId:", cursorId)
			if err := fn(i); err != nil {
				errorChan <- err
				return
			}
		}(i)
	}
	var err error
	for e := range errorChan {
		err = errors.New(fmt.Sprintf("%v%v\n", err, e))
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
