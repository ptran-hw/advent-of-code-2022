package day4

import "fmt"

type Solver struct {
}

type conditionFunc func(assignmentA, assignmentB []int) bool

func (s Solver) Solve() {
	countFullOverlappingAssignments()
	countOverlappingAssignments()
}

func countAssignments(conditionFunc conditionFunc) {
	//assignments := getSampleAssignments()
	assignments := readAssignmentsFromFile()

	count := 0

	for index := 0; index < len(assignments)-1; index += 2 {
		currAssignment := assignments[index]
		nextAssignment := assignments[index+1]

		if conditionFunc(currAssignment, nextAssignment) {
			count++
		}
	}

	fmt.Printf("full overlapping section assignments: %d\n", count)
}

func countFullOverlappingAssignments() {
	countAssignments(containsFullOverlap)
}

func containsFullOverlap(assignmentA, assignmentB []int) bool {
	return (assignmentA[0] <= assignmentB[0] && assignmentA[1] >= assignmentB[1]) ||
		(assignmentB[0] <= assignmentA[0] && assignmentB[1] >= assignmentA[1])
}

func countOverlappingAssignments() {
	countAssignments(containsOverlap)
}

func containsOverlap(assignmentA, assignmentsB []int) bool {
	return (assignmentA[0] <= assignmentsB[0] && assignmentA[1] >= assignmentsB[0]) ||
		(assignmentsB[0] <= assignmentA[0] && assignmentsB[1] >= assignmentA[0])
}
