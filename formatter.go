package logman

type Formatter interface {
	Format(entry *Entry) error
}
