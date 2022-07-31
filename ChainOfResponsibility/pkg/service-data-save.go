package pkg

import "fmt"

type DataService struct {
	Next Service
}

func (upd *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Printf("Data not update")
		return
	}
	fmt.Printf("Data save")
}

func (upd *DataService) SetNext(svc Service) {
	upd.Next = svc
}