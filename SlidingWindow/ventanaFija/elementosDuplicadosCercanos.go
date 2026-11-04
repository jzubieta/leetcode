package main

import "math"

/*
Problema: Dado un arreglo de enteros nums y un entero k,
 determina si existen dos índices distintos i y j
 en el arreglo tales que nums[i] == nums[j]
  y la distancia absoluta entre ellos sea
  como máximo k (el tamaño de la ventana es k + 1).
• Entrada: nums = [1, 2, 3, 1], k = 3
• Salida: true
• Explicación: nums[0] == nums[3] y 3 - 0 <= 3.


*/
func containsNearDuplicated(arr[] int ,k int)bool{

	m:=make(map[int]int)
	for i:=0;i<len(arr);i++{
		if v,ok:=m[arr[i]];ok{
			if i-v<=k{
				return true
			}else{
				m[arr[i]]=i
			}
		}else{
			m[arr[i]]=i
		}
	


	}

return false

}