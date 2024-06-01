package csvLover

import (
	"encoding/csv"
	"os"
)

type DateHandler func(data any) error

type CSVLover struct {
	csvReader   *csv.Reader
	destination any
}

func NewCSVLover(options *Options) (*CSVLover, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}
	f, err := os.Open(options.Path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	if options.Comma == 0 {
		r.Comma = ';'
	} else {
		r.Comma = options.Comma
	}
	r.LazyQuotes = options.LazyQuotes

	lover := new(CSVLover)
	lover.csvReader = r
	return lover, nil
}

func (lover *CSVLover) Persistence(limit *Limit, f DateHandler) error {
	_, err := lover.load(limit)
	if err != nil {
		return err
	}
	return nil
}

func (lover *CSVLover) load(limit *Limit) (records [][]string, err error) {
	csvReader := lover.csvReader
	row := limit.Row
	col := limit.Col
	for i := 0; i < row.To; i++ {
		if i < row.From {
			_, _ = csvReader.Read()
		} else {
			record, err := csvReader.Read()
			if err != nil {
				return nil, err
			}
			record = record[col.From:col.To]
			records = append(records, record)
		}
	}
	return records, nil
}
