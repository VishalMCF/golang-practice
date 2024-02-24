package main

import "testing"

type Student struct {
	name string
	age  int
}

var students []Student

func init() {
	students = []Student{
		{"Billa", 21},
		{"MahaBilla", 22},
		{"KhoonkharBilla", 23},
	}
}

func TestStudentDetails(t *testing.T) {

	test_students := []Student{
		{"Billa", 21},
		{"MahaBilla", 22},
		{"KhoonkharBilla", 23},
	}

	for i, student := range students {
		if student.name != test_students[i].name {
			t.Errorf("Expected '%s', but got '%s'", student.name, test_students[i].name)
		}
	}
}
