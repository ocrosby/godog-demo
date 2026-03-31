package models

// Geo holds the geographic coordinates of a user's address as returned by the
// JSONPlaceholder API. Both Lat and Lng are represented as strings to preserve
// the exact precision supplied by the API.
type Geo struct {
	// Lat is the latitude as a decimal string, e.g. "-37.3159".
	Lat string `json:"lat"`

	// Lng is the longitude as a decimal string, e.g. "81.1496".
	Lng string `json:"lng"`
}

// Address represents a user's postal address as returned by the JSONPlaceholder
// /users endpoint.
type Address struct {
	// Street is the street name and number.
	Street string `json:"street"`

	// Suite is the apartment, suite, or unit identifier.
	Suite string `json:"suite"`

	// City is the city name.
	City string `json:"city"`

	// Zipcode is the postal code.
	Zipcode string `json:"zipcode"`

	// Geo holds the geographic coordinates of this address.
	Geo Geo `json:"geo"`
}

// Company represents the employer of a user as returned by the JSONPlaceholder
// /users endpoint.
type Company struct {
	// Name is the company's name.
	Name string `json:"name"`

	// CatchPhrase is the company's slogan or tagline.
	CatchPhrase string `json:"catchPhrase"`

	// Bs is a freeform business-speak descriptor used by JSONPlaceholder as
	// placeholder content.
	Bs string `json:"bs"`
}

// User represents a user account as returned by the JSONPlaceholder /users
// endpoint, including their contact details, address, and employer.
type User struct {
	// ID is the unique identifier of the user.
	ID int `json:"id"`

	// Name is the user's full name.
	Name string `json:"name"`

	// Username is the user's login handle.
	Username string `json:"username"`

	// Email is the user's email address.
	Email string `json:"email"`

	// Address is the user's postal address including geographic coordinates.
	Address Address `json:"address"`

	// Phone is the user's phone number, which may include extensions.
	Phone string `json:"phone"`

	// Website is the URL of the user's personal website.
	Website string `json:"website"`

	// Company holds information about the user's employer.
	Company Company `json:"company"`
}
