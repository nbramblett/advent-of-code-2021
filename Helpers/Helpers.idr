module Helpers

export
parseFile : String -> IO (List String)
parseFile path = do file <- readFile path
                    case file of
                      Right content => pure (split (== '\n') content)
                      Left err => pure Nil
