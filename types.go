package main

type Movie struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
