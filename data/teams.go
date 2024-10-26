package data

type Rank struct {
	Ovr float64 `json:"ovr"`
	Off float64 `json:"off"`
	Def float64 `json:"def"`
}

type Team struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	Rank       Rank   `json:"rank"`
	Division   string `json:"division"`
	Conference string `json:"conference"`
}

var Teams = []Team{
	{
		Key:        "ducks",
		Name:       "Ducks",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 84, Off: 84, Def: 84},
	},
	{
		Key:        "bruins",
		Name:       "Bruins",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 90, Off: 85, Def: 95},
	},
	{
		Key:        "sabres",
		Name:       "Sabres",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 75, Off: 70, Def: 80},
	},
	{
		Key:        "redwings",
		Name:       "Red Wings",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 78, Off: 75, Def: 80},
	},
	{
		Key:        "panthers",
		Name:       "Panthers",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 88, Off: 90, Def: 85},
	},
	{
		Key:        "canadiens",
		Name:       "Canadiens",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 80, Off: 82, Def: 78},
	},
	{
		Key:        "senators",
		Name:       "Senators",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 74, Off: 70, Def: 78},
	},
	{
		Key:        "lightning",
		Name:       "Lightning",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 92, Off: 93, Def: 91},
	},
	{
		Key:        "mapleleafs",
		Name:       "Maple Leafs",
		Division:   "Atlantic",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 85, Off: 88, Def: 82},
	},
	{
		Key:        "hurricanes",
		Name:       "Hurricanes",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 87, Off: 85, Def: 89},
	},
	{
		Key:        "bluejackets",
		Name:       "Blue Jackets",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 76, Off: 74, Def: 78},
	},
	{
		Key:        "devils",
		Name:       "Devils",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 80, Off: 83, Def: 77},
	},
	{
		Key:        "islanders",
		Name:       "Islanders",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 82, Off: 81, Def: 83},
	},
	{
		Key:        "rangers",
		Name:       "Rangers",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 84, Off: 86, Def: 82},
	},
	{
		Key:        "flyers",
		Name:       "Flyers",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 79, Off: 78, Def: 80},
	},
	{
		Key:        "penguins",
		Name:       "Penguins",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 89, Off: 91, Def: 87},
	},
	{
		Key:        "capitals",
		Name:       "Capitals",
		Division:   "Metropolitan",
		Conference: "Eastern",
		Rank:       Rank{Ovr: 86, Off: 87, Def: 85},
	},
	{
		Key:        "blackhawks",
		Name:       "Blackhawks",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 82, Off: 80, Def: 84},
	},
	{
		Key:        "avalanche",
		Name:       "Avalanche",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 92, Off: 94, Def: 90},
	},
	{
		Key:        "stars",
		Name:       "Stars",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 88, Off: 87, Def: 89},
	},
	{
		Key:        "wild",
		Name:       "Wild",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 85, Off: 84, Def: 86},
	},
	{
		Key:        "predators",
		Name:       "Predators",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 83, Off: 82, Def: 84},
	},
	{
		Key:        "blues",
		Name:       "Blues",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 87, Off: 86, Def: 88},
	},
	{
		Key:        "jets",
		Name:       "Jets",
		Division:   "Central",
		Conference: "Western",
		Rank:       Rank{Ovr: 84, Off: 85, Def: 83},
	},
	{
		Key:        "coyotes",
		Name:       "Coyotes",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 70, Off: 68, Def: 72},
	},
	{
		Key:        "flames",
		Name:       "Flames",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 82, Off: 81, Def: 83},
	},
	{
		Key:        "oilers",
		Name:       "Oilers",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 90, Off: 93, Def: 87},
	},
	{
		Key:        "kings",
		Name:       "Kings",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 80, Off: 79, Def: 81},
	},
	{
		Key:        "sharks",
		Name:       "Sharks",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 78, Off: 77, Def: 79},
	},
	{
		Key:        "kraken",
		Name:       "Kraken",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 77, Off: 76, Def: 78},
	},
	{
		Key:        "canucks",
		Name:       "Canucks",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 81, Off: 82, Def: 80},
	},
	{
		Key:        "goldenknights",
		Name:       "Golden Knights",
		Division:   "Pacific",
		Conference: "Western",
		Rank:       Rank{Ovr: 88, Off: 90, Def: 86},
	},
}

var GoalsOddsPerPeriod = map[float64]float64{
	0: 25.0,
	1: 20.0,
	2: 15.0,
	3: 5.0,
	4: 2.0,
	5: 1.0,
	6: 0.1,
}
