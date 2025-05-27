package model

type Setup struct {
	Containers []Container
	Mocks      []Mock
}

type Container struct {
	Name  string
	Image string
	Tag   string
	Port  string
}

type Mock struct {
	Name string
	File *string
}
