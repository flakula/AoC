with open("../input.txt") as f:
    games = [g.split(":") for g in f.read().split("\n")]

def possible_games_ids(games):
    out = []
    for game in games:
        id_ = game[0].split()[-1]
        sets = game[1].split(";")
        cubes = {'red': 0, 'green': 0, 'blue': 0}
        for set_ in sets:
            turns = set_.strip().split(",")
            for turn in turns:
                count, color = turn.split()
                # print(count, color)
                count = int(count)
                if cubes[color] < count:
                    cubes[color] = count
        out.append(cubes)

    # print(out)
    power_out = sum(map(lambda x: x["red"]*x["green"]*x["blue"], out))
    return power_out

print(possible_games_ids(games))
