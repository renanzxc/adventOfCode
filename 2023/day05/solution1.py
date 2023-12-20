class CategoryNumber:
    def __init__(self, number):
        self.number = number 
        self.converted = False

lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()
    
seedsNeedPlanted = lines[0].strip()[7:].split(" ")

actualCategoryNumbers = []
for number in seedsNeedPlanted:
    actualCategoryNumbers.append(CategoryNumber(int(number))) 

for lineI, line in enumerate(lines[2:]):
    line = line.strip()
    if line == "":
        continue
    elif line.find(":") >= 0:
        for numberI, number in enumerate(actualCategoryNumbers):
            actualCategoryNumbers[numberI].converted = False
        continue
    
    line = line.split(" ") 
    rangeStrDest, rangeStrSource, range = int(line[0]), int(line[1]), int(line[2])
    
    for numberI, number in enumerate(actualCategoryNumbers):
        if number.converted:
            continue
        if rangeStrSource <= number.number and rangeStrSource+range > number.number:
            diff = number.number - rangeStrSource
            actualCategoryNumbers[numberI].number = rangeStrDest+diff
            actualCategoryNumbers[numberI].converted = True
    

lowestLocationNumber = ''
for number in actualCategoryNumbers:   
    if lowestLocationNumber == '' or number.number < lowestLocationNumber:
        lowestLocationNumber = number.number
        
print(lowestLocationNumber)