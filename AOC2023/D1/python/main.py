#  read input.txt file as string 
inputFile = open('input.txt', 'r')
inputString = inputFile.readLines()
for line in inputString:
    print(line+ ": " )
    
