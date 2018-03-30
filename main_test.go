package main

import "testing"

func Test_firstDayIndex(t *testing.T) {
	testCases := []struct {
		Tag                            string
		Date, DateIndex, FirstDayIndex int
	}{
		{
			Tag:           "benglai saal 1425, 1 boisakh expected first day zero based index is 6",
			Date:          1,
			DateIndex:     6,
			FirstDayIndex: 6,
		},
		{
			Tag:           "benglai saal 1425, 16 joishta expected first day zero based index is 2",
			Date:          16,
			DateIndex:     3,
			FirstDayIndex: 2,
		},
		{
			Tag:           "benglai saal 1425, 13 kartik expected first day zero based index is 0",
			Date:          13,
			DateIndex:     0,
			FirstDayIndex: 2,
		},
		{
			Tag:           "benglai saal 1425, 25 poush expected first day zero based index is 2",
			Date:          25,
			DateIndex:     2,
			FirstDayIndex: 6,
		},
		{
			Tag:           "english year 2018, 28 march expected first day zero based index is 3",
			Date:          28,
			DateIndex:     3,
			FirstDayIndex: 4,
		},
	}

	for _, tc := range testCases {
		if r := firstDayIndex(tc.Date, tc.DateIndex); r != tc.FirstDayIndex {
			t.Errorf("firstday index failed\nExpected: %d Got: %d", tc.FirstDayIndex, r)
		}
	}
}

func Test_enToBnNumber(t *testing.T) {
	testCases := map[int]string{
		1:    "১",
		2:    "২",
		3:    "৩",
		4:    "৪",
		5:    "৫",
		6:    "৬",
		7:    "৭",
		8:    "৮",
		9:    "৯",
		0:    "০",
		2018: "২০১৮",
	}
	for en, bn := range testCases {
		if r := enToBnNumber(en); r != bn {
			t.Error("e2dDigit failed", "expected: ", bn, " got: ", r)
		}
	}
}

func Test_highlight(t *testing.T) {
	testCases := map[string]string{
		"1":   "\033[1;32m1\033[0m",
		"2":   "\033[1;32m2\033[0m",
		"foo": "\033[1;32mfoo\033[0m",
	}

	for in, out := range testCases {
		if r := highlight(in); r != out {
			t.Error("highlight failed", "expected: ", out, " got: ", r)
		}
	}
}

func Test_bold(t *testing.T) {
	testCases := map[string]string{
		"1":   "\033[1m1\033[0m",
		"2":   "\033[1m2\033[0m",
		"foo": "\033[1mfoo\033[0m",
	}

	for in, out := range testCases {
		if r := bold(in); r != out {
			t.Error("bold failed", "expected: ", out, " got: ", r)
		}
	}
}
