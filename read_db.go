package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// data\\manana.dat

type coordenadas struct {
	x int
	y int

}

type request struct{
	xi int
	yi int
	xf int
	yf int
	t  int

}

func anadir(x int, y int) coordenadas {
	p :=coordenadas{x:x,y:y}
	return p
}

func anadirRequest(xi int, yi int,xf int,yf int, t int) request {
	p :=request{xi:xi,yi:yi,xf:xf,yf:yf,t:t}
	return p
}

func readData(file *os.File, reader *bufio.Reader)[]coordenadas{
	var arrayCoordenadas []coordenadas
	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		var_temp:=string(line)

		variable:=strings.Split(var_temp," ") //Separo la linea
		x_temp,err:=strconv.Atoi(variable[0])
		if err != nil {
			panic(err)
		}
		parsing:=strings.Split(variable[1],"\n")
		y_temp,err:=strconv.Atoi(parsing[0])

		//fmt.Println(x_temp)
		//fmt.Println(y_temp)

		arrayCoordenadas = append(
			arrayCoordenadas,
			anadir(x_temp,y_temp),
		)

	}

	return arrayCoordenadas
}


func helpParsing(cadena string)(int,int){
	ayuda:=strings.FieldsFunc(cadena, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})

	variable:=strings.Join(ayuda," ")

	ultimo:=strings.Split(variable," ") //Separo la linea

	x_temp,err:=strconv.Atoi(ultimo[0])
	if err != nil {
		panic(err)
	}
	parsing:=strings.Split(ultimo[1],"\n")
	y_temp,err:=strconv.Atoi(parsing[0])

	return x_temp,y_temp
}


func auxiliar(numero string)string{
	ayuda:=strings.FieldsFunc(numero, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	variable:=strings.Join(ayuda," ")

	ultimo:=string(variable)

	return ultimo

}


func readRequest(file *os.File, reader *bufio.Reader)[]request{
	var arrayRequest []request
	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}

		var_temp:=string(line)

		aloja:=strings.FieldsFunc(var_temp, func(r rune) bool {
			if r == '-' {
				return true
			}
			return false
		})

		//fmt.Println(var_temp)
		//Cada lisea representa un espaceado,

		//Ejemplo 492 720 - 521 353 - 9
		uno:=aloja[0]//492 720
		dos:=aloja[1]//521 353
		//fmt.Println(uno)
		//fmt.Println(dos)

		//aloja[2] // 9
		xi,yi:=helpParsing(uno)
		xf,yf:=helpParsing(dos)
		t,err:=strconv.Atoi(auxiliar(aloja[2]))
		if err!=nil{
			panic(err)
		}


		arrayRequest = append(
			arrayRequest,
			anadirRequest(xi,yi,xf,yf,t),
		)

		//fmt.Println(x_temp)
		//fmt.Println(y_temp)
	}

	return arrayRequest

}

func readDataRequest(path string)[]request{
	content, _ := os.Open(path)
	reader := bufio.NewReader(content)
	lines :=readRequest(content, reader)

	return lines

}


func readDataTime(path string)[]coordenadas{
	//This is the function you are going to use to read the data at night, morning, afternoon, etc...
	content, _ := os.Open(path)
	reader := bufio.NewReader(content)
	lines :=readData(content, reader)

	return lines
}