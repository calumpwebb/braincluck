package storage

type Storage interface {
	SetPtrPosition(position int)
	MovePtrRight()
	MovePtrLeft()
	IncrementCell()
	DecrementCell()
	GetValue() uint8
	GetTape() [MaxInMemoryTapeSize]uint8
	SetCellValue(value uint8)
}
