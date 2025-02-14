package testdata

import "github.com/SvetoslavAngelov/tourplan-app/src/attraction"

var Attractions = []attraction.Outline{
	{Id: 1,
		Name:      "Greenwich Observatory",
		IsOpen:    true,
		Distance:  2.64,
		Rating:    4.7,
		Latitude:  -0.000536,
		Longitude: 51.476833,
		ImageName: "default"},
	{Id: 2,
		Name:      "Cutty Sark",
		IsOpen:    true,
		Distance:  1.85,
		Rating:    4.5,
		Latitude:  -0.009586,
		Longitude: 51.482880,
		ImageName: "default"},
	{Id: 3,
		Name:      "Tower Bridge",
		IsOpen:    false,
		Distance:  3.45,
		Rating:    4.9,
		Latitude:  -0.075402,
		Longitude: 51.505507,
		ImageName: "default"},
	{Id: 4,
		Name:      "O2 Millenium Dome",
		IsOpen:    true,
		Distance:  2.35,
		Rating:    4.1,
		Latitude:  0.003182,
		Longitude: 51.502937,
		ImageName: "default"},
	{Id: 5,
		Name:      "Buckingham Palace",
		IsOpen:    true,
		Distance:  3.81,
		Rating:    4.8,
		Latitude:  -0.1440787,
		Longitude: 51.501364,
		ImageName: "default"},
}
