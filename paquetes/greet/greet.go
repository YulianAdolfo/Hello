package greet

var name string // función a nivel de paquete . No se puede utilizar los := se hace con var

func InEnglish(n string) string {
	return "Hello " + n
}
func InItalian(n string) string {
	return "Cia " + n
}

// cuando la primer letra de una variable o función es en Mayúscula, Go las exporta
// si están en Minúscula, Go NO las exporta
