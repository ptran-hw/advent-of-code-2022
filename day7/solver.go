package day7

import "fmt"

const directoryCandidateSizeLimit = 100000
const fileObjectType = "file"
const totalFileSystemSize = 70000000
const updateRequiredSpace = 30000000

type FileSystemObject struct {
	name       string
	objectType string
	size       int
	children   []*FileSystemObject
	parent     *FileSystemObject
}

type Solver struct {
}

func (s Solver) Solve() {
	analyzeDirectoryCandidates()
	analyzeDirectoryCandidatesForFileSystemUpdate()
}

func analyzeDirectoryCandidates() {
	//root := getSampleFileSystem()
	root := readFileSystemFromFile()

	updateDirectoriesSize(root)

	fmt.Printf("Directories with size <= %d had a total size of %d\n", directoryCandidateSizeLimit, sumDirectoryCandidatesSize(root, directoryCandidateSizeLimit))
}

func analyzeDirectoryCandidatesForFileSystemUpdate() {
	//root := getSampleFileSystem()
	root := readFileSystemFromFile()

	updateDirectoriesSize(root)

	currentFreeSpace := totalFileSystemSize - root.size
	additionalSpaceRequired := updateRequiredSpace - currentFreeSpace
	fmt.Printf("In order to update file system, we need additional space: %d\n", additionalSpaceRequired)

	directory := findSmallestDirectoryCandidate(root, additionalSpaceRequired)
	fmt.Printf("The best directory to delete for file system update: %v\n", directory) // print statements does display object values
}

func updateDirectoriesSize(node *FileSystemObject) {
	if node.children == nil {
		return
	}

	for _, child := range node.children {
		updateDirectoriesSize(child)
	}

	total := 0
	for _, child := range node.children {
		total += child.size
	}

	node.size = total
}

func sumDirectoryCandidatesSize(node *FileSystemObject, candidateSizeLimit int) int {
	if node.objectType == fileObjectType {
		return 0
	}

	total := 0
	if node.size <= candidateSizeLimit {
		total = node.size
	}

	for _, child := range node.children {
		total += sumDirectoryCandidatesSize(child, candidateSizeLimit)
	}

	return total
}

func findSmallestDirectoryCandidate(node *FileSystemObject, minSize int) *FileSystemObject {
	if node.objectType == fileObjectType {
		return nil
	}

	if node.size < minSize {
		return nil
	}

	smallestDir := node
	for _, child := range node.children {
		smallestSubDir := findSmallestDirectoryCandidate(child, minSize)

		if smallestSubDir != nil && smallestDir.size >= smallestSubDir.size {
			smallestDir = smallestSubDir
		}
	}

	return smallestDir
}
