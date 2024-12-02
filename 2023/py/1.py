import sys

numberAsWord = [
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
]
numberAsInt = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]


def getNumber(input, useWords=True):
    input = input.strip()

    if input == "":
        return None

    first = last = None
    firstPos = lastPos = None

    if useWords:
        for idx, i in enumerate(numberAsWord):
            occ = [j for j in range(len(input)) if input.startswith(i, j)]
            if not occ:
                continue

            if firstPos is None or occ[0] < firstPos:
                firstPos = occ[0]
                first = idx + 1

            if lastPos is None or occ[-1] > lastPos:
                lastPos = occ[-1]
                last = idx + 1

    for i in numberAsInt:
        occ = [j for j in range(len(input)) if input.startswith(i, j)]
        if not occ:
            continue

        if firstPos is None or occ[0] < firstPos:
            firstPos = occ[0]
            first = i

        if lastPos is None or occ[-1] > lastPos:
            lastPos = occ[-1]
            last = i

    if first is None:
        return None

    if last is None:
        last = first

    return int(f"{first}{last}")


def getValue(input, useWords=True):
    values = [getNumber(i.strip(), useWords) for i in input]
    return sum([i for i in values if i is not None])


# python 1.py < 1.txt (useWords=False )
# python 1.py < 1b.txt (useWords=True)
if __name__ == "__main__":
    data = sys.stdin.read().split("\n")
    print(
        f"A - {getValue([i for i in data if i.strip() != ""], useWords=True)}"
    )
    print(
        f"B - {getValue([i for i in data if i.strip() != ""], useWords=False)}"
    )
