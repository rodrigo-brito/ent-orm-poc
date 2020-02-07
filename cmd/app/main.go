package main

import (
	"context"
	"enttest/ent"
	"enttest/ent/artist"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		ent.Log(logrus.Info))
	if err != nil {
		logrus.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}

	if err = CreateArtist(ctx, client, "brito"); err != nil {
		logrus.Fatal(err)
	}

	if _, err = GetArtist(ctx, client, 1); err != nil {
		logrus.Fatal(err)
	}

	if _, err = GetArtist(ctx, client, 404); err != nil {
		logrus.Fatal("bla", err)
	}

	if _, err = UpdateArtist(ctx, client, 1, "batata"); err != nil {
		logrus.Fatal("foo", err)
	}

	if _, err = GetArtist(ctx, client, 404); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("Finished!")
}

func GetArtist(ctx context.Context, client *ent.Client, ID int) (*ent.Artist, error) {
	artist, err := client.Artist.Get(ctx, ID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if artist == nil {
		logrus.Println("Artist not found!")
		return nil, nil
	}

	logrus.Printf("Artist found: %d - %s", artist.ID, artist.Slug)
	return artist, nil
}

func CreateArtist(ctx context.Context, client *ent.Client, slug string) error {
	artist, err := client.Artist.Create().
		SetAge(26).
		SetGender(artist.GenderM).
		SetPlays(1000).
		SetSlug(slug).
		SetPhone("6666-6666").Save(ctx)
	if err != nil {
		return err
	}

	logrus.Printf("Artist saved = %d\n", artist.ID)
	return nil
}

func UpdateArtist(ctx context.Context, client *ent.Client, ID int, slug string) (*ent.Artist, error) {
	artist, err := client.Artist.UpdateOneID(ID).
		SetSlug(slug).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	logrus.Printf("Artist updated = %s\n", artist.Slug)

	return artist, nil
}
