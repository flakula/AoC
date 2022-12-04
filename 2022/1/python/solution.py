with open("../input.txt") as f:
	print(max(map(lambda e: sum(map(int, e.split("\n"))), f.read().split("\n\n"))))
