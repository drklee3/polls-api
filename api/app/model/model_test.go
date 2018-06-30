package model

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	originalPoll := Poll{
		Content: &PollContent{
			Choices: map[string]*PollChoice{
				"4": {
					ID:    4,
					Count: 100,
				},
				"25": {
					ID:    25,
					Count: 324,
				},
				"1": {
					ID:    1,
					Count: 123,
				},
				"5": {
					ID:    5,
					Count: 9,
				},
			},
		},
	}

	modifiedPoll := Poll{
		Content: &PollContent{
			Choices: map[string]*PollChoice{
				"4": {
					ID:    4,
					Count: 10000,
				},
				"25": {
					ID:    25,
					Count: 123525,
				},
				"1": {
					ID:    1,
					Count: 41234235,
				},
				"5": {
					ID:    5,
					Count: 12987249324,
				},
				"6": {
					ID:    6,
					Count: 123,
				},
			},
		},
	}

	shouldEqual := Poll{
		Content: &PollContent{
			Choices: map[string]*PollChoice{
				"4": {
					ID:    4,
					Count: 100,
				},
				"25": {
					ID:    25,
					Count: 324,
				},
				"1": {
					ID:    1,
					Count: 123,
				},
				"5": {
					ID:    5,
					Count: 9,
				},
				"6": {
					ID:    6,
					Count: 0,
				},
			},
		},
	}

	originalPoll.Update(&modifiedPoll)

	// make updated timestamp the same
	shouldEqual.UpdatedAt = originalPoll.UpdatedAt

	if !reflect.DeepEqual(originalPoll, shouldEqual) {
		t.Error("updated poll has invalid modified choices")
	}
}

func TestAddSubmission(t *testing.T) {
	poll := Poll{
		Content: &PollContent{
			Choices: map[string]*PollChoice{
				"4": {
					ID:    4,
					Count: 100,
				},
				"25": {
					ID:    25,
					Count: 324,
				},
				"1": {
					ID:    1,
					Count: 123,
				},
				"5": {
					ID:    5,
					Count: 9,
				},
			},
		},
	}

	pollAfter := Poll{
		Content: &PollContent{
			Choices: map[string]*PollChoice{
				"4": {
					ID:    4,
					Count: 101,
				},
				"25": {
					ID:    25,
					Count: 324,
				},
				"1": {
					ID:    1,
					Count: 123,
				},
				"5": {
					ID:    5,
					Count: 9,
				},
			},
		},
	}

	submission := Submission{
		ChoiceIDs: []uint{4},
	}

	poll.AddSubmission(&submission)

	if !reflect.DeepEqual(poll, pollAfter) {
		t.Error("poll does not have updated choice counts")
	}
}

func TestInitializePoll(t *testing.T) {
	poll := Poll{
		Content: &PollContent{
			Choices: map[string]*PollChoice{
				"0": {
					ID:   0,
					Name: "first",
				},
				"5": {
					ID:   5,
					Name: "this should fail",
				},
			},
		},
	}

	if err := poll.Initialize(); err == nil {
		t.Error("poll should error with invalid choice ID")
	}
}

func TestAddSubmissionEmpty(t *testing.T) {
	poll := Poll{}
	submission := Submission{}

	err := poll.AddSubmission(&submission)

	if err == nil && err.Error() == "Submission cannot be empty" {
		t.Error("poll allowed empty submission")
	}
}

func TestAddSubmissionArchived(t *testing.T) {
	poll := Poll{
		Archived: true,
	}
	submission := Submission{
		ChoiceIDs: []uint{1},
	}

	err := poll.AddSubmission(&submission)

	if err == nil && err.Error() == "Poll is archived" {
		t.Error("archived poll allowed submission")
	}
}
