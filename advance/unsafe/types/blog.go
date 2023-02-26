package types

import "time"

type BlogV1 struct {
	Name string
	Id int32
	Content []byte
	Title string
	Created time.Time
	Updated time.Time
}

type BlogV2 struct {
	Name string
	Id int32
	ParentId int32
	Content []byte
	Title string
	Created time.Time
	Updated time.Time
}

type BlogV3 struct {
	Name string
	Id int32
	Content []byte
	Title string
	Created time.Time
	Updated time.Time
	ParentId int32
}

