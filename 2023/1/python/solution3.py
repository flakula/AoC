import re

DIGIT_STRINGS = {
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}
regex = r"\d|" + r"|".join(DIGIT_STRINGS)
regex_reverse = r"\d|" + r"|".join(d[::-1] for d in DIGIT_STRINGS)


def get_total_sum(file_path: str) -> int:
    def get_value(line: str) -> int:
        first_digit = re.search(regex, line)[0]
        last_digit = re.search(regex_reverse, line[::-1])[0][::-1]
        return int(
            f"{DIGIT_STRINGS.get(first_digit, first_digit)}{DIGIT_STRINGS.get(last_digit, last_digit)}"
        )

    with open(file_path, "r") as f:
        return sum(get_value(line) for line in f.readlines())


print(get_total_sum("../input.txt"))
