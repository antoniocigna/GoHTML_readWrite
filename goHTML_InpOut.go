package main

	import (
		"fmt"
		"os"
		"os/signal"
		"strings"		
		"strconv"
		"runtime"	
		"github.com/zserge/lorca"	
		"github.com/lxn/win"	
		"bufio"		
		"io"			
	)
//---------------------
//g00_main.go
//------------------------------------------------------
//--------------------------
//g01_declare.go
//-----------------------------------------

const appName = "gohtml_sample"
const htmlFile= appName + ".html"
const parameter_path_html string  = "/subPack_html_js"
//--------------------------------------
// color.go
//---------------------------------------
/**         ***  COLORS: got from https // www.dolthub.com/blog/2024-02-23-colors-in-golang/ ***
    var Reset   = "\033[0m" 
	var Red     = "\033[31m" 
	var Green   = "\033[32m" 
	var Yellow  = "\033[33m" 
	var Blue    = "\033[34m" 
	var Magenta = "\033[35m" 
	var Cyan    = "\033[36m" 
	var Gray    = "\033[37m" 
	var White   = "\033[97m"
**/

func red(     str1 string) string { return "\033[31m" + str1 +  "\033[0m" }
func green(   str1 string) string { return "\033[32m" + str1 +  "\033[0m" }
func yellow(  str1 string) string { return "\033[33m" + str1 +  "\033[0m" }
//func blue(  str1 string) string { return "\033[34m" + str1 +  "\033[0m" }
func magenta( str1 string) string { return "\033[35m" + str1 +  "\033[0m" }
func cyan(    str1 string) string { return "\033[36m" + str1 +  "\033[0m" }
//func gray(  str1 string) string { return "\033[37m" + str1 +  "\033[0m" }
//func white( str1 string) string { return "\033[97m" + str1 +  "\033[0m" }

//---------------------------------------------------------------------------------
var sw_stop bool = false
var errorMSG = ""; 

var sw_begin_ended = false     
var sw_HTML_ready  = false     
//var apiceInverso = `40`  //  in windows:  tasto Alt + 96 (tastierino numerico)
//---------------------------

var fileHtml = ""
var parmFolder = ""
var parmFile = ""
//var fileInp [] string
//var fileOut [] string
var outFolder = "folderOutProva"
var outFileName = "textFileOut.txt"

//---------------
// end of g01_declare.go
//-----------------------------------------
//--------------------------------
func main() {

	fmt.Println("\n======================\n         My Main()  INIZIO di mainPack \n===============================\n")
	fmt.Println(  red( appName + " - Main") )
	
	//fmt.Println( "\ncolori:", red("rosso"), green("verde"), yellow("giallo"),  magenta("magenta"), cyan("ciano") , "\n"  )  
	
	//---------------
	getPgmArgs()	
	
	//-----------------------------------
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	//  err := lorca.New("", "", 480, 320, args...) moved out of main so that ui is available outside main()
	if err != nil {
		fmt.Println( red( "errore in lorca "), err )  //  //log.Fatal(err)
	}
	defer ui.Close()
	
	bind_go_func_to_js()  //  bind inside is executed asynchronously after calling from js function (html/js are ready) 
	
	begin_GO_HTML_Talk();  // this function is  executed firstily before html/js is ready  
	
	// the following in main() is executed at the end when the browser is close 
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
		case <-sigc:
		case <-ui.Done():
	}
	fmt.Println("exiting") // log.Println("exiting...")
}
//-----------------------------------------

func endBegin(wh string) {

	//fmt.Println("func endBegin (", wh,")")
	if sw_stop { 
		fmt.Println("\nXXXXXXXX  error found XXXXXXXXXXXXXX\n"); 
	}	
	sw_begin_ended = true 		
}
//--------------------------------

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//-----------------------------------------

func getInt(x string) int {	
	y1, e1 := strconv.Atoi( x ) 
	if e1 == nil { 
		return y1
	} 
	y2, e2 := strconv.Atoi(  "0"+strings.TrimSpace(x)  ) 
	if e2 == nil {
		return y2
	} else {
		fmt.Println("error in getInt(",x,") ", e2) 
	}
	return 0
}
//--------------------------
func getPgmArgs() {	
    args:= os.Args[1:]
	ixHtml:= -1
	ixParm:= -1
	var parmArr []string
	for v:=0; v < len(args); v++ {
		switch args[v] {
			case "-html"   : ixHtml= v
			case "-parm"   : ixParm= v
		}		
	} 
	if ixHtml >= 0 { fileHtml = args[ixHtml+1] }
	if ixParm >= 0 {
		for v:=ixParm+1; v < len(args); v++ {
			if args[v][0:1] == "-" {break}
			parmArr  = append(parmArr, args[v])
		}
	}
	if len(parmArr) > 0 { parmFolder = parmArr[0] }
	if len(parmArr) > 1 { parmFile   = parmArr[1] }
	
	fmt.Println("fileHtml = ", fileHtml)
	fmt.Println("parm     = ", parmFolder, " ", parmFile) 	
	
} // end of getPgmArgs
//-----------------------------------------------------------
//end of g00_main.go
//------------------------------------------------------

