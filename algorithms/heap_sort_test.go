package algorithms

import (
	"encoding/json"
	"strconv"
	"testing"
)
import "fmt"

func TestBuildHeap(t *testing.T) {
	values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}
	BuildHeap(values, len(values))
	fmt.Println("大堆构建之后的结果：", values)
}

func TestHeapSort(t *testing.T) {
	values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}
	ret := HeapSort(values)
	fmt.Println("堆排序之后：", ret)
}

type yy struct {
	name string
}
type xx struct {
	id   int
	name yy
}

func xxxx() (*xx, error) {
	return nil, nil
}

// iota is a predeclared identifier representing the untyped integer ordinal
// number of the current const specification in a (usually parenthesized)
// const declaration. It is zero-indexed.

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	listNode := &ListNode{}
	borrow := false
	for ret := listNode; l1 != nil || l2 != nil || borrow; {
		var tmp int
		if borrow {
			tmp++
			borrow = false
		}
		if l1 != nil && l2 != nil {
			tmp += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			tmp += l1.Val
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		if tmp >= 10 {
			borrow = true
			tmp = tmp % 10
		}
		(*ret).Val = tmp

		if l1 != nil || l2 != nil || borrow {
			(*ret).Next = &ListNode{}
		} else {
			break
		}
		ret = ret.Next
	}
	return listNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var one int
	for i, tmp := 1, l1; ; tmp, i = tmp.Next, i*10 {
		one += tmp.Val * i
		if tmp.Next == nil {
			break
		}
	}

	var two int
	for i, tmp := 1, l2; ; tmp, i = tmp.Next, i*10 {
		two += tmp.Val * i
		if tmp.Next == nil {
			break
		}
	}

	var ret ListNode
	for i, listNode := one+two, &ret; i != 0; listNode, i = listNode.Next, i/10 {
		*listNode = ListNode{
			Val: i % 10,
		}
		if i/10 != 0 {
			(*listNode).Next = &ListNode{}
		} else {
			break
		}
	}
	return &ret
}
func xxx() (int, error) {
	return 1, nil
}

