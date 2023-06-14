function streamOfLetters(arg) { 
  let [word, collection, code] = ['', '', ['c', 'n', 'o']]
  for (const ch of arg) {
    if(!ch.match(/[a-zA-Z]/)) continue
    if (code.includes(ch)) code = code.filter(x => x != ch) 
    else word += ch
     
    if (!code.length) {
      code = ['c', 'n', 'o']
      collection += word
      word = ' '
    }
  }
  return collection
}

//--------------------------------------------(2)--------------------------------

function streamOfLetters(arg) {
  let wordCollection = word = ''
  let [c, o, n] = [false, false, false]

  for (let i = 0; i < arg.length; i++) {
     let char = arg[i]
     
     if (/[^A-Za-z]/.test(char)) continue
  
     if (char == 'c' && !c || char == 'o' && !o || char == 'n' && !n) {
        if (char == 'c') c = true
        if (char == 'o') o = true
        if (char == 'n') n = true
     } else {
        word += char
     }

     if (c && o && n) {
        wordCollection += word + ' '
        c = o = n = false
        word = ''
     }
  }

  return wordCollection
}

//console.log((streamOfLetters(['H', 'n', 'e', 'l', 'l', 'o', 'o', 'c', 't', 'c', 'h', 'o', 'e', 'r', 'e', 'n', 'e', 'End'])))
//console.log((streamOfLetters(['%','!','c','^','B','`','o','%','o','o','M',')','{','n','\\','A','D','End']))) //BooM
