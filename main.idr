module Main

import Day1.Day1
import Helpers.Helpers

main : IO ()
main = do
  list <- (Helpers.parseFile "Day1/input.txt")
  printLn (solve list)
