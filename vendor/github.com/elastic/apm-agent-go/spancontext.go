package apm

import (
	"net/http"

	"go.elastic.co/apm/model"
)

// SpanContext provides methods for setting span context.
type SpanContext struct {
	model    model.SpanContext
	database model.DatabaseSpanContext
	http     model.HTTPSpanContext
}

// DatabaseSpanContext holds database span context.
type DatabaseSpanContext struct {
	// Instance holds the database instance name.
	Instance string

	// Statement holds the statement executed in the span,
	// e.g. "SELECT * FROM foo".
	Statement string

	// Type holds the database type, e.g. "sql".
	Type string

	// User holds the username used for database access.
	User string
}

func (c *SpanContext) build() *model.SpanContext {
	switch {
	case len(c.model.Tags) != 0:
	case c.model.Database != nil:
	case c.model.HTTP != nil:
	default:
		return nil
	}
	return &c.model
}

func (c *SpanContext) reset() {
	// TODO(axw) reuse space for tags
	*c = SpanContext{}
}

// SetTag sets a tag in the context. If the key is invalid
// (contains '.', '*', or '"'), the call is a no-op.
func (c *SpanContext) SetTag(key, value string) {
	if !validTagKey(key) {
		return
	}
	value = truncateString(value)
	if c.model.Tags == nil {
		c.model.Tags = map[string]string{key: value}
	} else {
		c.model.Tags[key] = value
	}
}

// SetDatabase sets the span context for database-related operations.
func (c *SpanContext) SetDatabase(db DatabaseSpanContext) {
	c.database = model.DatabaseSpanContext{
		Instance:  truncateString(db.Instance),
		Statement: truncateLongString(db.Statement),
		Type:      truncateString(db.Type),
		User:      truncateString(db.User),
	}
	c.model.Database = &c.database
}

// SetHTTPRequest sets the details of the HTTP request in the context.
//
// This function relates to client requests. If the request URL contains
// user info, it will be removed and excluded from the stored URL.
func (c *SpanContext) SetHTTPRequest(req *http.Request) {
	c.http.URL = req.URL
	c.model.HTTP = &c.http
}
