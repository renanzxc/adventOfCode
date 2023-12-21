import re

lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

timeRace = re.sub(r' +', '', lines[0][5:]).strip()
recordDistance = re.sub(r' +', '', lines[1][9:]).strip()

minWay = 0
maxWay = 0
timeRace = int(timeRace)
recordDistance = int(recordDistance)
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
    print(maxWay - minWay + 1)
else:
    print(0)
