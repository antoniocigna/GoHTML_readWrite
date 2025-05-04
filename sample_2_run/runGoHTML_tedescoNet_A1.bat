@echo off 
:: this file runs in a windows environment
:: specify absolute address in set htmlFile and exeFile  
:: the -parm parameter assume that 2 fields follow: folder address and file name of the parameter file    
::       you might use "%~n0.txt" as parameter file name, so that many copies of this file without change can be used to manage different file,   
::            just make a copy of this bat file and rename it as the parameter file name (apart the extension)        
:: -----------------------------------------------------

set htmlFile=D:/Users/Pc Anto/Documents/GitHub/GoHTML_readWrite/goLineByLine_HTML/goLineByLine_html_readWriteFile.html
set exeFile=D:/Users/Pc Anto/Documents/GitHub/GoHTML_readWrite/goHTML_InpOut.exe

:: "%exeFile%" -html "%htmlFile%" -parm "./" "parametri01.txt"

"%exeFile%" -html "%htmlFile%" -parm "./" "%~n0.txt"



