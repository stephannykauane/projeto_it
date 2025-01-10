package calculator

import "fmt"

func (a SatBases) calculo() float64 {

	
	s := a.Calcio + a.Magnesio + a.Potassio
	v1 := (s / a.Ctc) * 100

	fmt.Println("valor saturação atual: ", v1)
	fmt.Println("valor s atual: ", s)
	fmt.Println("valor saturação desejada: ", a.SatD)

	NC := ((a.Ctc * (a.SatD - v1)) / 10) * float64(a.Prnt)
    fmt.Println("valor ctc: ", a.Ctc)
	fmt.Println("valor Magnesio: ", a.Magnesio)
	fmt.Println("valor calcio: ", a.Calcio)
	fmt.Println("valor Potassio: ", a.Potassio)
	fmt.Println("valor prnt: ", a.Prnt)



	return NC

}

func (b Aluminio) calculo() float64 {
	var NC float64

	if b.Ctc > 4.0 || b.Argila > 15 || (b.Calcio+b.Magnesio) < 2.0 {

		NC = ((2 * b.Aluminio) + 2 - (b.Calcio + b.Magnesio)) * float64(b.Prnt)
		fmt.Println("esse foi utilizado!")

	} else {
		NC = (2 * b.Aluminio) * float64(b.Prnt)
		fmt.Println("esse AQUI foi utilizado!")
	}

	fmt.Println("valor ctc: ", b.Ctc)
	fmt.Println("valor argila: ", b.Argila)
	fmt.Println("valor calcio: ", b.Calcio)
	fmt.Println("valor magnesio: ", b.Magnesio)
	fmt.Println("valor aluminio: ", b.Aluminio)
	fmt.Println("valor prnt: ", b.Prnt)
	fmt.Println("nc: ", NC)

	return NC
}

type calculavel interface {
	calculo() float64
}

func gerarCalculo(c calculavel) float64 {
	return c.calculo()
}
