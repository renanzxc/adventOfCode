import re

lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()


timesRaces = re.sub(r' +', ' ', lines[0][5:]).strip().split(" ")
recordsDistance = re.sub(r' +', ' ', lines[1][9:]).strip().split(" ")

mulWaysRecord = 1
hasAnyWay = False

for raceI, timeRace in enumerate(timesRaces):
    timeRace = int(timeRace)
    recordDistance = int(recordsDistance[raceI])
    minWay = 0
    maxWay = 0
    maxHoldMS = timeRace-1

    tryHoldMS = 1
    while maxHoldMS >= tryHoldMS:
        if tryHoldMS * (timeRace - tryHoldMS) > recordDistance:
            minWay = tryHoldMS
            break
        tryHoldMS+=1

    tryHoldMS = maxHoldMS
    while 0 < tryHoldMS:
        if tryHoldMS * (timeRace - tryHoldMS) > recordDistance:
            maxWay = tryHoldMS
            break
        tryHoldMS-=1
    
    if minWay != 0 and maxWay != 0:
        mulWaysRecord *= (maxWay - minWay + 1)
        hasAnyWay = True

if hasAnyWay:
    print(mulWaysRecord)
else:
    print(0)
