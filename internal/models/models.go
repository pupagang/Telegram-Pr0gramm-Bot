package models

// pr0gramm response
type Pr0Response struct {
	AtEnd   bool        `json:"atEnd"`
	AtStart bool        `json:"atStart"`
	Error   interface{} `json:"error"`
	Items   []Items     `json:"items"`
	Ts      int         `json:"ts"`
	Cache   string      `json:"cache"`
	Rt      int         `json:"rt"`
	Qc      int         `json:"qc"`
}

type Items struct {
	ID       int    `json:"id"`
	Promoted int    `json:"promoted"`
	UserID   int    `json:"userId"`
	Up       int    `json:"up"`
	Down     int    `json:"down"`
	Created  int    `json:"created"`
	Image    string `json:"image"`
	Thumb    string `json:"thumb"`
	Fullsize string `json:"fullsize"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Audio    bool   `json:"audio"`
	Source   string `json:"source"`
	Flags    int    `json:"flags"`
	User     string `json:"user"`
	Mark     int    `json:"mark"`
	Gift     int    `json:"gift"`
}

type PostItem struct {
	MediaURL string
	Caption  string
	IsVideo  bool
}
