package httptool

const (
	HEADER_TRACEID  string = "Trace-Id"
	HEADER_OPERATOR string = "Operator"
	// HEADER_DURATION duration HTTP header key
	HEADER_DURATION string = "Duration"
	// CONTEXTKEY_TRACEID trace-id HTTP header key
	CONTEXTKEY_TRACEID int = 0
	// CONTEXTKEY_OPERATOR operator HTTP header key
	CONTEXTKEY_OPERATOR int = 1
	// CONTENTTYPE_JSON json content type value
	CONTENTTYPE_JSON string = "application/json; charset=utf-8"
)
