ALTER TABLE {{.Table | sanitize}} ADD COLUMN {{with .Column}}{{.Name | sanitize}} {{.Type}}{{end}}