package calculator

func (a SatBases) calculo() float64 {

	
	v2 := a.SatD 
	s := a.Calcio + a.Magnesio + a.Potassio 
    v1 := (s/a.Ctc) * 100
    
	NC := (a.Ctc * (v2 - v1))/ 10 * float64(a.Prnt)

	return NC

}

func (b Aluminio) calculo() float64 {
    var NC float64

	if (b.Ctc > 4.0 || b.Argila > 15 || (b.Calcio + b.Magnesio) < 2.0){

		NC = ((2 * b.Aluminio) + 2 - (b.Calcio + b.Magnesio) * float64(b.Prnt))

	}else{
		NC = (2 * b.Aluminio) * float64(b.Prnt)
	}
	
	return NC 
}

type calculavel interface {

	calculo() float64
}

func gerarCalculo (c calculavel) float64 {

	return c.calculo()
}


