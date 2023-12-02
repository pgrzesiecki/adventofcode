import sys
import math

exercise_2a_config = {"red": 12, "green": 13, "blue": 14}


def calculateSets(setLine):
    sets = setLine.split(";")
    sumNum = {}
    maxNum = {}

    for i in sets:
        game = i.strip().split(",")

        for g in game:
            [num, color] = g.strip().split(" ")
            num = int(num)

            if color in sumNum:
                sumNum[color] += num
            else:
                sumNum[color] = num

            if color not in maxNum or maxNum[color] < num:
                maxNum[color] = num

    return (maxNum, sumNum)


def canBePlayed(sets, settings):
    for color in settings:
        if color not in sets:
            return False

        if sets[color] > settings[color]:
            return False

    return True


def exercise2a(games):
    gameSum = 0
    for game in games:
        settings = game.split(":")

        if len(settings) != 2:
            continue

        gameID = int(settings[0].strip().replace("Game ", ""))
        (maxNum, _) = calculateSets(settings[1])

        if canBePlayed(maxNum, exercise_2a_config):
            gameSum += gameID

    return gameSum


def exercise2b(games):
    sumOfPowers = 0
    for game in games:
        settings = game.split(":")

        if len(settings) != 2:
            continue

        (maxNum, _) = calculateSets(settings[1])

        values = [maxNum[i] for i in maxNum]
        sumOfPowers += math.prod(values)

    return sumOfPowers


# python 2.py < 2.txt
# python 2.py < 2b.txt
if __name__ == "__main__":
    games = sys.stdin.read().split("\n")
    print(f"A - {exercise2a(games)}")
    print(f"B - {exercise2b(games)}")
