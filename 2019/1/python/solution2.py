total=0
with open("../input.txt") as f:
	for line in f.readlines():
		# print(line[:-1], end=": ")
		fuel = int(int(line)/3-2)
		line_total=0
		while fuel > 0:
			# print(fuel, end=" ")
			line_total += fuel
			fuel = int(int(fuel)/3-2)
		total += line_total
		# print(line_total, total)

print(total)

# 5068210