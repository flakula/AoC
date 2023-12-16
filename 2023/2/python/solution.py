with open("../input.txt") as f:
    games = [g.split(":") for g in f.read().split("\n")]

def possible_games_ids(games, red, green, blue):
    possible = []
    for game in games:
        id_ = game[0].split()[-1]
        sets = game[1].split(";")
        cubes = {'red': red, 'green': green, 'blue': blue}
        valid = True
        for set_ in sets:
            turns = set_.strip().split(",")
            for turn in turns:
                count, color = turn.split()
                print(count, color)
                count = int(count)
                if cubes[color] < count:
                    valid = False
                    break
                # else:
                #     cubes[color] -= count
            if not valid:
                break
        if valid:
            possible.append(int(id_))

    print(possible)
    return possible


red_cubes = 12
green_cubes = 13
blue_cubes = 14

sum_of_ids = sum(possible_games_ids(games, red_cubes, green_cubes, blue_cubes))
print(sum_of_ids)
