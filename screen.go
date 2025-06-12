package main

import (
	nato "github.com/benjamindimalanta/natomanga-api"
	
)

var screen Screen

// contains every main views and nato.Searcher
type Screen struct {

	// gocui's View with some extra fields
	sb *views.SearchBar
	sl *views.SearchList
	md *views.MangaDetails
	cl *views.ChapterList

	// natomanga-api's Searcher struct
	searcher nato.Searcher
}
