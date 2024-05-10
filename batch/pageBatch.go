package batch

import "fmt"

/*
page method
total : total number of data
pageSize
offset
*/

type PageBatch struct {
	handlePanic
	total    int
	pageSize int
	offset   int
}

func NewPageBatch(total int, pageSize int, offset int) *PageBatch {
	return &PageBatch{
		total:    total,
		pageSize: pageSize,
		offset:   offset,
	}
}

func (b *PageBatch) Total() int {
	return b.total
}

func (b *PageBatch) PageSize() int {
	return b.pageSize
}

func (b *PageBatch) Offset() int {
	return b.offset
}

func (b *PageBatch) Exec(fn func(pageIndex int, offset int) error) (err error) {
	defer b.handlePanicError(&err)
	pageIndex := 0
	for offset := b.offset; offset < b.total; offset += b.pageSize {
		fmt.Println("start_offset:", b.offset, "total:", b.total, "pageSize:", b.pageSize, "offset:", offset)
		if err := fn(pageIndex, offset); err != nil {
			return err
		}
		pageIndex++
	}
	return err
}
