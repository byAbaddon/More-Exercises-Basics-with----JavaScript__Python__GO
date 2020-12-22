function streamOfLetters([...arg]) {
   arg.pop()
   let wordCollection = ''
   let word = ''
   let [c, o, n] = [false, false, false]

   for (let i = 0; i < arg.length; i++) {
      let chars = arg[i]
      
      if (/[^A-Za-z]/.test(chars)) {
         continue
      }

      if ((chars == 'c' && c != true) || (chars == 'o' && o != true) || (chars == 'n' && n != true)) {

         if (chars === 'c') c = true
         if (chars === 'o') o = true
         if (chars === 'n') n = true
      } else {
         word += chars
      }

      if (c && o && n) {
         wordCollection += word + ' '
         c = false
         o = false
         n = false
         word = ''
      }
   }

   return wordCollection
}

//console.log((streamOfLetters(['H', 'n', 'e', 'l', 'l', 'o', 'o', 'c', 't', 'c', 'h', 'o', 'e', 'r', 'e', 'n', 'e', 'End'])))
//console.log((streamOfLetters(['%','!','c','^','B','`','o','%','o','o','M',')','{','n','\\','A','D','End']))) //BooM