package model

import (
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"time"
)

// Metadata defines the movie metadata.
type Metadata struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	CreateDate  string `json:"-"`
}

type Movies []*Metadata

func (p *Movies) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetMovies() Movies {
	return movieList
}

func GetMovie(id int32) *Metadata {
	return movieList[id]
}

func Add(m *Metadata) Movies {
	m.ID = int32(len(movieList) + 1)
	m.CreateDate = time.Now().UTC().String()
	movieList = append(movieList, m)
	return movieList
}
func Update(m *Metadata, id int32) Movies {
	for i, d := range movieList {
		if id == d.ID {
			m.ID = d.ID
			m.CreateDate = d.CreateDate
			movieList[i] = m
			break
		}
	}
	return movieList
}

// DeleteMovie removes a movie from the list.
func DeleteMovie(id int32) Movies {
	fmt.Println("inside delete movie")
	fmt.Printf("movieid before delete %d\n", id)
	for i, movie := range movieList {
		if movie.ID == id {
			//movieList = append(movieList[:i], movieList[i+1:]...)
			movieList = slices.Delete(movieList, i, i+1)
			break
		}
	}
	return movieList
}

// DeleteMovie removes a movie from the list.
// DeleteMovie removes a movie from the list.
// DeleteMovie removes a movie from the list.
// Delete removes a movie from the list.
var movieList = []*Metadata{
	{
		ID:          1,
		Title:       "Tare Zameen",
		Description: "Brave ",
		Director:    "AK",
		CreateDate:  time.Now().UTC().String(),
	},
	{
		ID:          2,
		Title:       "3 ID",
		Description: "Copy",
		Director:    "R K H",
		CreateDate:  time.Now().UTC().String(),
	},
}

func (p *Metadata) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func (p *Metadata) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
