package main

import "fmt"
import  "regexp"
func main() {
   var wordCollection, word string
   var c, o, n bool
   
   for {
      var chars string
      fmt.Scan(&chars)
      if chars == "End"{
         break
      }

      var digitCheck = regexp.MustCompile(`[^A-Za-z]+$`)

      if (digitCheck.MatchString(chars)) {
         continue
      }


      if (chars == "c" && c != true) || (chars == "o" && o != true) || (chars == "n" && n != true) {

         if chars == "c"{
             c = true
         }
         if chars == "o"{
             o = true
         }
         if chars == "n"{
            n = true
         } 
      } else {
         word += chars
      }

      if (c && o && n) {
         wordCollection += word + " "
         c = false
         o = false
         n = false
         word = ""
      }
   }

   fmt.Println(wordCollection)
}

/*
name   :03. Stream Of Letters
input  :H n e l l o o c t c h o e r e n e End
output :Hello there 
*/
