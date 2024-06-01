package csvLover

import "testing"

type testCase struct{}

func Test_CSVLoverBasic(t *testing.T) {
	opt := &Options{
		Path:       "winequality-red.csv",
		Comma:      ',',
		LazyQuotes: true,
		Dest:       &testCase{},
	}
	limit := &Limit{
		Row: &Row{
			From: 0,
			To:   21,
		},
		Col: &Col{
			From: 0,
			To:   12,
		},
	}
	csvLover, err := NewCSVLover(opt)
	if err != nil {
		t.Fatal(err)
	}
	records, err := csvLover.load(limit)
	if err != nil {
		t.Fatal(err)
	}
	for _, record := range records {
		t.Log(record)
	}
}
