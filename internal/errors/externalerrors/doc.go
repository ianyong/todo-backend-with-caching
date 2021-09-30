// Package externalerrors contains all external errors (errors with messages that are meant to be returned as part of
// the API response). All errors in this package implement the api.ExternalError interface and never cause the server
// to return a 500 Internal Server Error as the status code.
package externalerrors
