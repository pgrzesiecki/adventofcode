import sys
import numpy as np


def prepare_schema(lines):
    schema = []
    numbers_schema = {}
    num_start = None
    num_end = None

    def add_to_number_schema(num_start, num_end, line_idx, numbers_schema):
        if num_start is not None and num_end is not None:
            if line_idx not in numbers_schema:
                numbers_schema[line_idx] = []

            numbers_schema[line_idx].append((num_start, num_end))

        return numbers_schema

    for line_idx, line in enumerate(lines):
        if line == "":
            continue

        if line_idx not in schema:
            schema.append([])

        for char_idx, char in enumerate(line):
            if char_idx not in schema[line_idx]:
                schema[line_idx].append([])

            schema[line_idx][char_idx] = char

            if char.isdigit():
                if num_start is None:
                    num_start = char_idx

                num_end = char_idx
            else:
                numbers_schema = add_to_number_schema(
                    num_start, num_end, line_idx, numbers_schema
                )

                num_start = None
                num_end = None

        numbers_schema = add_to_number_schema(
            num_start, num_end, line_idx, numbers_schema
        )

        num_start = None
        num_end = None

    return (schema, numbers_schema)


def is_symbol(item):
    return item != "." and item.isdigit() is False


def check_borders(num_start, num_end, num_row, schema):
    next_row = num_row + 1
    prev_row = num_row - 1

    is_next_row = (num_row + 1) < len(schema)
    is_prev_row = num_row > 0

    next_char = num_end + 1
    prev_char = num_start - 1

    is_next_char = (next_char + 1) < len(schema[num_row])
    is_prev_char = prev_char > 0

    iter_from = prev_char if is_prev_char else num_start
    iter_to = next_char if is_next_char else num_end

    # check left
    if is_prev_char and is_symbol(schema[num_row][prev_char]):
        return True

    # check right
    if is_next_char and is_symbol(schema[num_row][next_char]):
        return True

    for i in range(iter_from, iter_to + 1):
        # Check top
        if is_prev_row and is_symbol(schema[prev_row][i]):
            return True

        # Check bottom
        if is_next_row and is_symbol(schema[next_row][i]):
            return True

    return False


def get_number_from_coords(num_start, num_end, num_row, schema):
    number = ""
    for i in range(num_start, num_end + 1):
        number += schema[num_row][i]

    return int(number)


def find_numbers_around(matrix, positions):
    rows, cols = matrix.shape
    positions_with_numbers = []

    for row, col in positions:
        print(f"row = {row}, col = {col}, {matrix[row][col]}")
        local_set = set()

        for dr, dc in [
            (-1, -1),
            (-1, 0),
            (-1, 1),
            (1, -1),
            (1, 0),
            (1, 1),
            (0, -1),
            (0, 1),
        ]:
            r, c = row + dr, col + dc
            if 0 <= r < rows and 0 <= c < cols and matrix[r, c].isdigit():
                print(f"r = {r}, c = {c}, {matrix[r][c]}")
                local_set.add(find_number_in(r, c, matrix))

    print(matrix)
    print(local_set)
    return positions_with_numbers


def find_number_in(row, col, matrix):
    (col_start, col_end) = (col, col)
    (found_start, found_end) = (False, False)
    number = ""

    while True:
        if matrix[row][col_start - 1].isdigit():
            col_start -= 1
        else:
            found_start = True

        if matrix[row][col_end + 1].isdigit():
            col_end += 1
        else:
            found_end = True

        if found_start and found_end:
            break

    for i in range(col_start, col_end):
        number += matrix[row][i]

    print(f"number {number} ({col_start}-{col_end} (row: {row}))")
    return int(number)


def exerciseA(matrix):
    is_digit = np.char.isdigit(matrix)
    is_dot = matrix == "."

    mask = np.logical_not(np.logical_or(is_digit, is_dot))

    positions = np.where(mask)
    special_char_positions = list(zip(positions[0], positions[1]))

    print(special_char_positions)

    print(find_numbers_around(matrix, special_char_positions))

    return

    # (schema, number_schema) = prepare_schema(lines)

    # sum = 0

    # for row_num in number_schema:
    #     for number_coords in number_schema[row_num]:
    #         (num_start, num_end) = number_coords
    #         num_start = int(num_start)
    #         num_end = int(num_end)

    #         if check_borders(num_start, num_end, row_num, schema):
    #             sum += get_number_from_coords(
    #                 num_start, num_end, row_num, schema
    #             )

    # return sum


def exerciseB():
    pass


if __name__ == "__main__":
    matrix = np.array(
        [
            np.array(list(line), dtype=str)
            for line in sys.stdin.read().split("\n")
        ]
    )

    print(f"A - {exerciseA(np.char.array(matrix))}")
    # print(f"B - {exerciseB(lines)}")
