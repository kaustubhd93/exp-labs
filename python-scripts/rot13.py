import re
import sys

sample = sys.argv[1]
rotDict = {"A" : "N",
           "B" : "O",
           "C" : "P",
           "D" : "Q",
           "E" : "R",
           "F" : "S",
           "G" : "T",
           "H" : "U",
           "I" : "V",
           "J" : "W",
           "K" : "X",
           "L" : "Y",
           "M" : "Z",
           "N" : "A",
           "O" : "B",
           "P" : "C",
           "Q" : "D",
           "R" : "E",
           "S" : "F",
           "T" : "G",
           "U" : "H",
           "V" : "I",
           "W" : "J",
           "X" : "K",
           "Y" : "L",
           "Z" : "M",
          }

def convert_to_text(rotText):
    convertedText = []
    for alphabet in rotText:
        try:
            convertedText.append(rotDict[alphabet.upper()])
        except KeyError:
            continue
    return "".join(convertedText)

if re.search(r'\s', sample, re.M | re.I):
    decoded = []
    tmp = sample.split(' ')
    for word in tmp:
        decoded.append(convert_to_text(word))
    decodedText = " ".join(decoded)

print decodedText.lower()

