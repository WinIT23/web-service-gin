package db

import (
	"fmt"
	"sync"

	"github.com/WinIT23/web-service-gin/models"
)

type DataBase struct {
	Albums []models.Album
}

var singleInstance *DataBase
var lock = &sync.Mutex{}

func GetDatabase() *DataBase {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &DataBase{}
			singleInstance.init()
			fmt.Println("[GetDatabase] Setting up database instance")
		} else {
			fmt.Println("[GetDatabase] Instance already available")
		}
	} else {
		fmt.Println("[GetDatabase] Instance already available")
	}
	return singleInstance
}

func (d *DataBase) init() {
	fmt.Println("[init] initializing database")
	d.Albums = []models.Album{
		{ID: "1", Title: "Some Title - 1", Artist: "Some Artist - 1", Price: 9.69},
		{ID: "2", Title: "Some Title - 2", Artist: "Some Artist - 2", Price: 6.09},
		{ID: "3", Title: "Some Title - 3", Artist: "Some Artist - 3", Price: 9.60},
		{ID: "4", Title: "Some Title - 4", Artist: "Some Artist - 4", Price: 6.69},
		{ID: "5", Title: "Some Title - 5", Artist: "Some Artist - 5", Price: 0.69},
	}
}
