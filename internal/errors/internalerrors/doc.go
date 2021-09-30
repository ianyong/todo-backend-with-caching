// Package internalerrors contains all internal errors (errors with messages that are not meant to be returned as part
// of the API response). All errors in this package do not implement the api.ExternalError interface and always cause
// the server to return a 500 Internal Server Error as the status code. In addition, all such errors are written to
// the logs.
package internalerrors
