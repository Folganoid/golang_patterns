package pkg

type DataBase struct {

}

func (db DataBase) GetData(user string) ([]string, error) {
	return []string {
		"Row 1",
		"Last row",
	}, nil
}