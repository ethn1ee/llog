package handler

const (
	dbGetEntryError    = "failed to retrieve entries from db: %w"
	dbAddEntryError    = "failed to create entries in db: %w"
	dbDeleteEntryError = "failed to delete entries from db: %w"
)

const (
	flagMutexError   = "option '%s' cannot be used with option '%s'"
	flagIdMutexError = "cannot use flags when providing IDs"
)

const (
	dateParseError = "failed to parse date: %w"
	dateRangeError = "invalid date range: %s"
)

const (
	idParseError = "invalid ID: %w"
	idVoidError  = "invalid ID: ID %d does not exist"
)

const (
	noArgsOrFlagsError = "arguments or flags must be provided"
)
