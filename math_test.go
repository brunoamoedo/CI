package main

import "testing"

func TestSoma(t *testing.T) 
{
	total:= Soma(15,15)
	if total != 30{
		t.Errorf("Resultado da soma e invalido: Resultado %d. Esoerado : %d",total,30)
	}
}
TestSoma()