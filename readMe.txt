
il programma goHTML_InpOut.exe     ( sorgente   goHTML_InpOut.go )

permette ad file HTML/JS in ambiente desktop di leggere e scrivere qualsiasi file di testo, 
senza alcuna necessità di modifiche ( lo stesso exe può essere usato da html diversi per gestire file diversi. 

La personalizzazione si trova soltanto nel file Html/Js.  L'utilizzo di file html diversi viene comunicato nei parametri nella linea di comando cmd.
I nomi dei file dentro html/js 

es.  file bat per chiamare il programma

-----------------------
@echo off 
set htmlFile="D:/ANTONIO/Anki_inputSchede/ProvaGo_HTML/mioHTML/gohtml_readWriteFile.html"

"../goHTML_InpOut.exe" -html %htmlFile% 
----------------------------------------

 file programma go, file html, file javascript, file di testo di input o di output possono stare anche in cartelle diverse tra loro
 
 vedi esempio di utilizzo nella cartella goLineByLine_HTML. 
  
