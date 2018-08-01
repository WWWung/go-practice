package main

import "fmt"

func main() {
  slice := []float64{8,6,4,3,2,1}
  fmt.Println(avgSlice(slice))  //  4

  fmt.Println(sortNumber(5, 10)) // 5 10

  var stack1 stack
  stack1.push(1)
  fmt.Printf("%v\n", stack1)  //  {1, [1]}
  stack1.push(2)
  fmt.Printf("%v\n", stack1) //   {2, [1 2]}
  stack1.pop()
  fmt.Printf("%v\n", stack1)  //  {1, [1]}
  stack1.pop()
  fmt.Printf("%v\n", stack1)  //  {0, []}

  fmt.Println(fibonacci(8)) //  [1 1 2 3 5 8 13 21]

  fmt.Println(map1(square, 1,2,3,4))//  [1 4 9 16]

  s := []int{1,3,4,2,5}
  fmt.Println(max(s)) //  5
  fmt.Println(min(s)) //  1

  intSlice := []int{8,6,4,3,2,1}
  fmt.Println(sort(intSlice)) //  [1 2 3 4 6 8]

  fmt.Println(plusX(5)(4))   // 9

}

//  1.返回float64类型切片的均值
func avgSlice(slice []float64) float64 {
  var total float64 = 0
  for _, num := range slice {
    total += num
  }
  var rsl = total / float64(len(slice))
  return rsl
}

//  2.两个整数的自然排序
func sortNumber(x, y int) (int, int) {
  if x > y {
    return y, x
  } else {
    return x, y
  }
}

//  3.1创建一个固定大小保存整数的栈。它无需超出限制的增长，定义push函数--将数据放入栈，和pop函数--从栈中取得内容。栈应该是后进先出。
//  3.2 编写一个string方法将栈转化为字符串形式的表达，可以这样的方式打印整个栈  fmt.Printf("My stack %v\n", stack) 输出 [0:m][1:l][2:k]

type stack struct {
  i int
  data []int
}

func (s *stack) push(num int) {
  s.data = append(s.data, num)
  s.i ++
}

func (s *stack) pop() {
  s.data = s.data[0:s.i-1]
  s.i --
}

//  4.斐波那契数列
func fibonacci(c int) []int {
  slice := make([]int, c)
  slice[0] = 1
  slice[1] = 1
  for i:=2; i<c; i++ {
    slice[i] = slice[i-1] + slice[i-2]
  }
  return slice
}

//  5.编写一个函数，函数接收一个函数和一个列表作为参数，返回这个函数执行列表里每一项的结果的集合
func map1(f func(int) int, data ...int) []int {
  lens := len(data)
  slice := make([]int, lens)
  for i, s := range data {
    slice[i] = f(s)
  }
  return slice
}

func square(num int) int {
  return num * num
}

//  6.求int slice里的最小值和最大值
func min(slice []int) int {
  max := int(^uint(0) >> 1)
  for _, s := range slice {
    if (max > s) {
      max = s
    }
  }
  return max
}

func max(slice []int) int {
  min := ^int(^uint(0) >> 1)
  for _, s := range slice {
    if (min < s) {
      min = s
    }
  }
  return min
}

//  冒泡排序
func sort(slice []int) []int {
  var finished = true
  for i, _ := range slice {
    if (i == 0) {
      continue
    }
    if (slice[i] < slice[i-1]) {
      slice[i], slice[i-1] = slice[i-1], slice[i]
      finished = false
    }
  }
  var rsl []int
  if finished {
    rsl = slice
  } else {
    rsl = sort(slice)
  }
  return rsl
}

//  函数接收一个x int参数，并且返回一个函数，返回函数接收一个参数p int 并返回p+x
func plusX(x int) func(int) int {
  return func (p int) int {
    return x + p
  }
}
