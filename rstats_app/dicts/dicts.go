package dicts

import (
	"embed"
	"encoding/json"
	"fmt"

	"dev.maizy.ru/rstats/internal/u"
)

type Location struct {
	Name string `json:"name"`
}

var UnknownLocation = Location{Name: "???"}

func (l Location) String() string {
	return l.Name
}

type trackJson struct {
	Name       string  `json:"name"`
	Length     float64 `json:"length"`
	LocationId int     `json:"location_id"`
}

type Track struct {
	Name     string  `json:"name"`
	Length   float64 `json:"length"`
	Location Location
}

func (t Track) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Location)
}

var UnknownTrack = Track{Name: "Unknown track", Length: 0.0, Location: UnknownLocation}

type CarClass struct {
	Name       string `json:"name"`
	Drivetrain string `json:"drivetrain"`
}

var UnknownCarClass = CarClass{Name: "???", Drivetrain: "???"}

type carJson struct {
	Name    string `json:"name"`
	ClassId int    `json:"car_class_id"`
}

type Car struct {
	Name  string
	Class CarClass
}

func (c Car) String() string {
	return c.Name
}

var UnknownCar = Car{Name: "Unknown car", Class: UnknownCarClass}

type Dicts struct {
	Tracks map[int]Track
	Cars   map[int]Car
}

func (dicts Dicts) GetTrackById(id int) Track {
	return u.GetOrElse(&dicts.Tracks, id, UnknownTrack)
}

func (dicts Dicts) GetCarById(id int) Car {
	return u.GetOrElse(&dicts.Cars, id, UnknownCar)
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
			Name:     track.Name,
			Length:   track.Length,
			Location: u.GetOrPanic(&locations, track.LocationId),
		}
	}

	carsRaw := map[int]carJson{}
	loadJsonOrPanic("cars.json", &carsRaw)
	cars := map[int]Car{}
	for carId, car := range carsRaw {
		cars[carId] = Car{
			Name:  car.Name,
			Class: u.GetOrPanic(&carClasses, car.ClassId),
		}
	}

	return Dicts{
		Tracks: tracks,
		Cars:   cars,
	}

}
