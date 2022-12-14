total=0
with open("../input.txt") as f:
	array = list(map(int, f.read().split(',')))
	array[1] = 12
	array[2] = 2
	# print(array, int(len(array)/4))
	for x in range(0, len(array), 4):
		# print(array[x:x+4])
		if array[x]==1:
			array[array[x+3]] = array[array[x+1]] + array[array[x+2]]	
		elif array[x]==2:
			array[array[x+3]] = array[array[x+1]] * array[array[x+2]]	
		elif array[x]==99:
			break

	# print(array)

print(array[0])

# 8017076