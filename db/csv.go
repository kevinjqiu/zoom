package db

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvDataProvider struct {
	locations []interface{}
	blocks    []interface{}
}

func newBlockFromRecord(record []string) interface{} {
	return Block{
		NetworkStartIp:              record[0],
		NetworkPrefixLength:         record[1],
		GeonameId:                   record[2],
		RegisteredCountryGeoNameId:  record[3],
		RepresentedCountryGeoNameId: record[4],
		PostalCode:                  record[5],
		Latitude:                    record[6],
		Longitude:                   record[7],
		IsAnonymousProxy:            record[8] != "0",
		IsSatelliteProvider:         record[9] != "0",
	}
}

func newLocationFromRecord(record []string) interface{} {
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
}

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
		transformedResult := transformer(record)
		result = append(result, transformedResult)
	}

	return result
}

func NewCsvDataProvider(dataFolderName string) *CsvDataProvider {
	result := CsvDataProvider{}
	blockFilePath := fmt.Sprintf("%s/GeoLite2-Country-Blocks.csv", dataFolderName)
	result.blocks = forEachRecord(blockFilePath, newBlockFromRecord).([]Block)

	locationFilePath := fmt.Sprintf("%s/GeoLite2-Country-Locations.csv", dataFolderName)
	result.locations = forEachRecord(locationFilePath, newLocationFromRecord)
	return &result
}
