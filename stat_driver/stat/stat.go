package stat

func Average(elems ...int) float64 {
    size := len(elems)
    var sum float64
    for _, e := range elems {
        sum += float64(e)
    }
    return (sum / float64(size))
}
