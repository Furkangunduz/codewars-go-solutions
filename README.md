# Codewars Go Solutions

This repository contains my solutions to various [Codewars](https://www.codewars.com/) kata (coding challenges) implemented in Go.

## Structure

Each file contains a solution to a specific Codewars kata, with the following format:

- The file is named after the kata (e.g., `duplicate-encoder.go`)
- Each file includes a comment with a link to the original kata
- The solution is implemented as a Go function within the `solutions` package

## Solutions

| Kata Name                                     | Difficulty | Description                                                                                                       | Link                                                                |
| --------------------------------------------- | ---------- | ----------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| [Duplicate Encoder](duplicate-encoder.go)     | 6 kyu      | Convert a string where each character is replaced with "(" if it appears once or ")" if it appears multiple times | [Kata Link](https://www.codewars.com/kata/54b42f9314d9229fd6000d9c) |
| [Next Smaller Number](next-smaller-number.go) | 4 kyu      | Find the next smaller number with the same digits                                                                 | [Kata Link](https://www.codewars.com/kata/5659c6d896bc135c4c00021e) |
| [Range Extraction](range-extraction.go)       | 4 kyu      | Format a list of integers as a comma separated list with ranges                                                   | [Kata Link](https://www.codewars.com/kata/51ba717bb08c1cd60f00002f) |
| [RGB to Hex](rgb-to-hex.go)                   | 5 kyu      | Convert RGB values to hexadecimal representation                                                                  | [Kata Link](https://www.codewars.com/kata/513e08acc600c94f01000001) |

## Running the Solutions

To run any solution, you can create a simple Go program that imports the solutions package and calls the function:

```go
package main

import (
    "fmt"
    "github.com/yourusername/codewars-go-solutions/solutions"
)

func main() {
    // Example for DuplicateEncode
    result := solutions.DuplicateEncode("recede")
    fmt.Println(result) // Should output "()()()"
}
```

## Testing

Each solution has been tested against the Codewars test cases to ensure correctness. You can also run tests locally by creating test files
for each solution.

## About Codewars

[Codewars](https://www.codewars.com/) is a platform that offers coding challenges called "kata" which are ranked by difficulty:

- 8 kyu: Beginner
- 7-6 kyu: Novice
- 5-4 kyu: Competent
- 3-1 kyu: Proficient
- 1-2 dan: Expert
- 3+ dan: Master

The higher the kyu number, the easier the challenge. The higher the dan number, the harder the challenge.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Feel free to fork this repository and add your own solutions or improvements to existing ones. Pull requests are welcome!
