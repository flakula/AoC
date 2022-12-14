import functools

def manhattan(a):
    return abs(a[0]) + abs(a[1])

with open("../input.txt") as f:
	lines = f.read().split("\n")
	sets = []
	line_no = 0
	for line in lines:
		steps = line.split(",")
		# print(steps)
		if len(steps)<2:
			continue
		# print(steps)
		sets.append({})
		pos = (0,0)
		pasos = 0
		for step in steps:
			direction = step[0]
			step_count = int(step[1:])
			# print(step, direction, step_count)
			while step_count > 0:
				pasos +=1 
				if direction=="U":
					pos = (pos[0], pos[1] + 1)
				if direction=="D":
					pos = (pos[0], pos[1] - 1)
				if direction=="L":
					pos = (pos[0] - 1, pos[1])
				if direction=="R":
					pos = (pos[0] + 1, pos[1])
				sets[line_no][pos] = pasos
				step_count -=1

		line_no += 1

	# print(sorted(sets[0].items()))
	# print(sorted(sets[1].items()))

	intersecction = functools.reduce(lambda a, b: set(a.keys()&b.keys()), sets )
	# print(sorted(intersecction))
	# print(len(intersecction))

	added_steps = [sets[0][x]+sets[1][x] for x in intersecction]
	# print(sorted(added_steps))
	print(min(added_steps))

	# oneliner
	# print(min([sets[0][x]+sets[1][x] for x in functools.reduce(lambda a, b: set(a.keys()&b.keys()), sets )]))

#10554