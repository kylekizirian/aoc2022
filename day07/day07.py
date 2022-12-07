from enum import Enum
from pathlib import Path
from typing import Dict, List, Optional

INPUT_FILE = Path(__file__).parent / "input.txt"
SIZE_CAP = 100000
SPACE_ALLOWED = 70000000 - 30000000


class Kind(Enum):
    DIR = 1
    FILE = 2


class Entity:
    def __init__(
        self,
        kind: Kind,
        name: str,
        size: int,
        parent=None,
        children=None,
    ):
        self.kind = kind
        self.name = name
        self.size = size
        self.parent = parent
        self.children = {} if children is None else children


def buildTree(lines: List[str]) -> Entity:
    root = Entity(Kind.DIR, "/", 0, None)
    curDir = root
    i = 0

    while i < len(lines):
        line = lines[i]
        _, cmd, *args = line.split()

        if cmd == "cd":
            if args[0] == "/":
                curDir = root
            elif args[0] == "..":
                curDir = curDir.parent
            else:
                curDir = curDir.children[args[0]]

            i += 1

        elif cmd == "ls":
            # iterate to next command
            i += 1
            while i < len(lines) and not lines[i].startswith("$"):
                line = lines[i]
                i += 1
                parts = line.split()

                # skip if we've already seen this file/folder
                if parts[1] in curDir.children:
                    continue

                if parts[0] == "dir":
                    curDir.children[parts[1]] = Entity(Kind.DIR, parts[1], 0, curDir)
                else:
                    size = int(parts[0])
                    curDir.children[parts[1]] = Entity(
                        Kind.FILE, parts[1], size, curDir
                    )
                    temp = curDir
                    while temp is not None:
                        temp.size += size
                        temp = temp.parent

    return root


def part1(root: Entity):
    # DFS from root and find dirs
    size_sum = 0
    stack = [root]
    while stack:
        cur = stack.pop()
        if cur.size <= SIZE_CAP:
            size_sum += cur.size
        stack += [child for child in cur.children.values() if child.kind == Kind.DIR]

    print(f"part 1: {size_sum}")


def part2(root: Entity):
    space_needed = root.size - SPACE_ALLOWED
    # find smallest dir that's greater than or equal to space_needed
    smallest = float("inf")
    stack = [root]
    while stack:
        cur = stack.pop()
        if cur.size >= space_needed and cur.size < smallest:
            smallest = cur.size
        stack += [child for child in cur.children.values() if child.kind == Kind.DIR]

    print(f"part 2: {smallest}")


if __name__ == "__main__":
    with open(INPUT_FILE) as f:
        lines = [l.strip() for l in f.readlines()]

    root = buildTree(lines)
    part1(root)
    part2(root)
