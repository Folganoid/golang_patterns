package pkg

import "fmt"

const (
	ServerType = "server"
	PCType = "personal"
	NotebookType = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New (typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s Unknown type\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PCType:
		return NewPC()
	case NotebookType:
		return NewNoteBook()
	}
}
