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


def part2(heights: List[List[int]]):
    scores = [
        [scenic_score(i, j, heights) for i in range(len(heights[j]))]
        for j in range(len(heights))
    ]
    max_score = 0
    for row in scores:
        max_score = max(max_score, max(row))

    print(f"part 2: {max_score}")


def scenic_score(i: int, j: int, heights: List[List[int]]) -> int:
    if i == 0 or j == 0 or i == len(heights) - 1 or j == len(heights[i]) - 1:
        return 0

    height = heights[i][j]
    above = below = left = right = 0
    for row in range(i - 1, -1, -1):
        above += 1
        if heights[row][j] >= height:
            break

    for row in range(i + 1, len(heights)):
        below += 1
        if heights[row][j] >= height:
            break

    for col in range(j - 1, -1, -1):
        left += 1
        if heights[i][col] >= height:
            break

    for col in range(j + 1, len(heights[i])):
        right += 1
        if heights[i][col] >= height:
            break

    return above * below * left * right


if __name__ == "__main__":
    heights = []
    with open(INPUT_FILE) as f:
        for row in f.readlines():
            heights += [[int(height) for height in row.strip()]]

    part1(heights)
    part2(heights)
