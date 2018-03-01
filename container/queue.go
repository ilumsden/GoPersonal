package container

type Queue struct {
    data Vector
}

func NewQueue(data ...int) Queue {
    q := Queue{NewVector(data...)}
    return q
}

func (q Queue) Empty() bool {
    return q.data.Empty()
}

func (q Queue) Size() int {
    return q.data.Length()
}

func (q Queue) Front() int {
    return q.data.Front()
}

func (q Queue) Back() int {
    return q.data.Back()
}

func (q *Queue) Push(elem int) {
    tmp := q.data.Data()
    data := make([]int, 0)
    data = append(data, elem)
    for _, d := range tmp {
        data = append(data, d)
    }
    q.data = NewVector(data...)
}

func (q *Queue) Pop() int {
    return q.data.PopBack()
}
