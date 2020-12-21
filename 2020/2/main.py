import os
path = os.path.abspath(os.path.dirname(__file__))


def A():
    with open(path + '/input.txt', 'r') as f:
        lines = f.readlines()
        passwords_valid = 0
        for line in lines:
            (condition,
             password) = [x.strip() for x in line.strip().split(':')]
            (times, required_letter) = condition.strip().split(' ')
            (tfrom, tto) = [int(x) for x in times.strip().split('-')]

            i = 0
            for letter in password:
                if letter == required_letter:
                    i += 1

            if i >= tfrom and i <= tto:
                passwords_valid += 1

    print(passwords_valid)


def B():
    with open(path + '/input.txt', 'r') as f:
        lines = f.readlines()
        passwords_valid = 0
        for line in lines:
            (condition,
             password) = [x.strip() for x in line.strip().split(':')]
            (poses, letter) = condition.strip().split(' ')
            (pos_a, pos_b) = [int(x) for x in poses.strip().split('-')]

            if (password[pos_a - 1] == letter or password[pos_b - 1]
                    == letter) and password[pos_a - 1] != password[pos_b - 1]:
                passwords_valid += 1

    print(passwords_valid)


if __name__ == "__main__":
    A()
    B()
