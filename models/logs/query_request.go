package logs

type ListQueryParams struct {
	SourceIDs string `json:"source_ids,omitempty"`

	Query string `json:"query,omitempty"` // Possible values

	Batch int `json:"batch,omitempty"` // 50-1000. Default: 100 rows.

	From string `json:"from,omitempty"` // Start of time range for the log query (ISO8601-formatted string: 2022-07-19T13:32:56+0000). Default: 30 minutes before to. If to is not specified then 30 minutes ago.

	To string `json:"to,omitempty"` // End of time range for the log query (ISO8601-formatted string: 2022-07-19T13:32:56+0000). Default: 30 minutes after from. If from is not specified then current time.

	Order string `json:"order,omitempty"` // Possible values: newest_first, oldest_first. Default: newest_first
}

// strings.Join(qp.SourceIDs, ",")

// q := []{"simple", "compound"}
// o := []{"newest_first", "oldest_first"}
