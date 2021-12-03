module Part1

import Data.String

mode : (Nat, Nat) -> Char
mode (a, b) = if a > b then '0' else '1'

indexedCounter : Nat -> List (List Char) -> (Nat, Nat)
indexedCounter n llc = foldl foldOp (0, 0) llc where
    foldOp : (Nat, Nat) -> List Char -> (Nat, Nat)
    foldOp (x, y) l = if (index' n l) == Just '0' then (x + 1, y) else (x, y + 1)

countModes : List (List Char) -> List (Nat, Nat)
countModes llc = map doer (natRange 12) where
    doer : Nat -> (Nat, Nat)
    doer n = indexedCounter n llc

export
solve : List String -> List (Char)
solve a = map mode (countModes (map unpack a))