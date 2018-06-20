package model

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	originalPoll := Poll{
		Content: PollContent{
			Choices: map[string]*PollChoice{
				"4": &PollChoice{
					ID:    4,
					Count: 100,
				},
				"25": &PollChoice{
					ID:    25,
					Count: 324,
				},
				"1": &PollChoice{
					ID:    1,
					Count: 123,
				},
				"5": &PollChoice{
					ID:    5,
					Count: 9,
				},
			},
		},
	}

	modifiedPoll := Poll{
		Content: PollContent{
			Choices: map[string]*PollChoice{
				"4": &PollChoice{
					ID:    4,
					Count: 10000,
				},
				"25": &PollChoice{
					ID:    25,
					Count: 123525,
				},
				"1": &PollChoice{
					ID:    1,
					Count: 41234235,
				},
				"5": &PollChoice{
					ID:    5,
					Count: 12987249324,
				},
				"6": &PollChoice{
					ID:    6,
					Count: 123,
				},
			},
		},
	}

	shouldEqual := Poll{
		Content: PollContent{
			Choices: map[string]*PollChoice{
				"4": &PollChoice{
					ID:    4,
					Count: 100,
				},
				"25": &PollChoice{
					ID:    25,
					Count: 324,
				},
				"1": &PollChoice{
					ID:    1,
					Count: 123,
				},
				"5": &PollChoice{
					ID:    5,
					Count: 9,
				},
				"6": &PollChoice{
					ID:    6,
					Count: 0,
				},
			},
		},
	}

	toUpdatePoll := originalPoll
	// deep copy array of choices
	toUpdatePoll.Content.Choices = map[string]*PollChoice{}
	for key, element := range originalPoll.Content.Choices {
		toUpdatePoll.Content.Choices[key] = element
	}

	toUpdatePoll.Update(&modifiedPoll)

	// make updated timestamp the same
	shouldEqual.UpdatedAt = toUpdatePoll.UpdatedAt

	if !reflect.DeepEqual(toUpdatePoll, shouldEqual) {
		t.Error("updated poll has invalid modified choices")
	}
}

func TestAddSubmission(t *testing.T) {
	poll := Poll{
		Content: PollContent{
			Choices: map[string]*PollChoice{
				"4": &PollChoice{
					ID:    4,
					Count: 100,
				},
				"25": &PollChoice{
					ID:    25,
					Count: 324,
				},
				"1": &PollChoice{
					ID:    1,
					Count: 123,
				},
				"5": &PollChoice{
					ID:    5,
					Count: 9,
				},
			},
		},
	}

	pollAfter := Poll{
		Content: PollContent{
			Choices: map[string]*PollChoice{
				"4": &PollChoice{
					ID:    4,
					Count: 101,
				},
				"25": &PollChoice{
					ID:    25,
					Count: 324,
				},
				"1": &PollChoice{
					ID:    1,
					Count: 123,
				},
				"5": &PollChoice{
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
