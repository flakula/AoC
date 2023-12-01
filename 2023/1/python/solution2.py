import re


def replace_digits(line):
    one = line.find("one")
    # if one != -1:
    #     line = line[:one] + "1" + line[one:]
    two = line.find("two")
    # if two != -1:
    #     line = line[:two] + "2" + line[two:]
    three = line.find("three")
    # if three != -1:
    #     line = line[:three] + "3" + line[three:]
    four = line.find("four")
    # if four != -1:
    #     line = line[:four] + "4" + line[four:]
    five = line.find("five")
    # if five != -1:
    #     line = line[:five] + "5" + line[five:]
    six = line.find("six")
    # if six != -1:
    #     line = line[:six] + "6" + line[six:]
    seven = line.find("seven")
    # if seven != -1:
    #     line = line[:seven] + "7" + line[seven:]
    eight = line.find("eight")
    # if eight != -1:
    #     line = line[:eight] + "8" + line[eight:]
    nine = line.find("nine")
    # if nine != -1:
    #     line = line[:nine] + "9" + line[nine:]
    first_digit = re.findall(r"\d", line)[0]
    last_digit = re.findall(r"\d", line)[-1]

    first_string_digit = min(one, two, three, four, five, six, seven, eight, nine)
    last_string_digit = max(one, two, three, four, five, six, seven, eight, nine)

    print(line)
    return output


with open("../input.txt") as f:
    print(([replace_digits(line) for line in f.readlines()]))
