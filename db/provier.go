package db

import (
	"encoding/csv"
	"io"
	"os"
)

func forEachRecord(fileName string, transformer func(record []string) interface{}) []interface{} {
	var result []interface{}
	file, _ := os.Open(fileName)
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Read()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		result = append(result, transformer(record))
	}

	return result
}

func Locations() []interface{} {
	return forEachRecord("_data/GeoLite2-Country-Locations.csv", func(record []string) interface{} {
		return Location{
			GeonameId:          record[0],
			ContinentCode:      record[1],
			ContinentName:      record[2],
			CountryISOCode:     record[3],
			CountryName:        record[4],
			SubDivisionISOCode: record[5],
			SubDivisionName:    record[6],
			CityName:           record[7],
			MetroCode:          record[8],
			TimeZone:           record[9],
		}
		return record
	})
}

func Blocks() []interface{} {
	return forEachRecord("_data/GeoLite2-Country-Blocks.csv", func(record []string) interface{} {
		return Block{
			NetworkStartIp:              record[0],
			NetworkPrefixLength:         record[1],
			GeonameId:                   record[2],
			RegisteredCountryGeoNameId:  record[3],
			RepresentedCountryGeoNameId: record[4],
			PostalCode:                  record[5],
			Latitude:                    record[6],
			Longitude:                   record[7],
			IsAnonymousProxy:            record[8],
			IsSatelliteProvider:         record[9],
		}
	})
}
