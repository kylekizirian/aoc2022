from pathlib import Path
from typing import List

INPUT_FILE = Path(__file__).parent / "input.txt"


def part1(heights: List[List[int]]):
    visible = []
    for row in heights:
        visible += [[False for _ in row]]

    for i in range(len(heights)):
        row = heights[i]
        tallest = -1
        for j in range(len(row)):
            height = row[j]
            if height > tallest:
                visible[i][j] = True
                tallest = height

    for i in range(len(heights)):
        row = heights[i]
        tallest = -1
        for j in range(len(row) - 1, -1, -1):
            height = row[j]
            if height > tallest:
                visible[i][j] = True
                tallest = height

    for i in range(len(heights)):
        col = [row[i] for row in heights]
        tallest = -1
        for j in range(len(col)):
            height = col[j]
            if height > tallest:
                visible[j][i] = True
                tallest = height

    for i in range(len(heights)):
        col = [row[i] for row in heights]
        tallest = -1
        for j in range(len(col) - 1, -1, -1):
            height = col[j]
            if height > tallest:
                visible[j][i] = True
                tallest = height

    num_visible = sum(sum(row) for row in visible)
    print(f"part 1: {num_visible}")


if __name__ == "__main__":
    heights = []
    with open(INPUT_FILE) as f:
        for row in f.readlines():
            heights += [[int(height) for height in row.strip()]]

    part1(heights)
