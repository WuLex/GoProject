package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // 交换元素
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    // 示例数组
    arr := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println("未排序的数组:", arr)

    bubbleSort(arr)

    fmt.Println("排序后的数组:", arr)
}