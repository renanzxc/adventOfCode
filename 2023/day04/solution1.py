lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

sumCardsValues = 0
for lineI, line in enumerate(lines):
    line = line.strip()[8:]
    winningNumbers, yourNumbers = line.split("|")
    winningNumbers = winningNumbers.split(" ")
    yourNumbers = yourNumbers.split(" ")
    numOfYourNumberOccurences = 0
    cardValue = 0
   
    for yourNumber in yourNumbers:
       if yourNumber == "":
           continue
       if yourNumber in winningNumbers:
           numOfYourNumberOccurences+=1
       
    if numOfYourNumberOccurences == 0:
        cardValue = 0
    elif numOfYourNumberOccurences == 1:
        cardValue = 1
    else:
        cardValue = 2 ** (numOfYourNumberOccurences - 1)
        
    sumCardsValues += cardValue
print(sumCardsValues)