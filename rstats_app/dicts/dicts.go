package dicts

import (
	"embed"
	"encoding/json"
)

type Location struct {
	Name string `json:"name"`
}

type trackJson struct {
	name        string
	length      float64
	location_id int
}

type Track struct {
	Name     string  `json:"name"`
	Length   float64 `json:"length"`
	Location Location
}

type CarClass struct {
	Name       string `json:"name"`
	Drivetrain string `json:"drivetrain"`
}

type carJson struct {
	name     string
	class_id int
}

type Car struct {
	Name  string
	Class CarClass
}

type Dicts struct {
	Tracks map[int]Track
	Cars   map[int]Car
}

//go:embed *.json
var dictsFS embed.FS

func loadJsonOrPanic(name string, value any) {
	data, err := dictsFS.ReadFile(name)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &value); err != nil {
		panic(err)
	}
}

func LoadDicts() Dicts {
	locations := map[int]Location{}
	loadJsonOrPanic("locations.json", &locations)

	carClasses := map[int]CarClass{}
	loadJsonOrPanic("car_classes.json", &carClasses)

	tracksRaw := map[int]trackJson{}
	loadJsonOrPanic("tracks.json", &tracksRaw)
	tracks := map[int]Track{}
	for trackId, track := range tracksRaw {
		tracks[trackId] = Track{
			Name:     track.name,
			Length:   track.length,
			Location: locations[track.location_id],
		}
	}

	carsRaw := map[int]carJson{}
	loadJsonOrPanic("tracks.json", &carsRaw)
	cars := map[int]Car{}
	for carId, car := range carsRaw {
		cars[carId] = Car{
			Name:  car.name,
			Class: carClasses[car.class_id],
		}
	}

	return Dicts{
		Tracks: map[int]Track{},
		Cars:   map[int]Car{},
	}

}
