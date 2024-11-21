package kubernetes

import (
	"log"
	"math/rand"
	"os"
)

type PersistantVolume struct {
	path string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreatePersistantVolume() (*PersistantVolume, error) {
	log.Println("Creating Persistant Volume...")
	path := "/tmp/" + randSeq(10)
	os.Mkdir(path, 0750)
	return &PersistantVolume{
		path: path,
	}, nil
}

func (pv *PersistantVolume) GetPath() string {
	return pv.path
}
