package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {
	type Grade struct {
		score    int
		expected string
	}
	grades := []Grade{
		{score: 80, expected: "A"},
		{score: 70, expected: "B"},
		{score: 60, expected: "C"},
		{score: 50, expected: "D"},
		{score: 40, expected: "F"},
	}

	for _, tc := range grades {

		t.Run(tc.expected, func(t *testing.T) {
			grade := services.CheckGrade(tc.score)
			expected := tc.expected

			// if grade != expected {
			// 	t.Errorf("got %v but expected %v", grade, expected)
			// }

			assert.Equal(t, expected, grade)

		})

	}

}

func BenchmarkCheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}

}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output : A
}
