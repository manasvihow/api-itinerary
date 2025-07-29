package models

import "html/template"

type Traveller struct {
	Name               string `json:"name"`
	NumberOfTravellers int    `json:"number_of_travellers"`
}

type Trip struct {
	DepartureFrom string `json:"departure_from"`
	DepartureDate string `json:"departure_date"`
	Destination   string `json:"destination"`
	ArrivalDate   string `json:"arrival_date"`
	ReturnDate    string `json:"return_date"`
}

type Flight struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Airline   string `json:"airline"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Price     int    `json:"price"`
}

type Hotel struct {
	Name          string `json:"name"`
	CheckIn       string `json:"check_in"`
	CheckOut      string `json:"check_out"`
	PricePerNight int    `json:"price_per_night"`
}

type Activity struct {
	Date        string `json:"date"`
	Time        string `json:"time"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DailyActivity struct {
	DayNumber int
	Date      string
	ImageURL  string
	Morning   []Activity
	Afternoon []Activity
	Evening   []Activity
}

type Installment struct {
	Label   string
	Amount  int
	DueDate string
}

type Visa struct {
	Type           string `json:"type"`
	Validity       string `json:"validity"`
	ProcessingDate string `json:"processing_date"`
}

type ItineraryRequest struct {
	Traveller            Traveller    `json:"traveller"`
	Trip                 Trip         `json:"trip"`
	Flights              []Flight     `json:"flights"`
	Hotels               []Hotel      `json:"hotels"`
	Activities           []Activity   `json:"activities"`
	NumberOfInstallments int          `json:"number_of_installments"`
	Visa                 Visa         `json:"visa"`
	CSS                  template.CSS `json:"-"`
	LogoDataURI          string       `json:"-"`
}