func x(a [3]int) {
	a[2] = 3
	//a[0] = 1
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var rev int
	for ; x != 0; {
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	fmt.Printf("%v", rev == x)
	if x == rev {
		return true
	}
	return false
}

func xxxxx() int {
	fmt.Println("xxxxx")
	return 0
}

func yyyy() int {
	defer func() {
		fmt.Println("yyyyy")
	}()
	return xxxxx()
}

func TestJson(t *testing.T) {
	type xxx struct {
		Ty   int
		Html string
	}
	xx := xxx{
		Ty:   1,
		Html: `<p>尊敬的客户您好！我们的在线人工客服将于除夕当天停止服务，正月初七恢复正常，如需帮助请与当地工作人员联系，您也可以点击<a onclick="window.location.href='https://uc.guazi.com/customer-service.html#/service-assistant?userMode=1'">用户服务中心</a>在线反馈或发邮件至<a onclick="window.location.href='mailto:kefuzhongxin@guazi.com'">kefuzhongxin@guazi.com</a>提交您的问题，我们会有专人查看，祝您春节愉快，再见。</p>`,
	}
	bytes, err := json.Marshal(xx)
	if err != nil {
		fmt.Printf("%v", err)
	}
	type yyy struct {
		Content string
	}

	yy := yyy{
		Content: string(bytes),
	}

	marshal, err := json.Marshal(&yy)
	fmt.Printf("%s", string(marshal))
}

//func isValid(s string) bool {
//	if strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
//		return isValid(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "{}", ""), "()", ""), "[]", ""))
//	} else {
//		return "" == s
//	}
//}

func isValid(s string) bool {
	var stack []int32
	m := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, c := range s {
		if len(stack) > 0 {
			if k, ok := m[stack[len(stack)-1]]; ok && k == c {
				if len(stack) == 1 {
					stack = nil
				} else {
					stack = stack[0 : len(stack)-1]
				}
				continue
			}
		}
		stack = append(stack, c)
	}
	return len(stack) == 0
}

func maxSubArray(nums []int) int {
	var ans, sum int
	for _, v := range nums {
		if sum > 0 {
			sum = sum + v
		} else {
			sum = v
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
func addBinary(a string, b string) string {
	stack := []int{}
	pop := func() int {
		popNode := stack[len(stack)-1 ]
		stack = stack[0 : len(stack)-1]
		return popNode
	}
	push := func(node int) {
		stack = append(stack, node)
	}
	push(1)
	i := pop()
	fmt.Println("xxxxx", i)

	//for i > 0 || i < -1 {
	//	fmt.Println("xxxxx")
	//	i = 1
	//
	//}
	axxx := []*ListNode{nil}
	fmt.Println(axxx)

	var ans string
	var ca int
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := ca
		if i >= 0 {
			sum = sum + int(a[i]-'0')
		} else {
			sum = sum + 0
		}

		if j >= 0 {
			sum = sum + int(b[j]-'0')
		} else {
			sum = sum + 0
		}
		ans = ans + strconv.Itoa(sum%2)
		ca = sum / 2
	}
	if ca == 1 {
		ans = ans + "1"
	}
	// 翻转字符串
	runes := []rune(ans)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestXXXX(t *testing.T) {

	binary := addBinary("11", "1")
	//binary := addBinary("1010", "1011")
	//binary := addBinary("1111", "1111")
	//fmt.Printf("%v", binary)
	fmt.Printf("%v", binary)
	//tmp := "11"
	//var x byte = tmp[0]
	//fmt.Printf("%v", x)
	//var x byte = '2'
	//fmt.Printf("%v", x + '1')
	//valid := isValid("{[]}")
	//fmt.Println(valid)
	//f1 := func() (int, int) {
	//	return 1, 2
	//}
	//
	//f2 := func() (int, int) {
	//	return 3, 4
	//}
	//
	//v1, v2 := f1()
	//fmt.Println(v1, v2)
	//
	//var v3 int
	//v2, v3 = f2()
	//
	//fmt.Println(v2)
	//fmt.Println(v3)
}

//palindrome := isPalindrome(121)
//fmt.Printf("%v", palindrome)
//a := [3]int{2, 3}
//x(a)
//fmt.Printf("%+v", a)
//for ; t != nil; {
//
//}

//var b bool
//var a int
//if b == false {
//	a, err := xxx()
//	if err != nil {
//	}
//	fmt.Printf("%d\n", a)
//}
//
//fmt.Printf("%d\n", a)

//tmp1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
//tmp2 := ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}
//v2 := addTwoNumbersV2(&tmp1, &tmp2)
//fmt.Printf("%+v\n", v2)
//
//var i = 1000000000000000000000000000001
//fmt.Printf("%+v\n", i)
//tmp := ListNode{
//	Val:  1,
//	Next: (*ListNode)(nil),
//}
//fmt.Printf("%+v\n", tmp.Next == (interface{})(nil))

//
//var ret []int
//ret = append(ret, )
//println(mutexLocked)
//println(mutexWoken)
//println(mutexStarving)
//println(mutexWaiterShift)
//tmp, err := xxxx()
//if err != nil {
//	t.Error(err)
//}
//
//fmt.Println(tmp.name.name)

//var y int64
//var x float64 = 111111111.2
//y = *(*int64)(unsafe.Pointer(&x))
//fmt.Println(y)
//fmt.Println(*(*uint64)(unsafe.Pointer(&x)))
//
//
//var pn int64
//pn = 100
//fmt.Println(*(*float64)(unsafe.Pointer(&pn)))
//var a *struct{} = nil
//var b []int = nil
//var c map[int]bool = nil
//var d chan string = nil
//var e func() = nil
//var f interface{} = nil

//fmt.Printf("%+v", a == b)
//var err error = nil
//fmt.Printf("%+v", err == nil)
//fmt.Printf("%+v", error(nil) == nil)
//fmt.Printf("%+v", interface{}(nil) == nil)
//fmt.Printf("%+v", interface{}(nil) == (*int)(nil))
//i := new(map[int]int)
//(*i)[1] = 1
//fmt.Printf("%+v", i)
//s := make([]xx, 3, 5)
//fmt.Println(s)
//
//n2 := new([]int)
//fmt.Println("result", n2, *n2)
//fmt.Println(len(*n2), cap(*n2))
//fmt.Println((*n2)[2])
//}
