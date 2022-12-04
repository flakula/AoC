with open("../input.txt") as file:
	print(max(map(lambda e: sum(map(int, e.split("\n"))), file.read().split("\n\n"))))
