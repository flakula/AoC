total=0
with open("../input.txt") as f:
	for line in f.readlines():
		fuel = int(int(line)/3-2)
		total += fuel
		# print(total, fuel)

print(total)

# 3380731