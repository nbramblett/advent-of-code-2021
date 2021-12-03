module Main

-- To solve a part for a day, import Day#.Part#
import Day3.Part1
import Day3.Part2
import Helpers.Helpers


main : IO ()
main = do
  list <- (parseFile "Day3/input.txt")
  printLn (solve list)
  printLn (solve2 list (solve list) 0 True)
  printLn (solve2 list (solve list) 0 False)
  
