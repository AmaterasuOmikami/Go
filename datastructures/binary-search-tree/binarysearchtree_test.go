package binary_search_tree

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type TestCase struct {
	param        int
	wantToReturn bool
	err          error
}

func TestInsert(t *testing.T) {
	testCases := []TestCase{
		{5, true, nil},
		{2, true, nil},
		{2, false, errors.New("item 2 already exists in tree")},
		{7, true, nil},
	}
	tree := newBinTree()
	for idx, tCase := range testCases {
		result, err := tree.Insert(tCase.param)
		if result != tCase.wantToReturn {
			t.Errorf("failed test case number %v\n want result %v, got %v", idx, tCase.wantToReturn, result)
		}
		if err != nil && err != tCase.err {
			t.Errorf("failed test case number %v\n want err %v got %v", idx, tCase.err.Error(), err.Error())
		}
	}
}

func TestSearch(t *testing.T) {
	testCases := []TestCase{
		{5, true, nil},
		{2, true, nil},
		{7, true, nil},
		{1, false, nil},
	}
	tree := newBinTree()
	wantError := errors.New("tree is empty")
	_, err := tree.Search(5)
	if err != nil && err.Error() != wantError.Error() {
		t.Errorf("failed test case with empty tree\n want err %v got %v", wantError.Error(), err.Error())
	}
	_, _ = tree.Insert(5)
	_, _ = tree.Insert(2)
	_, _ = tree.Insert(7)
	for idx, tCase := range testCases {
		result, err := tree.Search(tCase.param)
		if result != tCase.wantToReturn {
			t.Errorf("failed test case number %v\n want result %v, got %v", idx, tCase.wantToReturn, result)
		}
		if err != nil && err != tCase.err {
			t.Errorf("failed test case number %v\n want err %v got %v", idx, tCase.err.Error(), err.Error())
		}
	}
}

func TestDelete(t *testing.T) {
	testCases := []TestCase{
		{4, true, nil},
		{10, true, nil},
		{5, true, nil},
		{15, false, nil},
	}
	tree := newBinTree()
	_, _ = tree.Insert(5)
	_, _ = tree.Insert(4)
	_, _ = tree.Insert(2)
	_, _ = tree.Insert(1)
	_, _ = tree.Insert(3)
	_, _ = tree.Insert(10)
	_, _ = tree.Insert(9)
	_, _ = tree.Insert(11)
	_, _ = tree.Insert(7)
	_, _ = tree.Insert(8)
	_, _ = tree.Insert(12)

	for idx, tCase := range testCases {
		result := tree.Delete(tCase.param)
		if result != tCase.wantToReturn {
			t.Errorf("failed test case number %v\n want result %v, got %v", idx, tCase.wantToReturn, result)
		}
	}
	fmt.Println()
}

func TestIterWalk(t *testing.T) {
	tree := newBinTree()
	_, _ = tree.Insert(5)
	_, _ = tree.Insert(2)
	_, _ = tree.Insert(7)
	want := "2 5 7 "

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tree.IterBinWalk()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if want != string(out) {
		t.Errorf("Want: %v Out: %v", want, out)
	}
}

func TestRecursiveWalk(t *testing.T) {
	tree := newBinTree()
	_, _ = tree.Insert(5)
	_, _ = tree.Insert(2)
	_, _ = tree.Insert(7)
	want := "2 5 7 "

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tree.RecursiveBinWalk()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if want != string(out) {
		t.Errorf("Want: %v Out: %v", want, out)
	}
}
