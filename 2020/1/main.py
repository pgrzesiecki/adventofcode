import os
path = os.path.abspath(os.path.dirname(__file__))


def A():
    with open(path + '/input.txt', 'r') as f:
        numbers = f.readlines()
        numbers = [x.strip() for x in numbers]

    found = []
    for idx, number in enumerate(numbers):
        for sec_number in numbers[idx:]:
            number = int(number)
            sec_number = int(sec_number)
            if number + sec_number == 2020:
                found.append((number, sec_number))

    for item in found:
        print("A: {} x {} = {}".format(item[0], item[1], item[0] * item[1]))


def B():
    with open(path + '/input.txt', 'r') as f:
        numbers = f.readlines()
        numbers = [x.strip() for x in numbers]

    found = []
    for idx, number in enumerate(numbers):
        for sec_number in numbers[idx:]:
            for third_number in numbers[idx:]:
                number = int(number)
                sec_number = int(sec_number)
                third_number = int(third_number)
                if number + sec_number + third_number == 2020:
                    found.append((number, sec_number, third_number))

    for item in found:
        print("B: {} x {} x {} = {}".format(item[0], item[1], item[2],
                                            item[0] * item[1] * item[2]))


if __name__ == "__main__":
    A()
    B()
