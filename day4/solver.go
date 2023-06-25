package day4

import "fmt"

type Solver struct {
}

type Assignment struct {
	start int
	end   int
}

type conditionFunc func(assignmentA, assignmentB Assignment) bool

func (s Solver) Solve() {
	//assignmentPairs := getSampleAssignments()
	assignmentPairs := readAssignmentsFromFile()

	countFullOverlappingAssignments(assignmentPairs)
	countOverlappingAssignments(assignmentPairs)
}

func getFilteredAssignmentPairsCount(assignmentPairs [][]Assignment, conditionFunc conditionFunc) {
	count := 0

	for index := 0; index < len(assignmentPairs); index++ {
		pair := assignmentPairs[index]
		assignmentA := pair[0]
		assignmentB := pair[1]

		if conditionFunc(assignmentA, assignmentB) {
			count++
		}
	}

	fmt.Printf("full overlapping section assignment pairs: %d\n", count)
}

/*
Given [][]Assignment assignmentPairs,
where assignmentPairs[i] is a pair of assignments and assignments contains int start and end
Count the number of assignment pairs where one range is fully contained in the other
*/
func countFullOverlappingAssignments(assignmentPairs [][]Assignment) {
	getFilteredAssignmentPairsCount(assignmentPairs, containsFullOverlap)
}

func containsFullOverlap(assignmentA, assignmentB Assignment) bool {
	return (assignmentA.start <= assignmentB.start && assignmentA.end >= assignmentB.end) ||
		(assignmentB.start <= assignmentA.start && assignmentB.end >= assignmentA.end)
}

/*
Given [][]Assignment assignmentPairs,
where assignmentPairs[i] is a pair of assignments and assignments contains int start and end
Count the number of assignment pairs that one range is overlapping with the other
*/
func countOverlappingAssignments(assignmentPairs [][]Assignment) {
	getFilteredAssignmentPairsCount(assignmentPairs, containsOverlap)
}

func containsOverlap(assignmentA, assignmentsB Assignment) bool {
	return (assignmentA.start <= assignmentsB.start && assignmentA.end >= assignmentsB.start) ||
		(assignmentsB.start <= assignmentA.start && assignmentsB.end >= assignmentA.start)
}
