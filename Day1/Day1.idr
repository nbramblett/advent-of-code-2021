module Day1
import Data.String


foldCount : (Int, Int) -> Int -> (Int, Int)
foldCount (a, b) c = (c, (if a < c then b + 1 else b))

foldCount2 : (Int, Int, Int, Int) -> Int -> (Int, Int, Int, Int)
foldCount2 (a, b, c, d) e = (b, c, e, (if (a+b+c) < (b+c+e) then d + 1 else d))

export
solve : List String -> (Int, Int, Int, Int)
solve content = foldl foldCount2 (1000000, 1000000, 100000, 0) (mapMaybe parseInteger content)

