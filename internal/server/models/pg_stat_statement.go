package models

type PgStatStatement struct {
	Query        string  `gorm:"query" json:"query"`
	MinTimeSecs  float64 `gorm:"min_time_secs" json:"min_time_secs"`
	MaxTimeSecs  float64 `gorm:"max_time_secs" json:"max_time_secs"`
	MeanTimeSecs float64 `gorm:"mean_time_secs" json:"mean_time_secs"`
	Database     string  `gorm:"database"`
}

func (this PgStatStatement) TableName() string {
	return "pg_stat_statements"
}
