import functools

def manhattan(a):
    return abs(a[0]) + abs(a[1])

with open("../input.txt") as f:
	lines = f.read().split("\n")
	sets = []
	line_no = 0
	for line in lines:
		steps = line.split(",")
		if len(steps)<2:
			continue
		sets.append(set())
		pos = (0,0)
		for step in steps:
			direction = step[0]
			step_count = int(step[1:])
			# print(step, direction, step_count)
			while step_count > 0:
				if direction=="U":
					pos = (pos[0], pos[1] + 1)
				if direction=="D":
					pos = (pos[0], pos[1] - 1)
				if direction=="L":
					pos = (pos[0] - 1, pos[1])
				if direction=="R":
					pos = (pos[0] + 1, pos[1])
				sets[line_no].add(pos)
				step_count -=1

		line_no += 1

	# print(sorted(sets[0]))
	# print(sorted(sets[1]))

	intersecction = functools.reduce(lambda a, b: set(a&b), sets )
	# print(sorted(intersecction))
	# print(len(intersecction))

	distances = list(map(manhattan, intersecction ) ) 
	# # print(distances)
	# print(sorted(distances))
	print(min(distances))

	# oneliner
	# print(min(list(map(manhattan,functools.reduce(lambda a, b: set(a&b), sets)))))

# 280