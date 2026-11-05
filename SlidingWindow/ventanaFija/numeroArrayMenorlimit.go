package main

func arraysLessLimit(arr[] int , k,limit int)int{

	sumaVentana:=0
	ventanasQueCumplen:=0

	//calcular primera ventana
	for i:=0 ; i<k;i++{
		sumaVentana=sumaVentana+arr[i]
	}
	if sumaVentana< limit{
		ventanasQueCumplen++
	}
	//calculamos las siguientes ventanas

	for rigth:=k;rigth<len(arr);rigth++{
		sumaVentana=sumaVentana-arr[rigth-k]+arr[rigth]
		if sumaVentana<limit{
			ventanasQueCumplen++
		}
	}
return ventanasQueCumplen
}