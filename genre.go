package ncs

type Genre string

// genres
const (
	AlternativeDance Genre = "Alternative Dance"
	AlternativePop   Genre = "Alternative Pop"
	Ambient          Genre = "Ambient"
	AntiPop          Genre = "Anti-Pop"
	Bass             Genre = "Bass"
	BassHouse        Genre = "Bass House"
	BrazilianPhonk   Genre = "Brazilian Phonk"
	Breakbeat        Genre = "Breakbeat"
	Chill            Genre = "Chill"
	ChillPop         Genre = "Chill Pop"
	ColourBass       Genre = "Colour Bass"
	Complextro       Genre = "Complextro"
	DancePop         Genre = "Dance-Pop"
	DeepHouse        Genre = "Deep House"
	DiscoHouse       Genre = "Disco House"
	DrumBass         Genre = "Drum & Bass"
	Drumstep         Genre = "Drumstep"
	Dubstep          Genre = "Dubstep"
	EDM              Genre = "EDM"
	Electro          Genre = "Electro"
	Electronic       Genre = "Electronic"
	ElectronicPop    Genre = "Electronic Pop"
	ElectronicRock   Genre = "Electronic Rock"
	FutureBass       Genre = "Future Bass"
	FutureBounce     Genre = "Future Bounce"
	FutureHouse      Genre = "Future House"
	FutureRave       Genre = "Future Rave"
	FutureTrap       Genre = "Future Trap"
	Garage           Genre = "Garage"
	GlitchHop        Genre = "Glitch Hop"
	Hardstyle        Genre = "Hardstyle"
	House            Genre = "House"
	Hyperpop         Genre = "Hyperpop"
	IndieDance       Genre = "Indie Dance"
	JerseyClub       Genre = "Jersey Club"
	JumpUp           Genre = "Jump-Up"
	LiquidDnB        Genre = "Liquid DnB"
	LofiHipHop       Genre = "Lofi Hip-Hop"
	MelodicDubstep   Genre = "Melodic Dubstep"
	MelodicHouse     Genre = "Melodic House"
	MidtempoBass     Genre = "Midtempo Bass"
	Neurofunk        Genre = "Neurofunk"
	Phonk            Genre = "Phonk"
	Pluggnb          Genre = "Pluggnb"
	Pop              Genre = "Pop"
	ProgressiveHouse Genre = "Progressive House"
	TechHouse        Genre = "Tech House"
	Techno           Genre = "Techno"
	Trap             Genre = "Trap"
	TribalHouse      Genre = "Tribal House"
	UKG              Genre = "UKG"
)

// GenreMap
var GenreMap = map[Genre]int{
	AlternativeDance: 31,
	AlternativePop:   33,
	Ambient:          23,
	AntiPop:          34,
	Bass:             1,
	BassHouse:        18,
	BrazilianPhonk:   26,
	Breakbeat:        27,
	Chill:            2,
	ChillPop:         35,
	ColourBass:       85,
	Complextro:       65,
	DancePop:         36,
	DeepHouse:        66,
	DiscoHouse:       46,
	DrumBass:         3,
	Drumstep:         4,
	Dubstep:          5,
	EDM:              6,
	Electro:          47,
	Electronic:       7,
	ElectronicPop:    39,
	ElectronicRock:   83,
	FutureBass:       17,
	FutureBounce:     68,
	FutureHouse:      8,
	FutureRave:       69,
	FutureTrap:       57,
	Garage:           51,
	GlitchHop:        15,
	Hardstyle:        9,
	House:            10,
	Hyperpop:         41,
	IndieDance:       11,
	JerseyClub:       84,
	JumpUp:           28,
	LiquidDnB:        29,
	LofiHipHop:       60,
	MelodicDubstep:   12,
	MelodicHouse:     54,
	MidtempoBass:     22,
	Neurofunk:        30,
	Phonk:            16,
	Pluggnb:          86,
	Pop:              19,
	ProgressiveHouse: 55,
	TechHouse:        73,
	Techno:           80,
	Trap:             14,
	TribalHouse:      74,
	UKG:              21,
}
