package parsing_tests

import (
	"encoding/json"
	"testing"

	"github.com/LastSprint/pkg_go_jira_service/models"
)

func TestCommentsModelParseSuccessfully(t *testing.T) {
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
            "id": "79741",
            "self": "https://jira.surfstudio.ru/rest/client/2/issue/79741",
            "key": "SBB-4949",
            "fields": {
                "comment": {
                    "comments": [
                        {
                            "self": "https://jira.surfstudio.ru/rest/client/2/issue/79741/comment/28203",
                            "id": "28203",
                            "author": {
                                "self": "https://jira.surfstudio.ru/rest/client/2/user?username=polyanskih",
                                "name": "polyanskih",
                                "key": "polyanskih",
                                "emailAddress": "polyanskih@surfstudio.ru",
                                "avatarUrls": {
                                    "48x48": "https://jira.surfstudio.ru/secure/useravatar?ownerId=polyanskih&avatarId=12660",
                                    "24x24": "https://jira.surfstudio.ru/secure/useravatar?size=small&ownerId=polyanskih&avatarId=12660",
                                    "16x16": "https://jira.surfstudio.ru/secure/useravatar?size=xsmall&ownerId=polyanskih&avatarId=12660",
                                    "32x32": "https://jira.surfstudio.ru/secure/useravatar?size=medium&ownerId=polyanskih&avatarId=12660"
                                },
                                "displayName": "Артур Полянских",
                                "active": true,
                                "timeZone": "Europe/Moscow"
                            },
                            "body": "https://jira.sbibankllc.ru/browse/FML-3096",
                            "updateAuthor": {
                                "self": "https://jira.surfstudio.ru/rest/client/2/user?username=polyanskih",
                                "name": "polyanskih",
                                "key": "polyanskih",
                                "emailAddress": "polyanskih@surfstudio.ru",
                                "avatarUrls": {
                                    "48x48": "https://jira.surfstudio.ru/secure/useravatar?ownerId=polyanskih&avatarId=12660",
                                    "24x24": "https://jira.surfstudio.ru/secure/useravatar?size=small&ownerId=polyanskih&avatarId=12660",
                                    "16x16": "https://jira.surfstudio.ru/secure/useravatar?size=xsmall&ownerId=polyanskih&avatarId=12660",
                                    "32x32": "https://jira.surfstudio.ru/secure/useravatar?size=medium&ownerId=polyanskih&avatarId=12660"
                                },
                                "displayName": "Артур Полянских",
                                "active": true,
                                "timeZone": "Europe/Moscow"
                            },
                            "created": "2019-12-30T17:07:01.514+0300",
                            "updated": "2019-12-30T17:07:01.514+0300"
                        },
                        {
                            "self": "https://jira.surfstudio.ru/rest/client/2/issue/79741/comment/28208",
                            "id": "28208",
                            "author": {
                                "self": "https://jira.surfstudio.ru/rest/client/2/user?username=kuzyaev",
                                "name": "kuzyaev",
                                "key": "alex_kuzyaev",
                                "emailAddress": "kuzyaev@surfstudio.ru",
                                "avatarUrls": {
                                    "48x48": "https://jira.surfstudio.ru/secure/useravatar?ownerId=alex_kuzyaev&avatarId=12209",
                                    "24x24": "https://jira.surfstudio.ru/secure/useravatar?size=small&ownerId=alex_kuzyaev&avatarId=12209",
                                    "16x16": "https://jira.surfstudio.ru/secure/useravatar?size=xsmall&ownerId=alex_kuzyaev&avatarId=12209",
                                    "32x32": "https://jira.surfstudio.ru/secure/useravatar?size=medium&ownerId=alex_kuzyaev&avatarId=12209"
                                },
                                "displayName": "Александр Кузяев",
                                "active": true,
                                "timeZone": "Asia/Yekaterinburg"
                            },
                            "body": "Не воспроизводится",
                            "updateAuthor": {
                                "self": "https://jira.surfstudio.ru/rest/client/2/user?username=kuzyaev",
                                "name": "kuzyaev",
                                "key": "alex_kuzyaev",
                                "emailAddress": "kuzyaev@surfstudio.ru",
                                "avatarUrls": {
                                    "48x48": "https://jira.surfstudio.ru/secure/useravatar?ownerId=alex_kuzyaev&avatarId=12209",
                                    "24x24": "https://jira.surfstudio.ru/secure/useravatar?size=small&ownerId=alex_kuzyaev&avatarId=12209",
                                    "16x16": "https://jira.surfstudio.ru/secure/useravatar?size=xsmall&ownerId=alex_kuzyaev&avatarId=12209",
                                    "32x32": "https://jira.surfstudio.ru/secure/useravatar?size=medium&ownerId=alex_kuzyaev&avatarId=12209"
                                },
                                "displayName": "Александр Кузяев",
                                "active": true,
                                "timeZone": "Asia/Yekaterinburg"
                            },
                            "created": "2019-12-30T22:39:25.964+0300",
                            "updated": "2019-12-30T22:39:25.964+0300"
                        }
                    ],
                    "maxResults": 2,
                    "total": 2,
                    "startAt": 0
                }
            }
        }
    ]
}
`

	// Act

	var res models.IssueSearchWrapperEntity

	json.Unmarshal([]byte(jsonStr), &res)

	// Assert

	if len(res.Issues[0].Fields.Comments.Comments) == 0 {
		t.Fatal(res)
	}
}
