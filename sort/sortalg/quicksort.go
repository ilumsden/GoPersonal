package sortalg

import (
    "github/container"
)

func Quicksort(v *container.Vector, left int, right int) int {
    if left == right {
        return left
    }
    var middle int = ((right - left)/2) + left
    if v.At(middle) < v.At(left) {
        v.Swap(left, middle)
    }
    if v.At(right) < v.At(left) {
        v.Swap(left, right)
    }
    if v.At(right) < v.At(middle) {
        v.Swap(middle, right)
    }
    if right - left + 1 <= 3 {
        return middle
    }
    pindex := middle
    var pivot int = v.At(pindex)
    v.Swap(pindex, right - 1)
    var i int = left + 1
    var j int = right - 2
    for {
        for v.At(i) < pivot { i+=1 }
        for pivot < v.At(j) { j-=1 }
        if i >= j {
            break
        }
        v.Swap(i, j)
    }
    pindex = i
    v.Swap(pindex, right - 1)
    Quicksort(v, left, pindex - 1)
    Quicksort(v, pindex + 1, right)
    return pindex
}
