lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

numCardsCopies = [1 for _ in range(len(lines))]
for lineI, line in enumerate(lines):
    numCardCopies = numCardsCopies[lineI]
    line = line.strip()[8:]
    winningNumbers, yourNumbers = line.split("|")
    winningNumbers = winningNumbers.split(" ")
    yourNumbers = yourNumbers.split(" ")
    numOfYourNumberOccurences = 0
   
    for yourNumber in yourNumbers:
       if yourNumber == "":
           continue
       if yourNumber in winningNumbers:
           numOfYourNumberOccurences+=1
    
    if numOfYourNumberOccurences > 0:
        nextPosition = lineI+1
        for number in range(nextPosition, nextPosition+numOfYourNumberOccurences):
            numCardsCopies[number]+=numCardCopies

sumNumCopies = 0
for numCopies in numCardsCopies:
    sumNumCopies += numCopies

print(sumNumCopies)