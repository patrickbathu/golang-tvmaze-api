package models

// Episode representa um episódio na TVMaze API
type Episode struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Season  int    `json:"season"`
	Number  int    `json:"number"`
	Airdate string `json:"airdate"`
	Airtime string `json:"airtime"`
	Runtime int    `json:"runtime"`
	Summary string `json:"summary"`
	Image   *Image `json:"image"`
}

// Show representa um show de TV
type Show struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Language  string   `json:"language"`
	Genres    []string `json:"genres"`
	Status    string   `json:"status"`
	Premiered string   `json:"premiered"`
	Summary   string   `json:"summary"`
	Image     *Image   `json:"image"`
	Network   *Network `json:"network"`
}

// Network representa a rede de TV
type Network struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}

// Country representa o país
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Image representa as imagens
type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

// Schedule representa um item da programação
type Schedule struct {
	ID      int      `json:"id"`
	Airdate string   `json:"airdate"`
	Airtime string   `json:"airtime"`
	Show    Show     `json:"show"`
	Episode *Episode `json:"episode,omitempty"`
}
