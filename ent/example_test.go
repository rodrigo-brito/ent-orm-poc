// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"

	"enttest/ent/artist"
	"enttest/ent/song"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleArtist() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the artist's edges.
	s0 := client.Song.
		Create().
		SetName(1).
		SetPhone("string").
		SetPlays(1).
		SetGender(song.GenderF).
		SaveX(ctx)
	log.Println("song created:", s0)

	// create artist vertex with its edges.
	a := client.Artist.
		Create().
		SetSlug("string").
		SetAge(1).
		SetPhone("string").
		SetPlays(1).
		SetGender(artist.GenderF).
		AddSongs(s0).
		SaveX(ctx)
	log.Println("artist created:", a)

	// query edges.
	s0, err = a.QuerySongs().First(ctx)
	if err != nil {
		log.Fatalf("failed querying songs: %v", err)
	}
	log.Println("songs found:", s0)

	// Output:
}
func ExampleSong() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the song's edges.

	// create song vertex with its edges.
	s := client.Song.
		Create().
		SetName(1).
		SetPhone("string").
		SetPlays(1).
		SetGender(song.GenderF).
		SaveX(ctx)
	log.Println("song created:", s)

	// query edges.

	// Output:
}
