package ncs

type Mood string

// Mood constants
const (
	Angry       Mood = "Angry"
	AngryLower  Mood = "angry"
	Chasing     Mood = "Chasing"
	Dark        Mood = "Dark"
	DarkLower   Mood = "dark"
	Dramatic    Mood = "dramatic"
	Dreamy      Mood = "Dreamy"
	Eccentric   Mood = "Eccentric"
	Elegant     Mood = "Elegant"
	Energetic   Mood = "energetic"
	Epic        Mood = "Epic"
	Euphoric    Mood = "Euphoric"
	Exciting    Mood = "exciting"
	Fear        Mood = "Fear"
	Floating    Mood = "Floating"
	Funny       Mood = "Funny"
	Glamorous   Mood = "Glamorous"
	Gloomy      Mood = "Gloomy"
	Happy       Mood = "Happy"
	Heavy       Mood = "Heavy"
	Hopeful     Mood = "Hopeful"
	LaidBack    Mood = "Laid Back"
	Mysterious  Mood = "Mysterious"
	Negative    Mood = "negative"
	Neutral     Mood = "neutral"
	Peaceful    Mood = "Peaceful"
	Positive    Mood = "positive"
	Powerful    Mood = "powerful"
	Quirky      Mood = "Quirky"
	Relaxed     Mood = "relaxed"
	Restless    Mood = "Restless"
	Romantic    Mood = "romantic"
	Sad         Mood = "sad"
	Scary       Mood = "scary"
	Sentimental Mood = "Sentimental"
	Sexy        Mood = "Sexy"
	Suspense    Mood = "Suspense"
	Weird       Mood = "Weird"
)

// MoodMap
var MoodMap = map[Mood]int{
	Angry:       1,
	AngryLower:  33,
	Chasing:     26,
	Dark:        2,
	DarkLower:   35,
	Dramatic:    36,
	Dreamy:      3,
	Eccentric:   27,
	Elegant:     28,
	Energetic:   6,
	Epic:        4,
	Euphoric:    5,
	Exciting:    37,
	Fear:        7,
	Floating:    29,
	Funny:       8,
	Glamorous:   9,
	Gloomy:      10,
	Happy:       11,
	Heavy:       30,
	Hopeful:     12,
	LaidBack:    13,
	Mysterious:  14,
	Negative:    34,
	Neutral:     39,
	Peaceful:    15,
	Positive:    40,
	Powerful:    32,
	Quirky:      16,
	Relaxed:     17,
	Restless:    18,
	Romantic:    19,
	Sad:         20,
	Scary:       21,
	Sentimental: 31,
	Sexy:        22,
	Suspense:    23,
	Weird:       24,
}
