import sys

columnString = sys.argv[1]
columnList = columnString.split(",")
limit = int(sys.argv[2])

# Remove extra white spaces
columnList = map(lambda x : x.strip(), columnList)
# Make name of column bold
columnList = map(lambda x : '**'+x+'**', columnList)
columnsLen = len(columnList)

# Print Header 
print("| "+" | ".join(columnList) + " |")

# Print separator for rows.
columnIdentifier = "|"
for i in range(columnsLen):
    columnIdentifier += " ----- |"
print(columnIdentifier)

# Print blank rows.
blankRows = ""
for i in range(columnsLen):
    blankRows += " | "
for i in range(1, limit+1):
    print("| " + blankRows)
