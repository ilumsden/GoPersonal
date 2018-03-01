package container

type Vector struct {
    data []int
    length int
}

func NewVector(data ...int) Vector {
    v := Vector{data, len(data)}
    return v
}

func CopyVector(cop *Vector) Vector {
    var data []int
    for _, e := range cop.data {
        data = append(data, e)
    }
    v := Vector{data, len(data)}
    return v
}

func (v *Vector) Resize(nsize int) {
    if nsize > v.length {
        extra := make([]int, nsize-v.length)
        v.data = append(v.data, extra...)
        v.length += nsize
    } else {
        tmp := v.data
        v.data = make([]int, nsize)
        for _, d := range tmp {
            v.data = append(v.data, d)
        }
        v.length = nsize
    }
}

func (v *Vector) PushBack(elem int) {
    v.Resize(v.length + 1)
    v.data[v.length - 1] = elem
}

func (v Vector) At(i int) int {
    return v.data[i]
}

func (v *Vector) Swap(ind1 int, ind2 int) {
    tmp := v.data[ind1]
    v.data[ind1] = v.data[ind2]
    v.data[ind2] = tmp
}

func (v Vector) Data() []int {
    return v.data
}

func (v Vector) Length() int {
    return v.length
}

func (v Vector) Empty() bool {
    return len(v.data) == 0
}

func (v Vector) Front() int {
    return v.data[0]
}

func (v Vector) Back() int {
    return v.data[v.length - 1]
}

func (v *Vector) Assign(ind int, elem int) {
    v.data[ind] = elem
}

func (v *Vector) Clear() {
    v.data = make([]int, 0)
    v.length = 0
}

func (v *Vector) PopBack() int {
    ret := v.data[v.length - 1]
    v.Resize(v.length - 1)
    return ret
}

type VectorIterator struct {
    p *Vector
    ind int
}

func (i *VectorIterator) Next() bool {
    point := i.p
    i.ind += 1
    if i.ind >= point.length {
        return false
    }
    return true
}

func (i *VectorIterator) Prev() bool {
    i.ind -= 1
    if i.ind < 0 {
        return false
    }
    return true
}

func (i VectorIterator) Deref() int {
    return (i.p).At(i.ind)
}

func (v Vector) Begin() VectorIterator {
    point := &v
    index := 0
    return VectorIterator{point, index}
}

func (v Vector) End() VectorIterator {
    point := &v
    index := v.length
    return VectorIterator{point, index}
}
