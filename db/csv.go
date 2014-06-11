package db

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

type CsvDataProvider struct {
	locations map[string]Location
	blocks    []Block
}

func newBlockFromRecord(record []string) Block {
	prefixLength, err := strconv.Atoi(record[1])
	if err != nil {
		panic(err)
	}
	return Block{
		NetworkStartIp:              net.ParseIP(record[0]),
		NetworkPrefixLength:         prefixLength,
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

func newCsvReader(fileName string) *csv.Reader {
	file, _ := os.Open(fileName)
	defer file.Close()

	reader := csv.NewReader(file)
	return reader
}

func newLocationFromRecord(record []string) Location {
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

func loadBlocks(fileName string) []Block {
	var result []Block
	file, _ := os.Open(fileName)
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Read() // discard the header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		result = append(result, newBlockFromRecord(record))
	}

	return result
}

func loadLocations(fileName string) map[string]Location {
	file, _ := os.Open(fileName)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // discard the header

	result := map[string]Location{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		location := newLocationFromRecord(record)
		result[location.GeonameId] = location
	}

	return result
}

func NewCsvDataProvider(dataFolderName string) *CsvDataProvider {
	result := CsvDataProvider{}
	blockFilePath := fmt.Sprintf("%s/GeoLite2-Country-Blocks.csv", dataFolderName)
	result.blocks = loadBlocks(blockFilePath)

	locationFilePath := fmt.Sprintf("%s/GeoLite2-Country-Locations.csv", dataFolderName)
	result.locations = loadLocations(locationFilePath)
	return &result
}

func (this *CsvDataProvider) GetLocationByGeonameId(geonameid string) Location {
	return this.locations[geonameid]
}

func (this *CsvDataProvider) GetBlockByIP(ipaddr net.IP) Block {
	// naive implementation for now
	target := ipaddr.To16()

	for _, block := range this.blocks {
		if bytes.Compare(block.NetworkStartIp, target) < 0 {
			networkEndIp := block.NetworkEndIp()
			if bytes.Compare(target, networkEndIp) < 0 {
				return block
			}
		}
	}
	return Block{}
}
