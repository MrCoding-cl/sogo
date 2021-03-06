package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func anadir(x int, y int) UberFormat {
	p := UberFormat{X: x, Y: y}
	return p
}

func anadirRequest(xi int, yi int, xf int, yf int, t int) RequestFormat {
	p := RequestFormat{Xi: xi, Yi: yi, Xf: xf, Yf: yf, T: t}
	return p
}

func readData(file *os.File, reader *bufio.Reader) []UberFormat {
	var arrayCoordenadas []UberFormat
	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		var_temp := string(line)
		variable := strings.Split(var_temp, " ") //Separo la linea
		x_temp, err := strconv.Atoi(variable[0])
		if err != nil {
			panic(err)
		}
		parsing := strings.Split(variable[1], "\n")
		y_temp, err := strconv.Atoi(parsing[0])
		arrayCoordenadas = append(
			arrayCoordenadas,
			anadir(x_temp, y_temp),
		)
	}
	return arrayCoordenadas
}

func helpParsing(cadena string) (int, int) {
	ayuda := strings.FieldsFunc(cadena, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	variable := strings.Join(ayuda, " ")
	ultimo := strings.Split(variable, " ") //Separo la linea
	x_temp, err := strconv.Atoi(ultimo[0])
	if err != nil {
		panic(err)
	}
	parsing := strings.Split(ultimo[1], "\n")
	y_temp, err := strconv.Atoi(parsing[0])
	return x_temp, y_temp
}

func auxiliar(numero string) string {
	ayuda := strings.FieldsFunc(numero, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	variable := strings.Join(ayuda, " ")
	ultimo := string(variable)
	return ultimo
}

func readRequest(file *os.File, reader *bufio.Reader) []RequestFormat {
	var arrayRequest []RequestFormat
	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		var_temp := string(line)
		aloja := strings.FieldsFunc(var_temp, func(r rune) bool {
			if r == '-' {
				return true
			}
			return false
		})
		uno := aloja[0] //492 720
		dos := aloja[1] //521 353
		xi, yi := helpParsing(uno)
		xf, yf := helpParsing(dos)
		t, err := strconv.Atoi(auxiliar(aloja[2]))
		if err != nil {
			panic(err)
		}
		arrayRequest = append(
			arrayRequest,
			anadirRequest(xi, yi, xf, yf, t),
		)
	}

	return arrayRequest

}

func readDataRequest(path string) []RequestFormat {
	content, _ := os.Open(path)
	reader := bufio.NewReader(content)
	lines := readRequest(content, reader)
	return lines

}

func readDataTime(path string) []UberFormat {
	//This is the function you are going to use to read the data at night, mo0rning, afternoon, etc...
	content, _ := os.Open(path)
	reader := bufio.NewReader(content)
	lines := readData(content, reader)

	return lines
}
func addClientsToWorld(world *world, path string) {
	request := readDataRequest(path)
	for id, r := range request {
		client := createUberPassenger(id, r.Xi, r.Yi, r.Xf, r.Yf, r.T, world)
		world.addClient(world, &client)
	}
}
func addUbersToWorld(world *world, path string) {
	request := readDataTime(path)
	for id, r := range request {
		uber := createUber(id, r.X, r.Y, world)
		world.ubers = append(world.ubers, &uber)
	}
}
