package proxy

type DataBase struct {
}

func (db DataBase) GetData(string) ([]string, error) {
	return []string{
		"String1",
		"Last string",
	}, nil
}
