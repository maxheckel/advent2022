package main

import "testing"

func TestFindRepeatingCharsOfLength(t *testing.T) {
	type TestCase struct {
		input         string
		length        int
		expectedIndex int
	}
	testCases := []TestCase{
		{
			input:         "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			length:        4,
			expectedIndex: 7,
		},
		{
			input:         "bvwbjplbgvbhsrlpgdmjqwftvncz",
			length:        4,
			expectedIndex: 5,
		},
		{
			input:         "nppdvjthqldpwncqszvftbrmjlhg",
			length:        4,
			expectedIndex: 6,
		},
		{
			input:         "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			length:        4,
			expectedIndex: 10,
		},
		{
			input:         "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			length:        4,
			expectedIndex: 11,
		},
		{
			input:         "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			length:        14,
			expectedIndex: 19,
		},
		{
			input:         "bvwbjplbgvbhsrlpgdmjqwftvncz",
			length:        14,
			expectedIndex: 23,
		},
		{
			input:         "nppdvjthqldpwncqszvftbrmjlhg",
			length:        14,
			expectedIndex: 23,
		},
		{
			input:         "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			length:        14,
			expectedIndex: 29,
		},
		{
			input:         "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			length:        14,
			expectedIndex: 26,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			res := FindRepeatingCharsOfLength(testCase.input, testCase.length)
			if res != testCase.expectedIndex {
				t.Errorf("expected %d got %d", testCase.expectedIndex, res)
			}
		})
	}
}
