module Part2
import Data.String
import Debug.Trace

forward : (Int, Int, Int) -> Int -> (Int, Int, Int)
forward (a, b, c) d = (a+d, b + (d * c), c)


down : (Int, Int, Int) -> Int -> (Int, Int, Int)
down (a, b, c) d = (a, b, c + d)


up : (Int, Int, Int) -> Int -> (Int, Int, Int)
up (a, b, c) d = (a, b, c - d)


parseOp : String -> ((Int, Int, Int) -> Int -> (Int, Int, Int))
parseOp s = if s == "forward" then forward else if s == "down" then down else up


foldOp : (Int, Int, Int) -> (String, String) -> (Int, Int, Int)
foldOp (a, b) (s, t) = (parseOp s) (a, b) (fromMaybe 0 (parseInteger (trim t)))

splitInput : String -> (String, String)
splitInput s = break (== ' ') s

export
splitInputs : List(String) -> List(String, String)
splitInputs l = map splitInput l

export
solve : List(String) -> (Int, Int, Int)
solve l = foldl foldOp (0, 0, 0) (splitInputs (filter (/= "") l))
