lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

min1 = lambda val: 1 if val < 1 else val 

sumPowers = 0 
for line in lines:
    line = line.strip()
    subsets = line.split(":")[1].split(";")

    possibleGame = True
    minBlueCubesGame = 0
    minGreenCubesGame = 0
    minRedCubesGame = 0
    
    for subset in subsets:
        cubes = subset.split(",")

        numBlueCubes = 0
        numGreenCubes = 0
        numRedCubes = 0
        
        for cube in cubes:
            _, numCubeType, cubeType = cube.split(" ")
            numCubeType = int(numCubeType)
            if cubeType == "blue":
                numBlueCubes = numCubeType
            elif cubeType == "green":
                numGreenCubes = numCubeType
            elif cubeType == "red":
                numRedCubes = numCubeType
                
        if numBlueCubes > 0 and (numBlueCubes > minBlueCubesGame or minBlueCubesGame == 0):
            minBlueCubesGame = numBlueCubes
            
        if numGreenCubes > 0 and (numGreenCubes > minGreenCubesGame or minGreenCubesGame == 0):
            minGreenCubesGame = numGreenCubes
        
        if numRedCubes > 0 and (numRedCubes > minRedCubesGame or minRedCubesGame == 0):
            minRedCubesGame = numRedCubes
            

    sumPowers+= (min1(minBlueCubesGame) * min1(minGreenCubesGame) * min1(minRedCubesGame))
    
print(sumPowers)
