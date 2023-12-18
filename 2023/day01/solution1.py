lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

totalCalibrationValue = 0
for line in lines:
    digits = []
    calibrationValue = 0
    line = line.strip()
    for letter in line:
        if letter.isdigit():
            digit1 = letter
            break
    for letter in line[::-1]:
        if letter.isdigit():
            digit2 = letter
            break

    calibrationValue = int(digit1 + digit2)
        
    totalCalibrationValue+=calibrationValue

print(totalCalibrationValue)
