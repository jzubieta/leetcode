package main

func maxSumK(nums []int, k int) int {
    // 1. suma de la primera ventana (los primeros k elementos)
	suma:=0
	maxSuma:=0
	for i:=0;i<k;i++{
		maxSuma=maxSuma+nums[i]
	}

    // 2. el máximo empieza con ese valor
	maxSuma=suma
    // 3. for right := k; right < len(nums); right++ {
	for right := k; right < len(nums); right++ {
    //        actualiza la suma: entra ..., sale ...
			suma= suma+nums[k]-nums[k-1]
	//       actualiza el máximo
			maxSuma=max(maxSuma,suma)
    //    
	}
    // 4. return el máximo
	return maxSuma
}

