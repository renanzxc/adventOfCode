def isSymbol(char):
    return char != '.' and not char.isdigit()

lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

lenLines = len(lines)
sumAllValidGroups = 0
for lineI, line in enumerate(lines):
    line = line.strip()
    lenLine = len(line)
    actualGroup = ''
    actualGroupHasAdjSy = False 

    for charI, char in enumerate(line):                
        digitEndOfLine = False 
        if char.isdigit():
            actualGroup += char
            if lineI - 1 >= 0:
                charUp = lines[lineI - 1][charI]
                if isSymbol(charUp):
                    actualGroupHasAdjSy = True
            if lineI + 1 < lenLines:
                charDown = lines[lineI + 1][charI]
                if isSymbol(charDown):
                    actualGroupHasAdjSy = True
            if charI - 1 >= 0:
                charLeft = line[charI - 1]
                if isSymbol(charLeft):
                    actualGroupHasAdjSy = True
            if charI + 1 < lenLine:
                charRight = line[charI + 1]
                if isSymbol(charRight):
                    actualGroupHasAdjSy = True
            else:
               digitEndOfLine = True
                    
            if lineI - 1 >= 0 and charI - 1 >= 0:
                diagUpLeft = lines[lineI - 1][charI - 1]
                if isSymbol(diagUpLeft):
                    actualGroupHasAdjSy = True
            if lineI - 1 >= 0 and charI + 1 < lenLine:
                diagUpRight = lines[lineI - 1][charI + 1]
                if isSymbol(diagUpRight):
                    actualGroupHasAdjSy = True
            if lineI + 1 < lenLines and charI - 1 >= 0:
                diagDownLeft = lines[lineI + 1][charI - 1]
                if isSymbol(diagDownLeft):
                    actualGroupHasAdjSy = True     
            if lineI + 1 < lenLines and charI + 1 < lenLine:
                diagDownRight = lines[lineI + 1][charI + 1]
                if isSymbol(diagDownRight):
                    actualGroupHasAdjSy = True
        elif actualGroup != "" or charI + 1 == lenLine:
            if actualGroupHasAdjSy:
                sumAllValidGroups+=int(actualGroup)
            actualGroup = ""
            actualGroupHasAdjSy = False
        
        if digitEndOfLine:
            if actualGroupHasAdjSy:
                sumAllValidGroups+=int(actualGroup)
            actualGroup = ""
            actualGroupHasAdjSy = False     
   
        
print(sumAllValidGroups)
