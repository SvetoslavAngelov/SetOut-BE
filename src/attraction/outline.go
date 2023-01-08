package attraction

type Outline struct {
	/*
		Tourist attraction outline data model.
		This object is used to show the list
		of tourist attractions near the user and
		only key information is displayed in each
		card presented to the user.
	*/

	// The ID can be a simple int, which can start from 1.
	Id int32 `json:"id,omitempty"`

	// The name of the tourist attraction, example "Cutty Sark".
	Name string `json:"name"`

	// Based on the HoursOpen & DaysOpen information from the
	// Detail object, determine whether the attraction is open.
	// This can be determined when the client requests this object.
	IsOpen bool `json:"isOpen"`

	// The distance between the client's current location and
	// this tourist attraction. This can be determined when the
	// client requests this object.
	Distance float32 `json:"distance"`

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
