package csvdemo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)


func ReadCSVFileAll(path string) (result [][]string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %v", err)
	}

	r := csv.NewReader(f)
	result, err = r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("r.ReadAll failed with: %v", err)
	}

	return result, nil
}

func ReadCSVFileByLine(path string, nRecords int) (result [][]string, err error) {
	var record []string
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Open file failed: %v", err)
	}
	r := csv.NewReader(f)
	for nRecords > 0 {
		record, err = r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		nRecords--
		result = append(result, record)
	}

	return result, err
}

func WriteCsv(f io.Writer, record []string) error {
	w := csv.NewWriter(f)
	// records := [][]string{
	// 	{"date", "price", "name"},
	// 	{"2013-02-08", "15,07", "GOOG"},
	// 	{"2013-02-09", "15,09", "GOOG"},
	// }
	err := w.Write(record)
	if err != nil {
		return err
	}

	// csv.Writer might buffer writes for performance so we must
	// Flush to ensure all data has been written to underlying
	// writer
	w.Flush()

	// Flush doesn't return an error. If it failed to write, we
	// can get the error with Error()
	err = w.Error()
	if err != nil {
		return err
	}

	return err
}


func Run() {
	res, err := ReadCSVFileAll(filepath.Join("csvdemo", "csvdemo.csv"))
	if err != nil {
		panic(err)
	}
	for _, r := range res {
		fmt.Println(r)
	}
}