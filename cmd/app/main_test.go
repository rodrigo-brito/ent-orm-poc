package main

import (
	"context"
	"enttest/ent"
	"enttest/ent/artist"
	"testing"

	"github.com/stretchr/testify/require"
)

func getTestClient(t *testing.T) *ent.Client {
	t.Helper()

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Error(err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		t.Error(err)
	}

	return client
}

func TestGetArtist(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name   string
		before func(*testing.T, *ent.Client)
		expect func(*testing.T, *ent.Client)
	}{
		{
			name: "get artist/found",
			before: func(t *testing.T, client *ent.Client) {
				_, err := client.Artist.Create().
					SetSlug("brito").
					SetAge(10).
					SetPlays(100).
					SetGender(artist.GenderM).
					Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				artist, err := GetArtist(ctx, client, 1)
				require.Nil(t, err)
				require.Equal(t, "brito", artist.Slug)
			},
		},
		{
			name:   "get artist/not found",
			before: func(t *testing.T, client *ent.Client) {},
			expect: func(t *testing.T, client *ent.Client) {
				artist, err := GetArtist(ctx, client, 1)
				require.Nil(t, err)
				require.Nil(t, artist)
			},
		},
		{
			name: "get artist/fail",
			before: func(t *testing.T, client *ent.Client) {
				client.Close()
			},
			expect: func(t *testing.T, client *ent.Client) {
				artist, err := GetArtist(ctx, client, 1)
				require.EqualError(t, err, "sql: database is closed")
				require.Nil(t, artist)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := getTestClient(t)
			defer client.Close()

			tt.before(t, client)
			tt.expect(t, client)
		})
	}
}

func TestSaveArtist(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name   string
		before func(*testing.T, *ent.Client)
		expect func(*testing.T, *ent.Client)
	}{
		{
			name:   "save artist/success",
			before: func(t *testing.T, client *ent.Client) {},
			expect: func(t *testing.T, client *ent.Client) {
				err := CreateArtist(ctx, client, "brito")
				require.Nil(t, err)
				artist, err := client.Artist.Get(ctx, 1)
				require.Nil(t, err)
				require.Equal(t, "brito", artist.Slug)
			},
		},
		{
			name: "create artist/fail",
			before: func(t *testing.T, client *ent.Client) {
				client.Close()
			},
			expect: func(t *testing.T, client *ent.Client) {
				err := CreateArtist(ctx, client, "fail")
				require.EqualError(t, err, "sql: database is closed")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := getTestClient(t)
			defer client.Close()

			tt.before(t, client)
			tt.expect(t, client)
		})
	}
}
