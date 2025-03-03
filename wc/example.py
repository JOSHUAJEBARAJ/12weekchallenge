import os 


count=0 
with open('input.txt','r') as infile:
    for lines in infile:
        count= count+1
    
    print(count)