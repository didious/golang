package go-area

import "math"

//Pi Variavel com letra maiuscula é publica
const Pi = 3.1416

//Circulo func publica letra maiuscula
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Retangulo tbm é publica
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//N é publica por causa do underline
func _TriangEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}
