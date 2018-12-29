// Copyright 2015 The Loadcat Authors. All rights reserved.

package data

type Availability string

var Availabilities = []Availability{
	"unavailable",
	"available",
	"backup",
}

func (a Availability) Label() string {
	return AvailabilityLabels[a]
}

var AvailabilityLabels = map[Availability]string{
	"unavailable": "不可用",
	"available":   "可用",
	"backup":      "备用",
}