//------------------------------------------------
//g03_html_env.go	
//--------------------------------------------------------

const screen_perc_scale int = 80; 

var scrX, scrY int = getScreenXY();
//------------------------------

var ui, err = lorca.New("", "", scrX, scrY); // crea ambiente html e javascript  // if height and width set to more then maximum (eg. 2000, 2000), it seems it works  


//---------------------
func getScreenXY() (int, int) {
	
	// use ==>  var x, y int = getScreenXY();
	
	var width  int = int(win.GetSystemMetrics(win.SM_CXSCREEN));
	var height int = int(win.GetSystemMetrics(win.SM_CYSCREEN));
	fmt.Println("scr3en1 ", width , " x " , height); 
	if width == 0 || height == 0 {
		//fmt.Println( "errore" )
		return 2000, 2000; 
	}	
	width  = width  - 20;  // subtraction to make room for any decorations 
	height = height - 40;  // subtraction to make room for any decorations 
	
	width  = int( width  * screen_perc_scale / 100 ) 
	height = int( height * screen_perc_scale / 100 ) 
	fmt.Println("screen2 ", width , " x " , height); 
	return width, height
	
} // end of getScreenXY()
//----------------------------------	

func begin_GO_HTML_Talk() { 	
	fmt.Println("func begin_GO_HTML_Talk"); 
	setHtmlEnv();	
}
//---------------

//------------------------------------
func setHtmlEnv() {	
	fmt.Println("func setHtmlEnv:  start load html")
    // load file html 	
	
	var html_path = getCompleteHtmlPath( parameter_path_html ) 
	            
	fmt.Println("path html        = " + html_path)
	fmt.Println("fileHtml=",fileHtml)
	//ui.Load("file:///" + html_path + string(os.PathSeparator) + htmlFile ); 
	ui.Load(fileHtml); 
	
	fmt.Println("\n", "func setHtmlEnv: wait for html ( javascript function js_call_go()", "\n")  
	
} // end of setHtmlEnv
//--------------------------------------------------------
//-------------------------
func getCompleteHtmlPath( path_html string) string {
	
	//curDir    := "D:/ANTONIO/K_L_M_N/LINGUAGGI/GO/_WORDS_BY_FREQUENCE/WbF_prova1_input_piccolo
	 
	curDir, err := os.Getwd()
    if err != nil {
		fmt.Println("setHtmlEnv() 3 err=", err )
        //log.Fatal(err)
    }	
				
	fmt.Println("curDir           = " + curDir ); 
	
	curDirBack  := curDir
	k1:= strings.LastIndex(curDir, "/") 
	k2:= strings.LastIndex(curDir, "\\") 
	if k2 > k1 { k1 = k2 } 
	curDirBack = curDir[0:k1] 	
	
	var newPath string = ""
	if strings.Index(path_html,":") > 0 {
		newPath = path_html
	} else if path_html[0:2] == ".." {
		newPath = curDirBack  + path_html[2:] 
	} else {
		newPath = curDir + path_html
	}
	return newPath 
} 
//------------------------
func putFileError( msg1, inpFile string) {
	err1:= `document.getElementById("id_startwait").innerHTML = '<br><br> <span style="color:red;">§msg1§</span> <span style="color:blue;">§inpFile§</span>';` ; 		
	err1 = strings.ReplaceAll( err1, "§msg1§", msg1 ); 	 
	err1 = strings.ReplaceAll( err1, "§inpFile§", inpFile); 	
	ui.Eval( err1 );	
}   

//-----------------------------------
// end of g03_html_env.go	
//--------------------------------------------------------

//-------------------------------------------
func bind_go_func_to_js() {		

		ui.Bind("goFunCalledByJS_mgr", 
			func(goFunc string, js_function string, var1 string, var2 string, var3 string, var4 string, var5 string ) { 				
				switch goFunc {
					case "funCalledByJs_0_html_is_ready":    funCalledByJs_0_html_is_ready(  var1, js_function)  
					case "funCalledByJs_1_readFile"     :    funCalledByJs_1_readFile( var1, var2, js_function) 
					case "funCalledByJs_2_writeFile"    :    funCalledByJs_2_writeFile(var1, var2, var3, js_function)  			
					default: 
						fmt.Println( red("error in bind_go_func_to_js"), " the string '", green(goFunc) , "' is unknown, It cannot be to any go func")    
						return;
				}
			} )
			
		//--------------------------------------		
		
		
		return 
}

