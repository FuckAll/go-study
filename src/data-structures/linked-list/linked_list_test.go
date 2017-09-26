package list

import "testing"

var testCase1 = []*Node{NewNode("test1"), NewNode("test2"), NewNode("test3")}
var testCase2 = []*Node{NewNode("test4"), NewNode("test5"), NewNode("test6")}

func TestListPrepend(t *testing.T) {
	l := new(List)
	for _, n := range testCase1 {
		l.Prepend(n)
	}

	result := []string{"test3", "test2", "test1"}
	compareValueString(t, l, result)
}

func TestListAppend(t *testing.T) {
	l := new(List)
	for _, n := range testCase1 {
		l.Append(n)
	}

	var count int
	for n := l.Head; n != nil; n = n.Next {
		if n.Value.(string) != testCase1[count].Value.(string) {
			t.Errorf("TestListPrepend: got n.Value %s want %s", n.Value.(string), testCase1[len(testCase1)-count].Value.(string))
		}
		count++
	}
}

func TestListInsert(t *testing.T) {
	l := new(List)
	for _, n := range testCase1 {
		l.Append(n)
	}
	err := l.Insert(5, "inseart1")
	if err == nil {
		t.Errorf("TestListInsert: got err==nil want err!=nil")
	}

	caseResult := []string{"inseart1", "test1", "inseart3", "test2", "inseart2", "test3"}

	l.Insert(0, "inseart1")
	l.Insert(3, "inseart2")
	l.Insert(2, "inseart3")

	compareValueString(t, l, caseResult)
}

func compareValueString(t *testing.T, l *List, result []string) {
	var count int
	for n := l.Head; n != nil; n = n.Next {
		if n.Value.(string) != result[count] {
			t.Errorf("compareString: got n.Value %s want %s", n.Value.(string), result[count])
		}
		count++
	}
}

func TestListFind(t *testing.T) {
	l := new(List)
	for _, n := range testCase1 {
		l.Append(n)
	}

	result1 := "test1"
	result2 := 1
	result3 := struct {
		name string
		age  int
	}{
		name: "San",
		age:  11,
	}

	if i, err := l.Find(NewNode(result1)); err != nil && i != 0 {
		t.Errorf("TestListFind result1: want %d get %d", 0, i)
	}

	if i, err := l.Find(NewNode(result2)); err != nil && i != -1 {
		t.Errorf("TestListFind result2: want %d get %d", -1, i)
	}

	if i, err := l.Find(NewNode(result3)); err != nil && i != -1 {
		t.Errorf("TestListFind result3: want %d get %d", -1, i)
	}
}

func TestListRemove(t *testing.T) {
	l := new(List)
	for _, n := range testCase1 {
		l.Append(n)
	}

	result1 := []string{"test1", "test2", "test3"}
	result2 := []string{"test1", "test2"}
	result3 := []string{"test1"}
	result4 := []string{}

	if err := l.Remove(3); err != nil {
		compareValueString(t, l, result1)
	}

	if err := l.Remove(2); err == nil {
		compareValueString(t, l, result2)
	}

	if err := l.Remove(1); err == nil {
		compareValueString(t, l, result3)
	}

	if err := l.Remove(0); err == nil {
		compareValueString(t, l, result4)
	}

	l.Show()
}

func TestListContact(t *testing.T) {
	l := new(List)
	for _, n1 := range testCase1 {
		l.Append(n1)
	}

	nextList := new(List)
	for _, n2 := range testCase2 {
		nextList.Prepend(n2)
	}

	l.Contact(nextList)
	result := []string{"test1",
		"test2",
		"test3",
		"test6",
		"test5",
		"test4"}
	compareValueString(t, l, result)
}
