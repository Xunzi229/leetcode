package main

import "fmt"

func main(){
    start1 := []int{0, 3}
    end1 := []int{0, 6}

    start2 := []int{0, 1}
    end2 := []int{0, 5}

    fmt.Println(intersection(start1, end1, start2, end2))
}

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
    if (start1[0] > start2[0] && end1[0] > end2[0]) || (start1[0] < start2[0] && end1[0] < end2[0]) {
        return []float64{}
    }

    var a1, b1 float64
    if end1[0] - start1[0] == 0 {
        b1 = 0
        if start1[0] == 0 {
            a1 = 0
        } else {
            a1 = float64(start1[1]) / float64(start1[0])
        }
    } else {
        a1 = float64(end1[1] - start1[1]) / float64(end1[0] - start1[0])
        b1 = float64(end1[1]) - a1 * float64(end1[0])
    }
    var a2, b2 float64
    if end2[0] - start2[0] == 0 {
        b2 = 0
        if start2[0] == 0 {
            a2 = 0
        } else {
            a2 = float64(start2[1]) / float64(start2[0])
        }
    } else {
        a2 = float64(end2[1] - start2[1]) / float64(end2[0] - start2[0])
        b2 = float64(end2[1]) - a2 * float64(end2[0])
    }

    fmt.Println(a1, b1)
    fmt.Println(a2, b2)

    if a1 == a2 && a1 == 0 {

    }

    if a1 == a2 && b1 == b2 {
        if start1[0] > start2[0] && start1[0] < end1[0] {
            return []float64{float64(start1[0]), float64(start1[1])}
        }
        if start1[0] > start2[0] && start1[0] > end1[0] {
            return []float64{float64(end1[0]), float64(end1[1])}
        }
        if start1[0] < start2[0] && start2[0] < end2[0] {
            return []float64{float64(start2[0]), float64(start2[1])}
        }
        if start1[0] < start2[0] && start2[0] > end2[0] {
            return []float64{float64(end2[0]), float64(end2[1])}
        }
    }

    if  (a1 != a2 && b1 == b2) || (a1 / b1 == a2 / b2) {
        return []float64{}
    }
    x1 := (b2 - b1) / (a1 - a2)
    y1 := a1 * x1 + b1

    return []float64{x1, y1}
}