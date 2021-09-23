package zendesk

import (
	"time"
)

type GetTicketsResponse struct {
	Tickets      []Ticket `json:"tickets"`
	NextPage     *string  `json:"next_page"`
	PreviousPage *string  `json:"previous_page"`
	Count        int      `json:"count"`
}

type Ticket struct {
	URL        string      `json:"url"`
	ID         int         `json:"id"`
	ExternalID interface{} `json:"external_id"`
	Via        struct {
		Channel string `json:"channel"`
		Source  struct {
			Rel  string `json:"rel"`
			From struct {
				FormattedPhone string `json:"formatted_phone"`
				Phone          string `json:"phone"`
				Name           string `json:"name"`
			} `json:"from"`
			To struct {
				FormattedPhone string `json:"formatted_phone"`
				Phone          string `json:"phone"`
				Name           string `json:"name"`
				BrandID        int64  `json:"brand_id"`
			} `json:"to"`
		} `json:"source"`
	} `json:"via"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Type            interface{}   `json:"type"`
	Subject         string        `json:"subject"`
	RawSubject      string        `json:"raw_subject"`
	Description     string        `json:"description"`
	Priority        string        `json:"priority"`
	Status          string        `json:"status"`
	Recipient       string        `json:"recipient"`
	RequesterID     int64         `json:"requester_id"`
	SubmitterID     int64         `json:"submitter_id"`
	AssigneeID      int64         `json:"assignee_id"`
	OrganizationID  interface{}   `json:"organization_id"`
	GroupID         int64         `json:"group_id"`
	CollaboratorIds []interface{} `json:"collaborator_ids"`
	FollowerIds     []interface{} `json:"follower_ids"`
	EmailCcIds      []interface{} `json:"email_cc_ids"`
	ForumTopicID    interface{}   `json:"forum_topic_id"`
	ProblemID       interface{}   `json:"problem_id"`
	HasIncidents    bool          `json:"has_incidents"`
	IsPublic        bool          `json:"is_public"`
	DueAt           interface{}   `json:"due_at"`
	Tags            []string      `json:"tags"`
	CustomFields    []struct {
		ID    int64       `json:"id"`
		Value interface{} `json:"value"`
	} `json:"custom_fields"`
	SatisfactionRating  interface{}   `json:"satisfaction_rating"`
	SharingAgreementIds []interface{} `json:"sharing_agreement_ids"`
	Fields              []struct {
		ID    int64       `json:"id"`
		Value interface{} `json:"value"`
	} `json:"fields"`
	FollowupIds             []interface{} `json:"followup_ids"`
	TicketFormID            int64         `json:"ticket_form_id"`
	BrandID                 int64         `json:"brand_id"`
	SatisfactionProbability interface{}   `json:"satisfaction_probability"`
	AllowChannelback        bool          `json:"allow_channelback"`
	AllowAttachments        bool          `json:"allow_attachments"`
}
