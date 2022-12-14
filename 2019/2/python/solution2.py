total=0
for noun in range(100):
	for verb in range(100):
		with open("../input.txt") as f:
			array = list(map(int, f.read().split(',')))
			array[1], array[2] = noun, verb
			# print(array, int(len(array)/4))
			for x in range(0, len(array), 4):
				# print(array[x:x+4])
				if array[x]==1:
					array[array[x+3]] = array[array[x+1]] + array[array[x+2]]	
				elif array[x]==2:
					array[array[x+3]] = array[array[x+1]] * array[array[x+2]]	
				elif array[x]==99:
					break

			if array[0]==19690720:
				# print(noun, verb, array[0], array[0]==19690720)
				break
		if array[0]==19690720:
			break		
	if array[0]==19690720:
		break

print(100 * noun + verb)

# 19690720, 3146