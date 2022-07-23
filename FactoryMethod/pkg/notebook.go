package pkg

import "fmt"

type NoteBook struct {
	Type string
	Core int
	Memory int
	Monitor bool
}

func NewNoteBook() Computer {
	return NoteBook{
		Type: NotebookType,
		Core: 2,
		Memory: 8,
		Monitor: true,
	}
}

func (pc NoteBook) GetType() string {
	return pc.Type
}

func (pc NoteBook) PrintDetails() {
	fmt.Printf("%s Core:[%d], Mem: [%d]\n", pc.Type, pc.Core, pc.Memory)
}
