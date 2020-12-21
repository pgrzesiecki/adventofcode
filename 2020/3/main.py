import os
import math
path = os.path.abspath(os.path.dirname(__file__))


def count_trees(right, down):
    with open(path + '/input.txt', 'r') as f:
        area = f.readlines()

    height = len(area)
    width = len(area[0]) - 1

    trees_on_road = 0
    steps_down = math.ceil(height / down)

    y = 0
    for y in range(steps_down):
        x = (y * right)
        if x >= width:
            x = x % width

    for i in range(steps_down):
        x = (i * right)
        y = (i * down)

        if x >= width:
            x = x % width

        if area[y][x] == '#':
            trees_on_road += 1

    return trees_on_road


def A():
    print(count_trees(3, 1))


def B():
    a = count_trees(1, 1)
    b = count_trees(3, 1)
    c = count_trees(5, 1)
    d = count_trees(7, 1)
    e = count_trees(1, 2)

    print((a, b, c, d, e))

    print(a * b * c * d * e)


if __name__ == "__main__":
    A()
    B()
