package repositories

import "api/models"

// Database
var albums = []models.Album{ //Slice of albums with their attributs to simulate a database
	{ID: "1", Title: "Metaphorical Music", Artist: "Nujabes", Price: 56.99},
	{ID: "2", Title: "Short n Sweet", Artist: "Sabrina Carpenter", Price: 22.99},
	{ID: "3", Title: "MM FOOD", Artist: "MF DOOM", Price: 88.99},
}

func GetAbums() []models.Album { //Just return all the albums
	return albums
}

func AddAlbum(album models.Album) { //Adds a new album passing an album parameter
	albums = append(albums, album) //Appends the album
}
