package model

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	content := PollContent{
		Choices: []PollChoice{
			PollChoice{
				ID:    4,
				Count: 100,
			},
			PollChoice{
				ID:    25,
				Count: 324,
			},
			PollChoice{
				ID:    1,
				Count: 123,
			},
			PollChoice{
				ID:    5,
				Count: 9,
			},
		},
	}

	content.Sort()

	sorted := PollContent{
		Choices: []PollChoice{
			PollChoice{
				ID:    1,
				Count: 123,
			},
			PollChoice{
				ID:    4,
				Count: 100,
			},
			PollChoice{
				ID:    5,
				Count: 9,
			},
			PollChoice{
				ID:    25,
				Count: 324,
			},
		},
	}

	if !reflect.DeepEqual(content, sorted) {
		t.Error("Poll content is not sorted")
	}
}

func TestUpdate(t *testing.T) {
	originalPoll := Poll{
		Content: PollContent{
			Choices: []PollChoice{
				PollChoice{
					ID:    4,
					Count: 100,
				},
				PollChoice{
					ID:    25,
					Count: 324,
				},
				PollChoice{
					ID:    1,
					Count: 123,
				},
				PollChoice{
					ID:    5,
					Count: 9,
				},
			},
		},
	}

	modifiedPoll := Poll{
		Content: PollContent{
			Choices: []PollChoice{
				PollChoice{
					ID:    4,
					Count: 10000,
				},
				PollChoice{
					ID:    25,
					Count: 123525,
				},
				PollChoice{
					ID:    1,
					Count: 41234235,
				},
				PollChoice{
					ID:    5,
					Count: 12987249324,
				},
				PollChoice{
					ID:    6,
					Count: 123,
				},
			},
		},
	}

	shouldEqual := Poll{
		Content: PollContent{
			Choices: []PollChoice{
				PollChoice{
					ID:    4,
					Count: 100,
				},
				PollChoice{
					ID:    25,
					Count: 324,
				},
				PollChoice{
					ID:    1,
					Count: 123,
				},
				PollChoice{
					ID:    5,
					Count: 9,
				},
				PollChoice{
					ID:    6,
					Count: 0,
				},
			},
		},
	}

	toUpdatePoll := originalPoll
	// deep copy array of choices
	toUpdatePoll.Content.Choices = make([]PollChoice, len(toUpdatePoll.Content.Choices))
	copy(toUpdatePoll.Content.Choices, originalPoll.Content.Choices)

	toUpdatePoll.Update(&modifiedPoll)

	if !reflect.DeepEqual(toUpdatePoll, shouldEqual) {
		t.Error("Updated poll has invalid modified choices")
	}
}
