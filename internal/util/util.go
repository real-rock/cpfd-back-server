package util

import (
	log "cpfd-back/internal/log"
	"os"

	"github.com/gocarina/gocsv"
)

func WriteTempCSV(obj interface{}, seed string) (string, error) {
	f, err := os.CreateTemp("", seed)
	if err != nil {
		log.Logger.Errorf("failed to create temp %s file: %v", seed, err.Error())
		return "", err
	}
	defer f.Close()
	if err := gocsv.MarshalCSV(obj, gocsv.DefaultCSVWriter(f)); err != nil {
		log.Logger.Errorf("failed to write particles to file `%s: %s\n", f.Name(), err.Error())
		return "", err
	}
	return f.Name(), nil
}
