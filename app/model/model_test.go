package model

import (
	"fmt"
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

	if !reflect.DeepEqual(toUpdatePoll, shouldEqual) {
		t.Error("Updated poll has invalid modified choices")
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

	submission := SubmissionOptions{
		ChoiceIDs: []uint{4},
	}

	poll.AddSubmission(&submission)

	fmt.Printf("%v", poll.Content.Choices["4"])
	fmt.Printf("%v", pollAfter.Content.Choices["4"])

	if !reflect.DeepEqual(poll, pollAfter) {
		t.Error("Poll does not have updated choice counts")
	}
}
