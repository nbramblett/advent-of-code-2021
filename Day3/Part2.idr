module Part2

import Day3.Part1

matchingIndices : List Char -> Nat -> Bool -> (String -> Bool)
matchingIndices lc n most = \s => ((index' n lc) == Just (strIndex s (toIntNat n))) == most

export
partial solve2 : List String -> List Char -> Nat -> Bool -> Maybe String
solve2 ls lc n most = if (length ls) == 1 then index' 0 ls else solve2 nls (solve nls) (n+1) most where
    nls : List String
    nls = (filter (matchingIndices lc n most) ls)
    