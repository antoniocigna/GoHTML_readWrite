@echo off 
:: this file runs in a windows environment
:: specify absolute address in set htmlFile and exeFile  
:: the -parm parameter assume that 2 fields follow: folder address and file name of the parameter file    
:: -----------------------------------------------------
set htmlFile=D:/Users/Pc Anto/Documents/GitHub/LineByLine_V3/LineByLineV3/goLineByLine_html_readWriteFile.html
set exeFile=D:/Users/Pc Anto/Documents/GitHub/GoHTML_readWrite/goHTML_InpOut.exe

"%exeFile%" -html "%htmlFile%" -parm "./" "miei_parametri2.txt"

