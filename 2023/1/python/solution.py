import re

with open("../input.txt") as f:
    print(
        sum(
            [
                int(re.findall(r"\d", line)[0] + re.findall(r"\d", line)[-1])
                for line in f.readlines()
            ]
        )
    )
