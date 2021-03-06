// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// Counters represents a row from '[custom counters]'.
type Counters struct {
	Datname      Name  // datname
	Numbackends  int   // numbackends
	XactCommit   int64 // xact_commit
	XactRollback int64 // xact_rollback
	BlksRead     int64 // blks_read
	BlksHit      int64 // blks_hit
	TupReturned  int64 // tup_returned
	TupFetched   int64 // tup_fetched
	TupInserted  int64 // tup_inserted
	TupUpdated   int64 // tup_updated
	TupDeleted   int64 // tup_deleted
	Conflicts    int64 // conflicts
	TempFiles    int64 // temp_files
	TempBytes    int64 // temp_bytes
	Deadlocks    int64 // deadlocks
}

// GetCounters runs a custom query, returning results as Counters.
func GetCounters(db XODB) ([]*Counters, error) {
	var err error

	// sql query
	var sqlstr = `SELECT datname, numbackends, xact_commit, xact_rollback, ` +
		`blks_read, blks_hit, tup_returned, tup_fetched, tup_inserted, ` +
		`tup_updated, tup_deleted, conflicts, temp_files, ` +
		`temp_bytes, deadlocks ` +
		`FROM pg_stat_database ` +
		`ORDER BY datname`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Counters{}
	for q.Next() {
		c := Counters{}

		// scan
		err = q.Scan(&c.Datname, &c.Numbackends, &c.XactCommit, &c.XactRollback, &c.BlksRead, &c.BlksHit, &c.TupReturned, &c.TupFetched, &c.TupInserted, &c.TupUpdated, &c.TupDeleted, &c.Conflicts, &c.TempFiles, &c.TempBytes, &c.Deadlocks)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}
