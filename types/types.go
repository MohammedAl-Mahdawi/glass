package types

import "time"

type Common struct {
	TmpPath       string
	StartedAt     time.Time
	FailedStatus  string
	SuccessStatus string
}

type Log struct {
	Command     string    `yaml:"command"`
	UUID        string    `yaml:"uuid"`
	CreatedAt   time.Time `yaml:"created-at"`
	CompletedAt time.Time `yaml:"completed-at"`
	ExitStatus  int       `yaml:"exit-status"`
	Status      string    `yaml:"status"`
	Log         string    `yaml:"log"`
}
