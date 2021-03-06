// Package database provides the Chrome Debugging Protocol
// commands, types, and events for the Database domain.
//
// Generated by the chromedp-gen command.
package database

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// EnableParams enables database tracking, database events will now be
// delivered to the client.
type EnableParams struct{}

// Enable enables database tracking, database events will now be delivered to
// the client.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Database.enable against the provided context and
// target handler.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDatabaseEnable, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// DisableParams disables database tracking, prevents database events from
// being sent to the client.
type DisableParams struct{}

// Disable disables database tracking, prevents database events from being
// sent to the client.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Database.disable against the provided context and
// target handler.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDatabaseDisable, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// GetDatabaseTableNamesParams [no description].
type GetDatabaseTableNamesParams struct {
	DatabaseID ID `json:"databaseId"`
}

// GetDatabaseTableNames [no description].
//
// parameters:
//   databaseID
func GetDatabaseTableNames(databaseID ID) *GetDatabaseTableNamesParams {
	return &GetDatabaseTableNamesParams{
		DatabaseID: databaseID,
	}
}

// GetDatabaseTableNamesReturns return values.
type GetDatabaseTableNamesReturns struct {
	TableNames []string `json:"tableNames,omitempty"`
}

// Do executes Database.getDatabaseTableNames against the provided context and
// target handler.
//
// returns:
//   tableNames
func (p *GetDatabaseTableNamesParams) Do(ctxt context.Context, h cdp.Handler) (tableNames []string, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDatabaseGetDatabaseTableNames, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetDatabaseTableNamesReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, cdp.ErrInvalidResult
			}

			return r.TableNames, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, ctxt.Err()
	}

	return nil, cdp.ErrUnknownResult
}

// ExecuteSQLParams [no description].
type ExecuteSQLParams struct {
	DatabaseID ID     `json:"databaseId"`
	Query      string `json:"query"`
}

// ExecuteSQL [no description].
//
// parameters:
//   databaseID
//   query
func ExecuteSQL(databaseID ID, query string) *ExecuteSQLParams {
	return &ExecuteSQLParams{
		DatabaseID: databaseID,
		Query:      query,
	}
}

// ExecuteSQLReturns return values.
type ExecuteSQLReturns struct {
	ColumnNames []string              `json:"columnNames,omitempty"`
	Values      []easyjson.RawMessage `json:"values,omitempty"`
	SQLError    *Error                `json:"sqlError,omitempty"`
}

// Do executes Database.executeSQL against the provided context and
// target handler.
//
// returns:
//   columnNames
//   values
//   sqlError
func (p *ExecuteSQLParams) Do(ctxt context.Context, h cdp.Handler) (columnNames []string, values []easyjson.RawMessage, sqlError *Error, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, nil, nil, err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDatabaseExecuteSQL, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, nil, nil, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r ExecuteSQLReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, nil, nil, cdp.ErrInvalidResult
			}

			return r.ColumnNames, r.Values, r.SQLError, nil

		case error:
			return nil, nil, nil, v
		}

	case <-ctxt.Done():
		return nil, nil, nil, ctxt.Err()
	}

	return nil, nil, nil, cdp.ErrUnknownResult
}
