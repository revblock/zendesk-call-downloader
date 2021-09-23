package zendesk

import "time"

type GetTicketCommentsResponse struct {
	Comments     []Comment   `json:"comments"`
	NextPage     interface{} `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	Count        int         `json:"count"`
}

type Comment struct {
	ID          int64         `json:"id"`
	Type        string        `json:"type"`
	AuthorID    int64         `json:"author_id"`
	Body        string        `json:"body"`
	HTMLBody    string        `json:"html_body"`
	PlainBody   string        `json:"plain_body,omitempty"`
	Public      bool          `json:"public"`
	Attachments []interface{} `json:"attachments"`
	AuditID     int64         `json:"audit_id,omitempty"`
	Via         struct {
		Channel string `json:"channel"`
		Source  struct {
			From struct {
			} `json:"from"`
			To struct {
			} `json:"to"`
			Rel interface{} `json:"rel"`
		} `json:"source"`
	} `json:"via"`
	CreatedAt time.Time `json:"created_at"`
	Data      struct {
		From                   string      `json:"from"`
		To                     string      `json:"to"`
		RecordingURL           *string     `json:"recording_url"`
		RecordingType          string      `json:"recording_type"`
		Recorded               bool        `json:"recorded"`
		RecordingConsentAction interface{} `json:"recording_consent_action"`
		CallID                 int64       `json:"call_id"`
		CallDuration           int         `json:"call_duration"`
		AnsweredByID           int64       `json:"answered_by_id"`
		StartedAt              time.Time   `json:"started_at"`
		Location               string      `json:"location"`
		AuthorID               int64       `json:"author_id"`
		Public                 bool        `json:"public"`
		BrandID                int64       `json:"brand_id"`
		ViaID                  int         `json:"via_id"`
		LineType               string      `json:"line_type"`
		AnsweredByName         string      `json:"answered_by_name"`
		TranscriptionStatus    string      `json:"transcription_status"`
	} `json:"data,omitempty"`
	FormattedFrom        string `json:"formatted_from,omitempty"`
	FormattedTo          string `json:"formatted_to,omitempty"`
	TranscriptionVisible bool   `json:"transcription_visible,omitempty"`
	Trusted              bool   `json:"trusted,omitempty"`
}
