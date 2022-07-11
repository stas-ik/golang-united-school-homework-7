package coverage

import (
	"os"
	"testing"
	"time"
)

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

var (
	Pone = Person {
		firstName: "Annd",
		lastName: "Anikeyev",
		birthDay: time.Now(),
	}
	Ptwo = Person {
		firstName: "Boris",
		lastName: "Bobkov",
		birthDay: time.Now(),
	}
	Pthree = Person {
		firstName: "Dima",
		lastName: "Dudkin",
		birthDay: time.Now(),
	}
)

func TestSwapPeople(t *testing.T) {
	t.Parallel()
	people := People {Pone, Ptwo, Pthree}
	people.Swap(0, 2)

	if (Pone != people[2]) && (Pthree != people[0]) && (Ptwo == people[1]) {
		t.Errorf("persons didn't swap")
	}
}

func TestSwapPeopleWithSameIndex(t *testing.T) {
	t.Parallel()
	people := People {Pone, Ptwo, Pthree}
	people.Swap(0, 0)

	if !((Pone == people[0]) && (Ptwo == people[1]) && (Pthree == people[2])) {
		t.Errorf("persons should not to swap")
	}
}

func TestPeopleLen(t *testing.T) {
	t.Parallel()
	table := map[string]struct {
		Persons People
		Expected int
	} {
		"empty": {People{}, 0},
		"single person": {People{Pone}, 1},
		"many persons": {People{Pone, Ptwo, Pthree}, 3},
	}
	for name, tcase := range table {
		t.Run(name, func(t *testing.T) {
			if tcase.Persons.Len() != tcase.Expected {
				t.Errorf("Wrong len result for People")
			}
		})
	}
}

var ErrMsgForTestLessFunc = "Wrong result of Less function"

func TestPeopleLess(t *testing.T) {
	t.Parallel()

	date := time.Now()
	firstName := "a"
	lastName := "a"

	table := map[string]struct {
		Persons People
		Expected bool
	} {
		"same date, same firstName, diff lastName less": {
			People{
				Person {
					firstName: firstName,
					lastName: lastName + "b",
					birthDay: date,
				},
				Person {
					firstName: firstName,
					lastName: lastName + "c",
					birthDay: date,
				},
			},
			true,
		},
		"same date, diff firstName, diff lastName less": {
			People{
				Person {
					firstName: firstName + "b",
					lastName: lastName + "c",
					birthDay: date,
				},
				Person {
					firstName: firstName + "c",
					lastName: lastName + "b",
					birthDay: date,
				},
			},
			true,
		},
		"diff date, same firstName, same lastName less": {
			People{
				Person {
					firstName: firstName,
					lastName: lastName,
					birthDay: date.Add(time.Hour * 4),
				},
				Person {
					firstName: firstName,
					lastName: lastName,
					birthDay: date,
				},
			},
			true,
		},
	}
	for name, tcase := range table {
		t.Run(name, func (t *testing.T) {
			if tcase.Persons.Less(0, 1) != tcase.Expected {
				t.Error(ErrMsgForTestLessFunc)
			}
			if tcase.Persons.Less(1, 0) == tcase.Expected {
				t.Error(ErrMsgForTestLessFunc)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////

func cmpMatrix(m1, m2 Matrix) bool {
	if m1.rows != m2.rows {
		return false
	}
	if m1.cols != m2.cols {
		return false 
	}
	for i := range m1.data {
		if m1.data[i] != m2.data[i] {
			return false
		}
	}
	return true
}

func TestNewMatrixFunc(t *testing.T) {
	t.Parallel()
	table := map[string] struct {
		SrcStr string
		ExpMtrx *Matrix
		Positive bool
	} {
		"valid matrix" : {
			"1 2\n3 4",
			&Matrix {
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			true,
		},
		"Rows and Cols not equal": {
			"1\n3 4",
			nil,
			false,
		},
		"characters in src string": {
			"a a\n3 4",
			nil,
			false,
		},
	}
	for name, tcase := range table {
		t.Run(name, func(t *testing.T) {
			m, err := New(tcase.SrcStr)
			if tcase.Positive {
				if err != nil {
					t.Errorf("Unexpected err: %s", err)
				}
				if !cmpMatrix(*m, *tcase.ExpMtrx) {
					t.Error("Unexpected matrix result")
				}
			} else { // negative cases
				if m != nil {
					t.Error("Expected that matrix equals to nil")
				}
				if err == nil {
					t.Error("Expected that function returns some error") // check this error in production code :)
				}
			}
		})
	}
}

func TestRows(t *testing.T) {
	t.Parallel()
	m := Matrix {
		rows: 2,
		cols: 2,
		data: []int{1, 2, 3, 4},
	}
	rows := m.Rows()
	if rows[0][0] != m.data[0] || rows[0][1] != m.data[1] || rows[1][0] != m.data[2] || rows[1][1] != m.data[3] {
		t.Error("Wrond Rows result")
	}
}

func TestCols(t *testing.T) {
	t.Parallel()
	m := Matrix {
		rows: 2,
		cols: 2,
		data: []int{1, 2, 3, 4},
	}
	cols := m.Cols()
	if cols[0][0] != m.data[0] || cols[0][1] != m.data[2] || cols[1][0] != m.data[1] || cols[1][1] != m.data[3] {
		t.Error("Wrond Cols result")
	}
}

func TestSet(t *testing.T) {
	t.Parallel()
	
	m := Matrix {
		rows: 2,
		cols: 2,
		data: []int{1, 2, 3, 4},
	}

	table := map[string]struct {
		X, Y int
		Val int
		ExpRes bool
	} {
		"positive case for Matrix.Set": {
			0,
			0,
			0,
			true,
		},
		"negative case for Matrix.Set": {
			-1,
			0,
			0,
			false,
		},
	}

	for name, tcase := range table {
		t.Run(name, func(t *testing.T) {
			isSet := m.Set(tcase.X, tcase.Y, tcase.Val)
			if isSet != tcase.ExpRes {
				t.Error("Wrong bool result from Set method")
			}
			// check value in seting cell
		})
	}
}