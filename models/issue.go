package models

type CreatedIssue struct {
	Id string `json:"id"`
	Key string `json:"key"`
	Self string `json:"self"`
	Transition NestedResponse `json:"transition"`
}

type NestedResponse struct {
	Status int32 `json:"status"`
	ErrorCollection ErrorCollection `json:"errorCollection"`
}

type ErrorCollection struct {
	ErrorMessages []string `json:"errorMessages"`
	Errors Any `json:"errors"`
	Status int32 `json:"status"`
}

type CreatedIssues struct {
	Issues []CreatedIssue `json:"issues"`
	Errors []BulkOperationErrorResult `json:"errors"`
}

type BulkOperationErrorResult struct {
	Status int32 `json:"status"`
	ElementErrors ErrorCollection `json:"elementErrors"`
	FailedElementNumber int32 `json:"failedElementNumber"`
}

type IssueCreateMetadata struct {
	Expand string `json:"expand"`
	Projects []ProjectIssueCreateMetadata `json:"projects"`
}

type ProjectIssueCreateMetadata struct {
	Expand string `json:"expand"`
	Self string `json:"self"`
	Id string `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	AvatarUrls AvatarUrlsBean `json:"avatarUrls"`
	IssueTypes []IssueTypeIssueCreateMetadata `json:"issueTypes"`
}

type IssueTypeIssueCreateMetadata struct {
	Self string `json:"self"`
	Id string `json:"id"`
	Description string `json:"description"`
	IconUrl string `json:"iconUrl"`
	Name string `json:"name"`
	Subtask bool `json:"subtask"`
	AvatarId int64 `json:"avatarId"`
	Scope Scope `json:"scope"`
	Expand string `json:"expand"`
	Fields Any `json:"fields"`
}

type IssueBean struct {
	Expand string `json:"expand"`
	Id string `json:"id"`
	Self string `json:"self"`
	Key string `json:"key"`
	RenderedFields Any `json:"renderedFields"`
	Properties Any `json:"properties"`
	Names Any `json:"names"`
	Schema Any `json:"schema"`
	Transitions []Transition `json:"transitions"`
	Operations Operations `json:"operations"`
	EditMeta IssueUpdateMetadata `json:"editmeta"`
	ChangeLog PageOfChangelogs `json:"changelog"`
	VersionedRepresentations Any `json:"versionedRepresentations"`
	FieldsToInclude IncludedFields `json:"fieldsToInclude"`
	Fields Any `json:"fields"`
}

type Transition struct {
	Id string `json:"id"`
	Name string `json:"name"`
	To StatusDetails `json:"to"`
	HasScreen bool `json:"hasScreen"`
	IsGlobal bool `json:"isGlobal"`
	IsInitial bool `json:"isInitial"`
	IsConditional bool `json:"isConditional"`
	Fields Any `json:"fields"`
	Expand string `json:"expand"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type StatusDetails struct {
	Self string `json:"self"`
	Description string `json:"description"`
	IconUrl string `json:"iconUrl"`
	Name string `json:"name"`
	Id string `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type StatusCategory struct {
	Self string `json:"self"`
	Id int64 `json:"id"`
	Key string `json:"key"`
	ColorName string `json:"colorName"`
	Name string `json:"name"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type Operations struct {
	LinkGroup []LinkGroup `json:"linkGroup"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type LinkGroup struct {
	Id string `json:"id"`
	StyleClass string `json:"styleClass"`
	Header SimpleLink `json:"header"`
	Weight int32 `json:"weight"`
	Links []SimpleLink `json:"links"`
	Groups []LinkGroup `json:"groups"`
}

type SimpleLink struct {
	Id string `json:"id"`
	StyleClass string `json:"styleClass"`
	IconClass string `json:"iconClass"`
	Label string `json:"label"`
	Title string `json:"title"`
	Href string `json:"href"`
	Weight int32 `json:"weight"`
}

type IssueUpdateMetadata struct {
	Fields Any `json:"fields"`
}

type PageOfChangelogs struct {
	StartAt int32 `json:"startAt"`
	MaxResults int32 `json:"maxResults"`
	Total int32 `json:"total"`
	Histories []Changelog `json:"histories"`
}

type Changelog struct {
	Id string `json:"id"`
	Author UserDetails `json:"author"`
	Created string `json:"created"`
	Items []ChangeDetails `json:"items"`
	HistoryMetadata HistoryMetadata `json:"historyMetadata"`
}

type ChangeDetails struct {
	Field string `json:"field"`
	FieldType string `json:"fieldType"`
	FieldId string `json:"fieldId"`
	From string `json:"from"`
	FromString string `json:"fromString"`
	To string `json:"to"`
	ToString string `json:"toString"`
}

type HistoryMetadata struct {
	Type string `json:"type"`
	Description string `json:"description"`
	DescriptionKey string `json:"descriptionKey"`
	ActivityDescription string `json:"activityDescription"`
	ActivityDescriptionKey string `json:"activityDescriptionKey"`
	EmailDescription string `json:"emailDescription"`
	EmailDescriptionKey string `json:"emailDescriptionKey"`
	Actor HistoryMetadataParticipant `json:"actor"`
	Generator HistoryMetadataParticipant `json:"generator"`
	Cause HistoryMetadataParticipant `json:"cause"`
	ExtraData Any `json:"extraData"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type HistoryMetadataParticipant struct {
	Id string `json:"id"`
	DisplayName string `json:"displayName"`
	DisplayNameKey string `json:"displayNameKey"`
	Type string `json:"type"`
	AvatarUrl string `json:"avatarUrl"`
	Url string `json:"url"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type IncludedFields struct {
	Included []string `json:"included"`
	ActuallyIncluded []string `json:"actuallyIncluded"`
	Excluded []string `json:"excluded"`
}

type Transitions struct {
	Expand string `json:"expand"`
	Transitions []Transition `json:"transitions"`
}