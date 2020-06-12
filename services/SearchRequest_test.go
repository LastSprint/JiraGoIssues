package services

import "testing"

func TestMakeJiraRequestRightForFullModel(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "assignee = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithoutOrdering(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Priorities:       []string{"7"},
	}

	awaiting := "assignee = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\")"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithoutIS(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "assignee = 1 " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithoutES(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		IncludedStatuses: []string{"2"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "assignee = 1 and status in (\"2\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithoutIT(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "assignee = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithoutET(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "assignee = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithouPriorities(t *testing.T) {

	// Arrange

	model := SearchRequest{
		Assignee:         "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
	}

	awaiting := "assignee = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithoutAssignee(t *testing.T) {

	// Arrange

	model := SearchRequest{
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithProject(t *testing.T) {

	// Arrange

	model := SearchRequest{
		ProjectID:        "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "project = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithBoard(t *testing.T) {

	// Arrange

	model := SearchRequest{
		ProjectID:        "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		Board:            "IOS",
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "board = IOS and project = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}

func TestMakeJiraRequestRightWithReasons(t *testing.T) {

	// Arrange

	model := SearchRequest{
		ProjectID:        "1",
		IncludedStatuses: []string{"2"},
		ExcludedStatuses: []string{"3"},
		IncludedTypes:    []string{"4"},
		ExcludedTypes:    []string{"5"},
		IncludedReasons:  []string{"server"},
		Ordering: &OrderingModel{
			OrderTarget: "6",
			Type:        Ascending,
		},
		Priorities: []string{"7"},
	}

	awaiting := "project = 1 and status in (\"2\") " +
		"and status not in (\"3\") " +
		"and Reason in (\"server\") " +
		"and issuetype in (\"4\") " +
		"and issuetype not in (\"5\") " +
		"and priority in (\"7\") " +
		"ORDER BY 6 ASC"

	// Act

	result := model.MakeJiraRequest()

	// Assert

	if result != awaiting {
		t.Fatal(result)
	}
}