//---------------------------------------------
func funCalledByJs_0_html_is_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func bind_go_passToJs_html_is_ready () " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   ", green("html has been loaded"), "   XXXXXXXXXXXX")
	
	begin() 
	
	fmt.Println( green( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" + 
		"xxxxxxxxxxxxxx you can use the tool xxxxxxxxxxxxxxxxxx\n"  + 
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n\n"  ) ) 
	
} // end of bind_go_passToJs_html_is_ready
 
//--------------------------------------

func funCalledByJs_1_readFile(inpFolder string, inpTxtFile string, js_function string) {
	sw_stop = false
	bytesPerRow:= 100
    righe := rowListFromFile( inpFolder, inpTxtFile, "file input", "read_file", bytesPerRow)  
	if sw_stop { return }
	
	fmt.Println("\nletti in ", inpTxtFile  + " " , len(righe) , " righe")   
	
	go_exec_js_function( js_function, strings.Join(righe, "\n"));
	
} // end of funCalledByJs_1_readFile
//-----------------------------------------

func funCalledByJs_2_writeFile(outFolder string, outTxtFile string, outStr string,  js_function string) {
	sw_stop = false
	o_listLines:= strings.Split(outStr,"\n")	
	writeList2( outFolder, outTxtFile, o_listLines) 
	if sw_stop == false {
		go_exec_js_function( js_function, fmt.Sprint("scritte ", len(o_listLines), " righe") )		
	} 
	
} // end of funCalledByJs_2_writeFile	

//--------------------------------------



//-----------
//g04_begin.go
//------------------------------

func begin() { 	

	if sw_stop { 	
		fmt.Println("UI is ready, but run stopped because of some error")		
		
		go_exec_js_function("UI is ready, but run stopped because of some error","")	
		return
		
	} else {		
		fmt.Println(cyan("\nREADY"), "\n") 
		
		go_exec_js_function("js_go_ready",parmFolder + "," + parmFile)	
	}
	
}// end of begin	

//--------------------------------

// end g04_begin.go
//-------------------------------------- 

//------------------------
//g05_go_exec_js_function.go
//----------------------------------------------------------------
func go_exec_js_function(js_function0 string, inpstr string) {
	sw_stop = false
	var goFunc string 
 	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		goFunc = strings.ReplaceAll(details.Name(), "main.","")
	} else {
		goFunc=""
	}
	js_fun        := strings.Split( (js_function0 + ",,,,") ,",") 	
	js_function   := strings.TrimSpace( js_fun[0] )
	if js_function == ""  { return }
	jsInpFunction := strings.TrimSpace( js_fun[1] )
	
	js_parm:=""
	k1:= strings.Index(js_function, "(") 
	if k1 > 0 {
		js_parm     = strings.ReplaceAll(  js_function[k1+1:], ")","")			
		js_function = strings.TrimSpace(js_function[0:k1] )
	} 	
	
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by " "   	
	*/
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	   	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 
	
	evalStr := fmt.Sprintf( "%s(`%s`,`%s`,`%s`,`%s`);",  js_function, inpstr, js_parm, "js=" + jsInpFunction, "go=" + goFunc ) ; 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function

//----------------------------------------------------------------

// end of g05_go_exec_js_function.go
//----------------------------------------------------

//--------------------------------------
//g06_read_any_text_file.go
//-----------------------------------------
func test_folder_exist( myDir string, msg2 string) {
    _, err := os.Open( myDir )
    if err != nil {
		msg0:= `la cartella <span style="color:blue;">` + myDir + "</span> non esiste"			
		msg1:= `la cartella ` + myDir + "</span> non esiste"				
					
		errorMSG = `<br><br> <span style="color:red;">` + msg0 + `</span>` +  
			`<br><span style="font-size:0.7em;">(func ` + msg2 	 + ")" + `</span>` 		
		showErrMsg(errorMSG, msg1, msg2 )	
		
		sw_stop = true 
		return		
    }
} // end of test_folder_exist	
//------------------

//------------------
func getFileByteSize( path1 string,   fileName string) int {
	path2:=""
	if path1 != "" {
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	fileInfo, _ := os.Stat( fileN )  
	
	//fmt.Println("getFileByteSize fileN=", fileN, " fileInfo = ", fileInfo) 
	if fileInfo == nil { return 0 }
	return int( fileInfo.Size() )
} // end of getFileByteSize

//--------------
func myOpenRead( path1 string,   fileName string,   descr string,  func1    string) (*os.File, int) {
	path2:="";
	path10:=""
	if path1 != "" {
		path10 = " in " + cyan(path1)
		path2 = path1 + string(os.PathSeparator) 
	} 	
	fileN := path2 + fileName 
	
	fmt.Println("\n" + yellow("open file"),  green(fileName) , path10 )
	
	sizeByte:= getFileByteSize(path1,fileName)
	readFile, err := os.Open( fileN )  
    if err == nil {				
		fmt.Println( "\t", "size: ", sizeByte, " bytes" )	
		return readFile, sizeByte
	}
	msg1_Js:= `il file "` + fileN + `" (` + descr + " " + func1 + ")" + " non esiste"
		
	errorMSG = `<br><br>il file ` + 
				`<span style="font-size:0.7em;color:black;">(`	+ descr + `)</span>` +
				`<br><span style="color:blue;" >` + fileName + `</span>`	+ 				
				`<br><span style="font-size:0.7em;color:red;">`	+ "non esiste" 	+ `</span>` +				
				`<br><span style="font-size:0.7em; color:black;">nella cartella ` + path2    + `</span>` 
				
	showErrMsg2(errorMSG, msg1_Js)	
	
	return readFile, 0		
	
} // end of myOpenRead


//----------------------------------

func rowListFromFile( path1 string, fileName string, descr string, func1 string, bytesPerRow int ) []string { 
	
	file, sizeByte := myOpenRead( path1, fileName, descr, func1 )  
	if file == nil {			
		sw_stop = true 
		return nil
	} 
	numEleMax:= int( sizeByte / bytesPerRow ); 
	if numEleMax < 10 {numEleMax=10}
	
	fmt.Println("    allocate for a maximum of ", numEleMax, " rows (assumed ", bytesPerRow, " bytes per row as average)" )
	
	retRowL := make( [] string, 0, numEleMax)	
	
	r := bufio.NewReader(file)
	for {
	  line, _, err := r.ReadLine()
	  if err != nil {
		if err == io.EOF {
			break
		}
		break
	  }	 
	  retRowL = append( retRowL, string(line) ) 	
	}
	defer file.Close()
	
	fmt.Println("letto file " , fileName, "  num lines=", len(retRowL) )
	
	return retRowL 
	
} // end of rowListFromFile	

//-------------------------------------
// end of g06_read_any_text_file.go
//-----------------------------------------

//-----------------------------------------------
//g07_write_file.go
//------------------------------------------------

//----------------------
func writeList2( outFolder string, outTxtFile string, lines []string)  {
	test_folder_exist(outFolder, "writeList2")
	if sw_stop {return}	
	fileName := outFolder + "/" + outTxtFile
	
	// create file
    f, err := os.Create( fileName )
    if err != nil {
        fmt.Println( red("error")," in writeList file=", outTxtFile, "\n\t" , err ) //  log.Fatal(err)
		fmt.Println("path completo=", fileName) 
    }
    // remember to close the file
    defer f.Close()

    // create new buffer
    buffer := bufio.NewWriter(f)

    for _, line := range lines {
        _, err := buffer.WriteString(line + "\n")
        if err != nil {
		   descr:= 	  fmt.Sprintln( " in buffer.WriteString file=", fileName,"\n\t" , err )	
           fmt.Println( red("error"), " in buffer.WriteString file=", fileName,"\n\t" , err ) //log.Fatal(err)
		   errorMSG = `<br><br>il file ` + 
				`<span style="font-size:0.7em;color:black;">(`	+ descr + `)</span>` +
				`<br><span style="color:blue;" >` + fileName + `</span>`	
			showErrMsg2(errorMSG, "write line")			
		    return
        }
    }
    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
        fmt.Println( red("error"), " in buffer.Flush()cls file=", fileName,"\n\t" , err ) //  log.Fatal(err)
		descr:=  fmt.Sprintln("in buffer.Flush()cls file=", fileName,"\n\t" , err )
		errorMSG = `<br><br>il file ` + 
				`<span style="font-size:0.7em;color:black;">(`	+ descr + `)</span>` +
				`<br><span style="color:blue;" >` + fileName + `</span>`	
		showErrMsg2(errorMSG, "writeList")			
		return
    } else {
	  fmt.Println( green("scritte") , len(lines), " righe nel file ", fileName ) 
	}
} 
//-----------------------------------------------
// end of g07_write_file.go
//------------------------------------------------


//-------------------------------------
func showErrMsg(errorMSG0 string, msg1 string, func1 string) {
	errorMSG = strings.ReplaceAll( errorMSG0, "\n"," ") 
	fmt.Println(msg1, " func ", func1)
	go_exec_js_function( "js_go_setError", errorMSG ); 	
}
//--------------------------------
func showErrMsg2(errorMSG0 string, msg1 string) {
	errorMSG = strings.ReplaceAll( errorMSG0, "\n"," ") 
	fmt.Println(msg1)
	go_exec_js_function( "js_go_setError", errorMSG ); 	
}
//------------------------------

