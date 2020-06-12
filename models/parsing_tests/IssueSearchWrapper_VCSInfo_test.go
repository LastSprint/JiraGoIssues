package parsing_tests

import (
	"encoding/json"
	"testing"

	"github.com/LastSprint/pkg_go_jira_service/models"
)

func Test_VCSInfo_ParseSuccessfully(t *testing.T) {
	// Arrange

	jsonStr := `
	{
    "expand": "names,schema",
    "startAt": 0,
    "maxResults": 50,
    "total": 1,
    "issues": [
        {
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "78607",
            "self": "https://jira.surfstudio.ru/rest/client/2/issue/78607",
            "key": "SBB-4852",
            "fields": {
                "customfield_10100": "{summaryBean=com.atlassian.jira.plugin.devstatus.rest.SummaryBean@440ebfb2[summary={pullrequest=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@1100d390[overall=PullRequestOverallBean{stateCount=1, state='MERGED', details=PullRequestOverallDetails{openCount=0, mergedCount=1, declinedCount=0}},byInstanceType={source_repo_services=com.atlassian.jira.plugin.devstatus.summary.beans.ObjectByInstanceTypeBean@36fdd86c[count=1,name=Bitbucket Cloud]}], build=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@43cb0da8[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BuildOverallBean@53deb410[failedBuildCount=0,successfulBuildCount=0,unknownBuildCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], review=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@26ef824d[overall=com.atlassian.jira.plugin.devstatus.summary.beans.ReviewsOverallBean@743a4133[stateCount=0,state=<null>,dueDate=<null>,overDue=false,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], deployment-environment=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@6f07b451[overall=com.atlassian.jira.plugin.devstatus.summary.beans.DeploymentOverallBean@49dff395[topEnvironments=[],showProjects=false,successfulCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], repository=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@13f8fd09[overall=com.atlassian.jira.plugin.devstatus.summary.beans.CommitOverallBean@35ed159c[count=2,lastUpdated=2020-01-12T17:43:34.000+0300,lastUpdatedTimestamp=2020-01-12T17:43:34.000+03:00],byInstanceType={source_repo_services=com.atlassian.jira.plugin.devstatus.summary.beans.ObjectByInstanceTypeBean@e548963[count=2,name=Bitbucket Cloud]}], branch=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@516f9e61[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BranchOverallBean@5ab31f0[count=1,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={source_repo_services=com.atlassian.jira.plugin.devstatus.summary.beans.ObjectByInstanceTypeBean@240ef6fd[count=1,name=Bitbucket Cloud]}]},errors=[],configErrors=[]], devSummaryJson={\"cachedValue\":{\"errors\":[],\"configErrors\":[],\"summary\":{\"pullrequest\":{\"overall\":{\"count\":1,\"lastUpdated\":\"2020-01-12T17:43:37.000+0300\",\"stateCount\":1,\"state\":\"MERGED\",\"details\":{\"openCount\":0,\"mergedCount\":1,\"declinedCount\":0,\"total\":1},\"open\":false},\"byInstanceType\":{\"source_repo_services\":{\"count\":1,\"name\":\"Bitbucket Cloud\"}}},\"build\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"failedBuildCount\":0,\"successfulBuildCount\":0,\"unknownBuildCount\":0},\"byInstanceType\":{}},\"review\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":null,\"dueDate\":null,\"overDue\":false,\"completed\":false},\"byInstanceType\":{}},\"deployment-environment\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"topEnvironments\":[],\"showProjects\":false,\"successfulCount\":0},\"byInstanceType\":{}},\"repository\":{\"overall\":{\"count\":2,\"lastUpdated\":\"2020-01-12T17:43:34.000+0300\"},\"byInstanceType\":{\"source_repo_services\":{\"count\":2,\"name\":\"Bitbucket Cloud\"}}},\"branch\":{\"overall\":{\"count\":1,\"lastUpdated\":null},\"byInstanceType\":{\"source_repo_services\":{\"count\":1,\"name\":\"Bitbucket Cloud\"}}}}},\"isStale\":false}}"
            }
        }
    ]
}
`

	// Act

	var data models.IssueSearchWrapperEntity

	json.Unmarshal([]byte(jsonStr), &data)

	// Assert

	if data.Issues[0].Fields.VCSInfo == nil {
		t.Fatal(data)
	}
}
