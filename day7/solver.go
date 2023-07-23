package day7

import (
	"log"
)

const defaultCandidateSizeLimit = 100000
const defaultFileSystemSize = 70000000
const defaultUpdateSize = 30000000

const dirType = "dir"
const fileType = "file"
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
	//root := getSampleFileSystem()
	root := readFileSystemFromFile()

	updateDirectoriesSize(root)
	analyzeDirectoryCandidates(root, defaultCandidateSizeLimit)
	analyzeDirectoryCandidatesForFileSystemUpdate(root, defaultFileSystemSize, defaultUpdateSize)
}

/*
Given *FileSystemObject root with
*/
func analyzeDirectoryCandidates(root *FileSystemObject, candidateSizeLimit int) {
	totalSize := sumDirectoryCandidatesSize(root, candidateSizeLimit)
	log.Printf("Directories with size <= %d had a total size of %d\n", candidateSizeLimit, totalSize)
}

/*

*/
func analyzeDirectoryCandidatesForFileSystemUpdate(root *FileSystemObject, fileSystemSize int, updateSize int) {
	currentFreeSpace := fileSystemSize - root.size
	additionalSpaceRequired := updateSize - currentFreeSpace
	log.Println("In order to update file system, we need additional space:", additionalSpaceRequired)

	directory := findSmallestDirectoryCandidate(root, additionalSpaceRequired)
	log.Printf("The best directory to delete: %s, which results in saving space: %d\n", directory.name, directory.size)
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
	if node.objectType == fileType {
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
	if node.objectType == fileType {
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
