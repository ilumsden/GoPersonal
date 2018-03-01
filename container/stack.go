package container

type Stack struct {
    data Vector
}

func NewStack(data ...int) Stack {
    s := Stack{NewVector(data...)}
    return s
}

func (s Stack) Empty() bool {
    return s.data.Empty()
}

func (s Stack) Size() int {
    return s.data.Length()
}

func (s Stack) Top() int {
    return s.data.Back()
}

func (s *Stack) Push(elem int) {
    s.data.PushBack(elem)
}

func (s *Stack) Pop() int {
    return s.data.PopBack()
}
