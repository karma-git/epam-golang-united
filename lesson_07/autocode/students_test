package coverage

import (
	"os"
	"time"
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var ppl People
var mtx *Matrix

func setupTestCase(){
	ppl = People{
		// 1980-07-31 15:59:05 +0000 UTC
		Person{firstName: "Harry", lastName: "Potter", birthDay: time.Date(1980, 07, 31, 15, 59, 05, 0, time.UTC)},
		Person{firstName: "Harry", lastName: "ARandomLastName", birthDay: time.Date(1980, 07, 31, 15, 59, 05, 0, time.UTC)},
		Person{firstName: "Tom", lastName: "Riddle", birthDay: time.Date(1926, 12, 31, 0, 0, 0, 0, time.UTC)},
		Person{firstName: "A", lastName: "Voldemort", birthDay: time.Date(1926, 12, 31, 0, 0, 0, 0, time.UTC)},
	}
	mtx = &Matrix{rows:2,cols:2,data: []int{1,2,3,4}}
}

func tearDownTestCase() {
	ppl = nil
	mtx = nil
}

// First Part

func TestPeopleLength(t *testing.T) {
	// t.Parallel() // NOTE: we couldn't use parallel test due to mutation of data in setup & teardown
	assert.Len(t, ppl, ppl.Len())
}

func TestPeopleSwap(t *testing.T) {
	// swap HP & V
	ppl.Swap(0, 3)
	assert.Equal(t, ppl[0].lastName, "Voldemort")
	assert.Equal(t, ppl[3].lastName, "Potter")  // this case doesn't increase coverage, it's just for visibility
}

func TestPeopleLess(t *testing.T) {
	// younger person return true
	assert.Truef(t, ppl.Less(1, 3), "HP is younger than V")
	assert.Falsef(t, ppl.Less(3, 1), "V is older than HP")  // this case doesn't increase coverage, it's just for visibility

	// same birthday, first name return true
	assert.Truef(t, ppl.Less(3, 2), "V wins TR")
	assert.Falsef(t, ppl.Less(2, 3), "TR lose V")

	// same birthday and name, last name return true
	assert.Truef(t, ppl.Less(1, 0), "ARand wins Potter")
	assert.Falsef(t, ppl.Less(0, 1), "Potter lose ARand")
}

// Second Part

func TestMatrixNew(t *testing.T) {
	m, _ := New("1 2 \n 3 4")
	assert.Equal(t, m, mtx)

	_, err := New("1 2 \n 3")
	assert.Error(t, err, "Rows need to be the same length")

	_, ErrAtoi := strconv.Atoi("a")
	_, err = New("a \n a")
	assert.Equal(t, err, ErrAtoi)
}

func TestRows(t *testing.T){
	assert.Equal(t, mtx.Rows(), [][]int{{1,2}, {3,4}})
}

func TestCols(t *testing.T){
	assert.Equal(t, mtx.Cols(), [][]int{{1,3}, {2,4}})
}

func TestSet(t *testing.T){
	assert.True(t, mtx.Set(1, 1, 7))
	assert.False(t, mtx.Set(2, 1, 7))  // this case doesn't increase coverage, it's just for visibility
	assert.Equal(t, mtx.data, []int{1,2,3,7})
}

// Entrypoint

func TestMain(m *testing.M) {
	setupTestCase()
	code := m.Run()
	tearDownTestCase()
	os.Exit(code)
}
