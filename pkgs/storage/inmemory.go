package storage

const MaxInMemoryTapeSize = 10_000

type InMemoryStorage struct {
	pointerIndex int
	tape         [MaxInMemoryTapeSize]uint8
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (s *InMemoryStorage) boundPointerIndex() {
	// bound index to [0, MaxInMemoryTapeSize]
	if s.pointerIndex < 0 || s.pointerIndex >= MaxInMemoryTapeSize {
		s.pointerIndex = 0
	}
}

func (s *InMemoryStorage) MovePtrRight() {
	s.pointerIndex += 1
	s.boundPointerIndex()
}

func (s *InMemoryStorage) MovePtrLeft() {
	s.pointerIndex -= 1
	s.boundPointerIndex()
}

func (s *InMemoryStorage) IncrementCell() {
	s.tape[s.pointerIndex] += 1
}

func (s *InMemoryStorage) DecrementCell() {
	s.tape[s.pointerIndex] -= 1
}

func (s *InMemoryStorage) GetPtrPosition() int {
	return s.pointerIndex
}

func (s *InMemoryStorage) SetPtrPosition(index int) {
	s.pointerIndex = index
	s.boundPointerIndex()
}

func (s *InMemoryStorage) GetValue() uint8 {
	return s.tape[s.pointerIndex]
}

func (s *InMemoryStorage) GetTape() [MaxInMemoryTapeSize]uint8 {
	return s.tape
}

func (s *InMemoryStorage) SetCellValue(value uint8) {
	s.tape[s.pointerIndex] = value
}
