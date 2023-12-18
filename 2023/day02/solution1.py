lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

maxRedCubes = 12
maxGreenCubes = 13
maxBlueCubes = 14

gameID = 1
sumIDs = 0 
for line in lines:
    line = line.strip()
    subsets = line.split(":")[1].split(";")

    possibleGame = True
    for subset in subsets:
        cubes = subset.split(",")
     
        for cube in cubes:
            _, numCubeType, cubeType = cube.split(" ")
            numCubeType = int(numCubeType)
            if cubeType == "blue" and numCubeType > maxBlueCubes:
                possibleGame = False
                break
            elif cubeType == "green" and numCubeType > maxGreenCubes:
                possibleGame = False
                break
            elif cubeType == "red" and numCubeType > maxRedCubes:
                possibleGame = False
                break
            
    if possibleGame:
        sumIDs+=gameID
    
    gameID+=1
print(sumIDs)