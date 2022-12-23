package attraction

type Detail struct {
	/*
		Tourist attraction detail data model.
		This object is passed when the user request
		detailed information about the attraction
		they're reviewing via the iOS app. Not all of
		the information here is used when displaying the
		attractions in a list, so to reduce the payload
		size, separate object types are used for the same
		tourist attraction.
	*/

	// The ID can be a simple int, which can start from 1.
	Id int32 `json:"id,omitempty"`

	// The name of the tourist attraction, example "Cutty Sark".
	Name string `json:"name"`

	// The city it is located, example "London"
	City string `json:"city"`

	// The country it is located in, example "United Kingdom".
	Country string `json:"country"`

	// The hours of the day it is open, example "09:00 - 18:00"
	// Assumes the same opening hours for all days of the week.
	HoursOpen string `json:"hoursOpen"`

	// The days it's opne, example "weekdays"/"weekends"/"all".
	DaysOpen string `json:"daysOpen"`

	// Apple Maps rating, if applicable, example "4.5".
	Rating float32 `json:"rating"`

	// 2D coordinate latitude, example "51.48286"
	Latitude float64 `json:"latitude"`

	// 2D coordinate longitutde, example "-0.00145"
	Longitude float64 `json:"longitude"`

	// Image name to be used for testing purposes.
	// The image will be stored on the client.
	ImageName string `json:"imageName"`
}
