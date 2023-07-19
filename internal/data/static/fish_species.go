package static

var ShallowFishNames = []string{
	"Atlantic bluefin tuna",
	"Blue tang",
	"Cod",
	"Dolphinfish",
	"Grouper",
	"Halibut",
	"Mahi-mahi",
	"Marlin",
	"Salmon",
	"Swordfish",
	"Atlantic cod",
	"Bigeye tuna",
	"Black seabass",
	"Blacktip shark",
	"Bluefin tuna",
	"Butterfish",
}

var MiddleFishNames = []string{
	"California halibut",
	"Capelin",
	"Catfish",
	"Chub mackerel",
	"Coho salmon",
	"Common dolphinfish",
	"Dover sole",
	"Flounder",
	"Giant grouper",
	"Golden trevally",
	"Gray triggerfish",
	"Haddock",
	"Shark",
	"Skipjack tuna",
	"Skate",
	"Sole",
	"Swordfish",
	"Tilefish",
	"Yellowfin tuna",
}

var DeepFishNames = []string{
	"Hake",
	"Hoki",
	"King mackerel",
	"Largemouth bass",
	"Lemon shark",
	"Lingcod",
	"Longfin squid",
	"Mahi-mahi",
	"Marlin",
	"Menhaden",
	"Mullet",
	"Napoleon wrasse",
	"Ocean perch",
	"Pacific halibut",
	"Pacific salmon",
	"Perch",
	"Pink salmon",
	"Pompano",
	"Red snapper",
	"Rockfish",
	"Salmon",
	"Scallop",
}

const (
	ShallowLevel                    = 200
	MiddleLevel                     = 400
	DeepLevel                       = 700
	MinShallowLevelFishSpeciesCount = 10
	MaxShallowLevelFishSpeciesCount = 50
	MinMiddleLevelFishSpeciesCount  = 5
	MaxMiddleLevelFishSpeciesCount  = 25
	MinDeepLevelFishSpeciesCount    = 0
	MaxDeepLevelFishSpeciesCount    = 10
)
