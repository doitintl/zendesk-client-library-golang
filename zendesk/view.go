package zendesk

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type TicketReference struct {
	ID          int64   `json:"id"`
	Status      *string `json:"status,omitempty"`
	Priority    *string `json:"priority,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

type ViewColumn struct {
	ID         interface{} `json:"id,omitempty"`
	Title      *string     `json:"title,omitempty"`
	Type       *string     `json:"type,omitempty"`
	URL        *string     `json:"url,omitempty"`
	Filterable *bool       `json:"filterable,omitempty"`
	Sortable   *bool       `json:"sortable,omitempty"`
}

type ViewInfo struct {
	ID  *int64  `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`
}

type ViewUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type ViewRow struct {
	CustomFields         []CustomField    `json:"custom_fields,omitempty"`
	TicketID             *int64           `json:"ticket_id"`
	AssigneeID           *int64           `json:"assignee_id,omitempty"`
	RequesterID          *int64           `json:"requester_id,omitempty"`
	SubmitterID          *int64           `json:"submitter_id,omitempty"`
	GroupID              *int64           `json:"group_id,omitempty"`
	OrganizationID       *int64           `json:"organization_id,omitempty"`
	Subject              *string          `json:"subject,omitempty"`
	Priority             *string          `json:"priority,omitempty"`
	SatisfactionScore    *string          `json:"satisfaction_score,omitempty"`
	UpdatedByType        *string          `json:"updated_by_type,omitempty"`
	Locale               *string          `json:"locale,omitempty"`
	TicketForm           *string          `json:"ticket_form,omitempty"`
	Brand                *string          `json:"brand,omitempty"`
	Type                 *string          `json:"type,omitempty"`
	Created              *time.Time       `json:"created,omitempty"`
	Updated              *time.Time       `json:"updated,omitempty"`
	Assigned             *time.Time       `json:"assigned,omitempty"`
	Solved               *time.Time       `json:"solved,omitempty"`
	SLANextBreachAt      *time.Time       `json:"sla_next_breach_at,omitempty"`
	GroupSLANextBreachAt *time.Time       `json:"group_sla_next_breach_at,omitempty"`
	RequesterUpdatedAt   *time.Time       `json:"requester_updated_at,omitempty"`
	AssigneeUpdatedAt    *time.Time       `json:"assignee_updated_at,omitempty"`
	Ticket               *TicketReference `json:"ticket,omitempty"`
}

type ViewResults struct {
	Rows         []ViewRow    `json:"rows"`
	Columns      []ViewColumn `json:"columns"`
	View         ViewInfo     `json:"view"`
	Users        []ViewUser   `json:"users"`
	NextPage     *string      `json:"next_page"`
	PreviousPage *string      `json:"previous_page"`
	Count        *int64       `json:"count"`
}

func (c *client) ExecuteView(id int64, options *ListOptions) (*ViewResults, error) {
	params, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	out := new(ViewResults)
	err = c.get(fmt.Sprintf("/api/v2/views/%d/execute.json?%s", id, params.Encode()), out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
