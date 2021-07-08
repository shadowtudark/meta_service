package model

// Partition refers to a partition meta info
type Partition struct {
	JobId         string             `json:"job_id"`
	PartitionNum  string             `json:"partition_num"`
	Leader        string             `json:"leader"`
	IsInSync      []string           `json:"is_in_sync"`
	Followers     map[string][]int64 `json:"followers"`
	LowWaterMark  int64              `json:"low_water_mark"`
	HighWaterMark int64              `json:"high_water_mark"`
}

// NestedInfo referes to all partitions info
type NestedInfo struct {
	JobId      string      `json:"job_id"`
	Partitions []Partition `json:"partitions"`
}
