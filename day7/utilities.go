package day7

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const fileSystemInputFile = "./day7/fileSystemInput.txt"


// input file has the following formats:
// $ cd /
// $ cd ..
// $ cd <dir-name>
// $ ls
// dir <dir-name>
// <file-size> <file-name>
func readFileSystemFromFile() *FileSystemObject {
	file, err := os.Open(fileSystemInputFile)
	if err != nil {
		log.Panicf("unable to read file system input: %v", err)
	}
	defer file.Close()

	listDirectoryCommand := "$ ls"
	changeDirPattern := regexp.MustCompile("\\$ cd (.+)")
	fileSystemFilePattern := regexp.MustCompile("([0-9]+) (.+)")
	fileSystemDirPattern := regexp.MustCompile("dir (.+)")
	root := &FileSystemObject{
		name:       "/",
		objectType: "dir",
		children:   make([]*FileSystemObject, 0),
	}
	head := root

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == listDirectoryCommand:
			continue
		case changeDirPattern.MatchString(line):
			path := changeDirPattern.FindStringSubmatch(line)[1]
			switch path {
			case "/":
				head = root
			case "..":
				head = head.parent
			default:
				for index, child := range head.children {
					if child.name == path {
						head = child
						break
					}

					if index == len(head.children) - 1 {
						log.Panicf("unable to change directory with head: %s, target: %s", head.name, path)
					}
				}
			}
		case fileSystemDirPattern.MatchString(line):
			dirName := fileSystemDirPattern.FindStringSubmatch(line)[1]
			child := &FileSystemObject{
				name:       dirName,
				objectType: dirType,
				children:   make([]*FileSystemObject, 0),
				parent:     head,
			}
			head.children = append(head.children, child) // append creates a new slice, so we need to reassign here
		case fileSystemFilePattern.MatchString(line):
			matches := fileSystemFilePattern.FindStringSubmatch(line)
			fileSize, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Panicf("unable to read filesize value: %v", err)
			}

			fileName := matches[2]
			child := &FileSystemObject{
				name:       fileName,
				objectType: fileType,
				size:       fileSize,
				parent:     head,
			}
			head.children = append(head.children, child)
		default:
			log.Panicf("input line was not recognized: %s", line)
		}
	}

	return root
}

// add field names explicitly to avoid breaking with struct changes
// parent field not set for simplicity
func getSampleFileSystem() *FileSystemObject {
	return &FileSystemObject{
		name:       "/",
		objectType: "dir",
		children: []*FileSystemObject{
			{
				name:       "a",
				objectType: "dir",
				children: []*FileSystemObject{
					{
						name:       "e",
						objectType: "dir",
						children: []*FileSystemObject{
							{
								name:       "i",
								objectType: "file",
								size:       584,
							},
						},
					},
					{
						name:       "f",
						objectType: "file",
						size:       29116,
					},
					{
						name:       "g",
						objectType: "file",
						size:       2557,
					},
					{
						name:       "h.lst",
						objectType: "file",
						size:       62596,
					},
				},
			},
			{
				name:       "b.txt",
				objectType: "file",
				size:       14848514,
			},
			{
				name:       "c.dat",
				objectType: "file",
				size:       8504156,
			},
			{
				name:       "d",
				objectType: "dir",
				children: []*FileSystemObject{
					{
						name:       "j",
						objectType: "file",
						size:       4060174,
					},
					{
						name:       "d.log",
						objectType: "file",
						size:       8033020,
					},
					{
						name:       "d.ext",
						objectType: "file",
						size:       5626152,
					},
					{
						name:       "k",
						objectType: "file",
						size:       7214296,
					},
				},
			},
		},
	}
}
