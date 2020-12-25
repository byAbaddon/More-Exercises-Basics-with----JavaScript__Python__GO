function primePairs(arg) {
   let [fStart, sStart, fEnd, sEnd] = [...arg.map(Number)]
   fEnd += fStart
   sEnd += sStart
   const primeList = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101]

   for (let i = fStart; i <= fEnd; i++) {
      for (let j = sStart; j <= sEnd; j++) {
         if (primeList.includes(i) && primeList.includes(j))
            console.log('' + i + j)
      }
   }
}

// primePairs([10, 20, 5, 5])
// primePairs([10, 30, 9, 6])