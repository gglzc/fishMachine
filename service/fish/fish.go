package services

import (
	"errors"
	"sync"
)

type FishService struct{
	mu sync.Mutex
}
// generate fish group
func (fs *FishService)GenerateFishGroup(FishConfig interface{})error{
	if FishConfig!=nil{
		return errors.New("fish Config Setting wrong")
	}
	fs.mu.Lock()
	defer fs.mu.Unlock()
	// generate 
	// nums :=20

	return nil
}