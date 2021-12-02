module Main

-- To solve a part for a day, import Day#.Part#
import Day2.Part2
import Helpers.Helpers

main : IO ()
main = do
  list <- (parseFile "Day2/input.txt")
  printLn (solve list)
