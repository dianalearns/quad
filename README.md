# Quadrilateral Drawing Project

A beginner-friendly Go project that demonstrates how to draw different types of quadrilaterals (rectangles) using ASCII characters in the terminal.

## Overview

This project contains five different functions that draw rectangles with various border styles. Each function uses a core `Quad()` function with different parameters to create unique visual patterns.

## Project Structure

```sh
quad/
├── go.mod                 # Go module configuration
├── quada.go              # Draws rectangles with 'o' corners and '|'/'-' borders
├── quadb.go              # Draws rectangles with '/','\\' corners and '*' borders
├── quadc.go              # Draws rectangles with 'A','C' corners and 'B' borders
├── quadd.go              # Draws rectangles with alternating 'A','C' corners
├── quade.go              # Draws rectangles with mirrored 'A','C' corners
├── piscine_test.go       # Comprehensive test suite
└── test/main.go          # Example usage of all functions
```

## How It Works

### The Core `Quad()` Function

All five functions use a single core function that takes these parameters:

```go
func Quad(x, y int, tlc, trc, blc, brc, vl, hl rune)
```

**Parameters explained:**

- `x` - Width of the rectangle (number of columns)
- `y` - Height of the rectangle (number of rows)
- `tlc` - Top-left corner character
- `trc` - Top-right corner character
- `blc` - Bottom-left corner character
- `brc` - Bottom-right corner character
- `vl` - Vertical line character (sides)
- `hl` - Horizontal line character (top and bottom)

The function works by:

1. Looping through each row (from 1 to y)
2. For each row, looping through each column (from 1 to x)
3. Checking the position to decide which character to print:
   - Corners get special characters (tlc, trc, blc, brc)
   - Top/bottom edges get horizontal line characters (hl)
   - Left/right edges get vertical line characters (vl)
   - Interior gets spaces
4. Printing a newline after each row

## The Five Quad Functions

### QuadA - Classic Rectangle

Draws a rectangle with `'o'` corners, `'|'` sides, and `'-'` top/bottom.

**Code:**

```go
piscine.QuadA(5, 3)
```

**Output:**

```sh
o---o
|   |
o---o
```

**Use case:** Perfect for creating simple boxes, frames, or borders around text.

---

### QuadB - Diagonal Rectangle

Draws a rectangle with `'/'` and `'\\'` corners, `'*'` for everything else.

**Code:**

```go
piscine.QuadB(5, 3)
```

**Output:**

```terminal
/---\
|   |
\---/
```

**Use case:** Creates a diamond-like effect, useful for highlighting or decorative borders.

---

### QuadC - Letter A/C Rectangle

Draws a rectangle with `'A'` corners (top), `'C'` corners (bottom), and `'B'` for edges.

**Code:**

```go
piscine.QuadC(5, 3)
```

**Output:**

```sh
A---A
|   |
C---C
```

**Use case:** Demonstrates using letters as design elements; good for labeling sections.

---

### QuadD - Alternating Corners

Draws a rectangle with alternating `'A'` and `'C'` corners.

**Code:**

```go
piscine.QuadD(5, 3)
```

**Output:**

```sh
A---C
|   |
A---C
```

**Use case:** Creates an asymmetric design; useful for directional indicators.

---

### QuadE - Mirrored Corners

Draws a rectangle with mirrored `'A'` and `'C'` corners.

**Code:**

```go
piscine.QuadE(5, 3)
```

**Output:**

```sh
A---C
|   |
C---A
```

**Use case:** Creates a balanced, symmetrical pattern; good for decorative elements.

## Running the Examples

### Run the test program

```bash
go run test/main.go
```

This will display several examples of different quadrilaterals. Uncomment different lines in `test/main.go` to see various outputs.

### Run the tests

```bash
go test -v
```

This runs comprehensive tests for all five functions with various input sizes.

## Examples for Beginners

### Example 1: Simple Box

```go
package main

import "piscine"

func main() {
    // Draw a 10x5 box
    piscine.QuadA(10, 5)
}
```

**Output:**

```sh
o----------o
|          |
|          |
|          |
o----------o
```

### Example 2: Small Square

```go
package main

import "piscine"

func main() {
    // Draw a 3x3 square
    piscine.QuadA(3, 3)
}
```

**Output:**

```sh
o-o
| |
o-o
```

### Example 3: Single Line

```go
package main

import "piscine"

func main() {
    // Draw a horizontal line (1 row tall)
    piscine.QuadA(7, 1)
}
```

**Output:**

```sh
o-----o
```

### Example 4: Single Column

```go
package main

import "piscine"

func main() {
    // Draw a vertical line (1 column wide)
    piscine.QuadA(1, 5)
}
```

**Output:**

```sh
o
|
|
|
o
```

### Example 5: Mixing Different Styles

```go
package main

import "piscine"

func main() {
    // Draw different styles side by side
    piscine.QuadA(5, 3)
    println()
    piscine.QuadB(5, 3)
    println()
    piscine.QuadC(5, 3)
}
```

## Understanding the Code

### Key Concepts Used

1. **Rune**: A Go data type that represents a Unicode code point (essentially a character)
2. **Nested Loops**: The outer loop handles rows (height), inner loop handles columns (width)
3. **Conditional Logic**: `if/else` or `switch` statements determine which character to print based on position
4. **Position Checking**: By comparing current row/column with total rows/columns, we know if we're at a corner, edge, or interior

### The Logic Pattern

```sh
For each row from 1 to y:
    For each column from 1 to x:
        IF at top-left corner: print top-left character
        ELSE IF at top-right corner: print top-right character
        ELSE IF at bottom-left corner: print bottom-left character
        ELSE IF at bottom-right corner: print bottom-right character
        ELSE IF at top or bottom edge: print horizontal line
        ELSE IF at left or right edge: print vertical line
        ELSE: print space (interior)
    Print newline
```

## Common Use Cases

- Creating text-based user interfaces
- Building simple games (like tic-tac-toe boards)
- Formatting terminal output
- Learning about loops and conditionals
- Understanding coordinate systems

## Tips for Beginners

1. **Start small**: Try `QuadA(1, 1)` first to see a single point
2. **Experiment**: Change the numbers and see what happens
3. **Trace the code**: Follow along with pencil and paper for a 3x3 rectangle
4. **Modify it**: Try creating your own function with different characters
5. **Break it**: See what happens with 0 or negative numbers

## Troubleshooting

**Q: Nothing prints?**
A: Make sure you're calling the function and your terminal supports the output

**Q: The rectangle looks wrong?**
A: Check that x and y are positive numbers greater than 0

**Q: Characters overlap?**
A: This shouldn't happen, but ensure you're using the function correctly

## Learning Resources

- [Go Tour](https://go.dev/tour/) - Official Go tutorial
- [Go by Example](https://gobyexample.com/) - Practical Go examples
- [The Go Programming Language Specification](https://go.dev/ref/spec) - Official documentation

## Automated Testing

The `piscine_test.go` file contains comprehensive tests for all five functions. You can run these tests using the command:

```bash
    go test -v
```

## License

This is an educational project for learning Go programming.

## Author

Created as part of the gitlab piscine programming curriculum.
