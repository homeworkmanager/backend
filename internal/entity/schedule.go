package entity

import "time"

type OutputClass struct {
	Class    Class
	Homework []Homework
}

type Day struct {
	Date                time.Time
	OutputClass         []OutputClass
	IndependentHomework []Homework
}
