package shared

import (
)

type Asset struct {
	Id  			string				`json:"id"`
	Name 			string				`json:"name"`
	Type 			string				`json:"type"`
	Desc 			string				`json:"description"`
	OwnerIds 	[]string			`json:"owner"`
	Tags 			[]string			`json:"tags"`
}

type Source struct {
	Id  	string						`json:"id"`
	Url 	string						`json:"url"`
	Tags 	[]string					`json:"tags"`
}

type Event struct {
	Id  			string				`json:"id"`
	EntityIds	[]string			`json:"entities"`
	Date 			string				`json:"date"`
	Desc 			string				`json:"description"`
	SourceIds	[]string			`json:"sources"`
	Tags			[]string			`json:"tags"`
}

type Assocation struct {
	EntityIds	[]string			`json:"entities"`
	Desc 			string				`json:"description"`
	SourceIds	[]string			`json:"sources"`
	Tags 			[]string			`json:"tags"`
}

type Entity struct {
	Id  				string			`json:"id"`
	Name 				string			`json:"name"`
	Type 				string			`json:"type"`
	Nationality string			`json:"nationality"`
	Notes				string			`json:"notes"`
	Tags 				[]string		`json:"tags"`
}
