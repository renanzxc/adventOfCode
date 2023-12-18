def validateDigit(letter, letters, lenLetters, numberWords, isReverse):    
    if isReverse:
        numberWords = numberWords[::-1]

    if letter != numberWords[0]:
        return False

    lenNumberWords = len(numberWords)

    return  lenLetters >= lenNumberWords and letters[0:lenNumberWords] == numberWords

def getDigit(letters, isReverse):
    actualDigit = ''
    letter = letters[0]
    lenLetters = len(letters)

    if letter.isdigit():
        actualDigit = letter
    else:
        if validateDigit(letter, letters, lenLetters, 'one', isReverse):
            actualDigit = '1'
        elif validateDigit(letter, letters, lenLetters, 'two', isReverse):
            actualDigit = '2'
        elif validateDigit(letter, letters, lenLetters, 'three', isReverse):
            actualDigit = '3'
        if validateDigit(letter, letters, lenLetters, 'four', isReverse):
            actualDigit = '4'
        elif validateDigit(letter, letters, lenLetters, 'five', isReverse):
            actualDigit = '5'
        elif validateDigit(letter, letters, lenLetters, 'six', isReverse):
            actualDigit = '6'
        elif validateDigit(letter, letters, lenLetters, 'seven', isReverse):
            actualDigit = '7'
        elif validateDigit(letter, letters, lenLetters, 'eight', isReverse):
            actualDigit = '8'
        elif validateDigit(letter, letters, lenLetters, 'nine', isReverse): 
            actualDigit = '9'
    
    if actualDigit != '':
        return actualDigit
    

    return getDigit(letters[1:], isReverse)

lines = []

with open('input.txt', 'r') as file:
    lines = file.readlines()

totalCalibrationValue = 0
for line in lines:
    digits = []
    calibrationValue = 0
    line = line.strip()
    digit1 = getDigit(line, False)
    digit2 = getDigit(line[::-1], True)

    calibrationValue = int(digit1 + digit2)
    
    totalCalibrationValue += calibrationValue

print(totalCalibrationValue)


