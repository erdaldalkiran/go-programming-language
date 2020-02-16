package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
	fmt.Println("")
	fmt.Println("")
}

type byArtist []*Track

func (b byArtist) Len() int           { return len(b) }
func (b byArtist) Less(i, j int) bool { return b[i].Artist < b[j].Artist }
func (b byArtist) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type byTitle []*Track

func (b byTitle) Len() int           { return len(b) }
func (b byTitle) Less(i, j int) bool { return b[i].Title < b[j].Title }
func (b byTitle) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("byArtistReverse:")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	h := string(33)
	fmt.Printf("%T ###%v###\n", h, h)

	fmt.Println("byTitle:")
	sort.Sort(byTitle(tracks))
	printTracks(tracks)

	fmt.Println("byTitleReverse:")
	sort.Sort(sort.Reverse(byTitle(tracks)))
	printTracks(tracks)
}
