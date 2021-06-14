package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var bdongs []BDONG

type BDONG struct {
	Numero        int    `json:"numero"`
	Institucion   string `json:"institucion"`
	Departamento  string `json:"departamento"`
	Provincia     string `json:"provincia"`
	Distrito      string `json:"distrito"`
	Representante string `json:"representante"`
	Sector        string `json:"sector"`
}

func lineToStruc(lines [][]string) {
	// Recorre líneas y conviértete en objeto
	for _, line := range lines {
		Numero, _ := strconv.Atoi(strings.TrimSpace(line[0]))

		bdongs = append(bdongs, BDONG{
			Numero:        Numero,
			Institucion:   strings.TrimSpace(line[1]),
			Departamento:  strings.TrimSpace(line[2]),
			Provincia:     strings.TrimSpace(line[3]),
			Distrito:      strings.TrimSpace(line[4]),
			Representante: strings.TrimSpace(line[5]),
			Sector:        strings.TrimSpace(line[6]),
		})
	}
}

func readFileUrl(filePathUrl string) ([][]string, error) {
	// Abrir archivo CSV
	f, err := http.Get(filePathUrl)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Body.Close()

	// Leer archivo en una variable
	lines, err := csv.NewReader(f.Body).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func main() {
	//filePathUrl := "dataset/Base-de-Datos-de-las-ONGD-I-Trimestre-2018_0.csv"
	filePathUrl := "https://raw.githubusercontent.com/sigiandre/TA2-Programacion-Concurrente-y-Distribuida-Backend/master/dataset/Base-de-Datos-de-las-ONGD-I-Trimestre-2018_0.csv"
	lines, err := readFileUrl(filePathUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Leyo archivos")
	lineToStruc(lines)
	fmt.Println("Parseo Archivos")
}
