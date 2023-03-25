package day4

import "fmt"

type Solver struct {
}

func (s Solver) Solve() {
	countFullOverlappingAssignments()
	countOverlappingAssignments()
}

func countFullOverlappingAssignments() {
	//assignments := getSampleAssignments()
	assignments := readAssignmentsFromFile()

	count := 0

	for index := 0; index < len(assignments)-1; index += 2 {
		currAssignment := assignments[index]
		nextAssignment := assignments[index+1]

		if containsFullOverlap(currAssignment, nextAssignment) {
			count++
		}
	}

	fmt.Printf("full overlapping section assignments: %d\n", count)
}

func containsFullOverlap(assignmentA, assignmentB []int) bool {
	return (assignmentA[0] <= assignmentB[0] && assignmentA[1] >= assignmentB[1]) ||
		(assignmentB[0] <= assignmentA[0] && assignmentB[1] >= assignmentA[1])
}

func countOverlappingAssignments() {
	//assignments := getSampleAssignments()
	assignments := readAssignmentsFromFile()

	count := 0

	for index := 0; index < len(assignments)-1; index += 2 {
		currAssignment := assignments[index]
		nextAssignment := assignments[index+1]

		if containsOverlap(currAssignment, nextAssignment) {
			count++
		}
	}

	fmt.Printf("overlapping section assignments: %d\n", count)
}

func containsOverlap(assignmentA, assignmentsB []int) bool {
	return (assignmentA[0] <= assignmentsB[0] && assignmentA[1] >= assignmentsB[0]) ||
		(assignmentsB[0] <= assignmentA[0] && assignmentsB[1] >= assignmentA[0])
}
