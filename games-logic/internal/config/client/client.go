package client

type ClientList struct {
	Players  Rest `json:"players"`
	Games    Rest `json:"games"`
	GamePlay Rest `json:"game_play"`
}

type Rest struct {
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
}
