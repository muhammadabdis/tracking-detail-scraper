package main

type Tracking struct {
	Tracks []Track
}

func (t *Tracking) Find() []Track {
	return t.Tracks
}

func (t *Tracking) Create(tr Track) Track {
	t.Tracks = append(t.Tracks, tr)

	return tr
}
