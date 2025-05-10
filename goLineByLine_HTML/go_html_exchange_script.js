"use strict";
/*  
gohtml_sample - Antonio Cigna 2025
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//--------------------------------------------------------   
//  ... html_EXCHANGE_SCRIPT
//-----------------------------------------------
let file_LOCALSTOR_folder = "";
let file_LOCALSTOR_name   = "";
let file_LOCALSTOR_str    = "";

var parm_title=""
var parm_folder 		= "./"; 
var parm_file   		= "mioFile.csv";      // file formato csv
var parm_separ  		= "|";  			  // |   \t \n  
var parm_col_lang1 	= 0;                  // numero di colonna con frasi in lingua originale 
var parm_col_lang2 	= 1;                  // numero di colonna con frasi tradotte 
var parm_col_lang3 	= -1;                 // numero di colonna con informaz. supplementari es. paradigma verbo, plurale dei nomi 
var parm_language1_2 = "de_DE";           // de_DE  sigla della lingua l'originale      (es. de_DE, en_EN )
var parm_language2_2 = "it_IT";           // de_DE  sigla della lingua per la traduzione(es. it_IT )
var parm_sortTextLength = "false";        // if true sort input by length of the text in original language  
var sw_parm_sortTextLength = false;
//-----------------------------------------------------------------------
var runByGo = true;	
//-----------------------------------------
function js_call_go_on_html_body_load() {  // called by  html page body onload
	var msg1 = "html loaded"; 
	console.log("html is ready"); 
	
	/*
	goFunCalledByJS_mgr is not a js function, but a link to a go function (see go func 'bind_go_func_to_js()' ) 
		parameters: 1 : the actual go func name to run 
		            2 : the js fuction to be run when the go func has ended
					3,4,5,6,7: the parameters for the go func ( all of them must be passed as strings )  	
	*/
	goFunCalledByJS_mgr( "funCalledByJs_0_html_is_ready", "", msg1, "2","3","4","5" ) ; 
	
}

//-------------------------------------		

function js_go_ready(inpStr) {
	// this function is run by go
	//console.log("go is ready ", inpStr); 
	
	document.getElementById("id_start001").style.display = "none";
	//document.getElementById("id_myPage01").style.display = "flex"; 
	
	//????leggiFileZERO( inpStr ); 
	
	var str1 = inpStr.split("::LOCALSTOR=");
	if (str1.length < 4)  { str1.push(""); str1.push(); str1.push();str1.push();}
	var parmStr=""; 
	
	[ parmStr, file_LOCALSTOR_folder,   file_LOCALSTOR_name,  file_LOCALSTOR_str ] = str1;

	console.log("%cjs_go_ready","color:blue;")
	/**
	console.log("parmfile str=\n\t", str1[0],"\nfile_LOCALSTOR_folder", file_LOCALSTOR_folder, "\nfile_LOCALSTOR_name=", file_LOCALSTOR_name, 
		"\nfile_LOCALSTOR_str=", file_LOCALSTOR_str);
	**/	
	if (file_LOCALSTOR_str == "") {  file_LOCALSTOR_str = JSON.stringify("")}	
		
	localStorage.setItem(cbc_LOCALSTOR_key,  file_LOCALSTOR_str)
	
	js_go_1_returnFileZero( parmStr );
	
	
} // end of js_go_ready
//--------------------------------------------------------------
function TOGLIleggiFileZERO(inpStr) {
	console.log("LEGGI FILE ZERO XXXXXXXXX" , "  parm=>", inpStr); 
	var parmFile = inpStr.split(",")
	js_go_setError("");
	let inpFolder  = parmFile[0];
	let inpTxtFile = parmFile[1];  //  "parametri.txt";	
	console.log("parmFile=", parmFile)
	goFunCalledByJS_mgr( "funCalledByJs_1_readFile", "js_go_1_returnFileZero", inpFolder, inpTxtFile,"","","");  	
}
//-----------------
function js_go_1_returnFileZero(inpStr) {
	//console.log("return File ZERO ==>" , inpStr) 
	parm_title=""
	parm_folder 		= "./"; 
	parm_file   		= "mioFile.csv";   // file formato csv
	parm_separ  		= "|";  		   // {tab}  significa \t cioè tabulazione 
	parm_col_lang1 	= 0;                   // numero di colonna con frasi in lingua originale 
	parm_col_lang2 	= 1;                   // numero di colonna con frasi tradotte 
	parm_col_lang3 	= -1;                  // numero di colonna informaz. suppl. es. plurale nomi, paradigma verbi 
	parm_language1_2 = "de_DE";            // de_DE  sigla della lingua l'originale      (es. de_DE, en_EN )
	parm_language2_2 = "it_IT";            // de_DE  sigla della lingua per la traduzione(es. it_IT )
	parm_sortTextLength = "false";         // if true sort input by length of the text in original language  
	sw_parm_sortTextLength = false;
	var righe = inpStr.split("\n");
	
	for(var v=0; v < righe.length; v++) {
		var cols = righe[v].split("=");	
		if (cols.length < 2) continue;
		var key = cols[0].trim().toLowerCase();
		var valueC = cols[1].split("//");
		var value  = valueC[0].trim(); 
		//console.log("key=" + key +", value=" + value + "<==");  
		if	 (key == "title") 			parm_title  = value; 	
		else if (key == "folder") 		parm_folder = value; 	
		else if (key == "file")   		parm_file   = value; 	
		else if (key == "separ")  		parm_separ  = value;
		/**
		else if (key == "col_lang1")   	parm_col_lang1   = value; 	
		else if (key == "col_lang2") 	parm_col_lang2   = value;
		else if (key == "col_lang3") 	parm_col_lang3   = value;
		**/
		else if (key == "language1_2" ) parm_language1_2 = value; 	
		else if (key == "language2_2" ) parm_language2_2 = value; 
		else if (key == "sorttextlength") sw_parm_sortTextLength = (value == "true");
		else if (key.substr(0,8) == "col_lang") {
				var valueNum=-1;				
				try{ 
					valueNum = parseInt(value); 
					if (isNaN(valueNum)) valueNum=-1;
				} catch(err) {}; 
				if 		(key == "col_lang1")   	parm_col_lang1   = valueNum; 	
				else if (key == "col_lang2") 	parm_col_lang2   = valueNum;
				else if (key == "col_lang3") 	parm_col_lang3   = valueNum;
		}		
	} 
	
	/**
	console.log("\n",  
		"title", 		" \t", parm_title, 	"\n",  
		"folder",  		" \t", parm_folder, 	"\n",  
		"file",  		" \t", parm_file , 	"\n", 
		"separ",  		" \t", parm_separ , 	"\n",  
		"col_lang1",  	" \t", parm_col_lang1 , 	"\n", 
		"col_lang2",  	" \t", parm_col_lang2 , 	"\n", 
		"col_lang3",  	" \t", parm_col_lang3 , 	"\n", 
		"language1_2",  " \t", parm_language1_2, "\n",
		"language2_2",  " \t", parm_language2_2, "\n",
		"sw_parm_sortTextLength",  " \t", sw_parm_sortTextLength, "\n",
		""); 
	**/
	
	if (parm_separ == "{tab}") {parm_separ = "\t"; }	
		
	if (parm_title != "") {
		document.getElementById("titleB2" ).innerHTML = parm_title;
		document.getElementById("id_title").innerHTML = parm_title;
	}
	goFunCalledByJS_mgr( "funCalledByJs_1_readFile", "js_go_1_returnReadText", parm_folder, parm_file,"","","");  
	
	
	// get string read by go program 
	/***
	let ele1 = document.getElementById("inpFileContent");
	ele1.value = inpStr; 
	ele1.parentElement.style.display = "block"; 	
	***/
}	

//--------------------------------------------------------------
function leggiFileUno() {
	console.log("LEGGI FILE UNO XXXXXXXXX"); 
	
	js_go_setError("");
	let inpFolder  = "D:/ANTONIO/Anki_inputSchede/UnaProva_con_GoHTML_readWrite/fileCSV";
	let inpTxtFile = "ANKI_tedescoNet_A1_antonio.csv";	
	goFunCalledByJS_mgr( "funCalledByJs_1_readFile", "js_go_1_returnReadText", inpFolder, inpTxtFile,"","","");  	
}
//------------------
function onclick_1_readFile( id_inpFolder, id_inpTxtFile) {
	
	// ask go program to read file
		
	js_go_setError("");
	let inpFolder  = document.getElementById(id_inpFolder ).value;
	let inpTxtFile = document.getElementById(id_inpTxtFile).value;	
	goFunCalledByJS_mgr( "funCalledByJs_1_readFile", "js_go_1_returnReadText", inpFolder, inpTxtFile,"","","");  	
} // end of onclick_1_readFile
//-----------------
function js_go_1_returnReadText(inpStr) {
	
	// get string read by go program 
	/***
	let ele1 = document.getElementById("inpFileContent");
	ele1.value = inpStr; 
	ele1.parentElement.style.display = "block"; 	
	***/
	var FILLER_num = 100000;
	var righe = inpStr.split("\n");
	var lenArr =[];
	var v;	
	var numInpRighe = righe.length; 
	let ele2 = document.getElementById("id_readMsg");
	ele2.innerHTML = "lette " + numInpRighe + " righe";
	var maxPrmCol=Math.max(parm_col_lang1,parm_col_lang2, parm_col_lang3);
	if (sw_parm_sortTextLength) {
		console.log("sort righe per lunghezza testo in lingua originale");
		for(v=0; v < numInpRighe; v++) {
			var cols = righe[v].split( parm_separ );
			var numCols=cols.length; 	
			var len1=0;	
			if ((parm_col_lang1 >=0) && (parm_col_lang1 < numCols))  {
				 len1 = cols[parm_col_lang1].length;
			}	
			lenArr.push( len1*FILLER_num + v );
		}
		lenArr.sort(function(a, b){return a - b});
			
	} else {
		for(v=0; v < numInpRighe; v++) {
			lenArr.push(FILLER_num+v);
		}
	}	
	//--------------------------	
	var inpTab = "";
	for(var v00=0; v00 < lenArr.length; v00++) {
		var len1 = lenArr[v00];	
		v = len1%FILLER_num;
		if (v == len1) continue; // lunghezza del testo è zero 
		var cols = righe[v].split( parm_separ );
		var numCols=cols.length; 		
		if ((parm_col_lang1 >=0) && (parm_col_lang1 < numCols))  inpTab +=  "<tr><td>" + cols[parm_col_lang1]; else  inpTab +=  "<tr><td>"; 	
		if ((parm_col_lang2 >=0) && (parm_col_lang2 < numCols))  inpTab += "</td><td>" + cols[parm_col_lang2]; else  inpTab += "</td><td>"; 
		if ((parm_col_lang3 >=0) && (parm_col_lang3 < numCols))  inpTab += "</td><td>" + cols[parm_col_lang3]+ "</td></tr>\n"; else  inpTab += "</td><td>" + "</td></tr>\n"; 
	} 
	document.getElementById("inp1").innerHTML = "\n" + inpTab + "\n";
		
	begin_lbl2(); 	
	
} // end of js_go_1_returnReadText
//--------------------------------------------------

function TOGLIonclick_2_writeFile( id_outFolder, id_outTxtFile, id_outStr) {
	
	// ask go program to write string to a file
	
	js_go_setError("");
	let outFolder  = document.getElementById(id_outFolder ).value;
	let outTxtFile = document.getElementById(id_outTxtFile).value;
	let outStr     = document.getElementById(id_outStr    ).value;
	
	
	console.log("da pag html ", "outFolder =", outFolder);	
	console.log("da pag html ", "outTxtFile=", outTxtFile);	
	console.log("da pag html ", "outStr    =", outStr); 
	
	goFunCalledByJS_mgr( "funCalledByJs_2_writeFile","js_go_2_returnWriteFile", outFolder, outTxtFile, outStr,"","");  	
	
} // end of onclick_2_writeFile
//-------------------------------------
function js_go_2_returnWriteFile(msg1) {		
	// return message with the write operation result
	
	/**
	let ele1 = document.getElementById("id_writeMsg");
	ele1.innerHTML = msg1;
	**/
}
//--------------------------------------------------	

function js_go_setError(msg1) {
	// error message got by the go program 
	let ele1 = document.getElementById("id_msgErr");
	ele1.innerHTML = msg1; 
}
//--------------------

//------------------------------------------------------------
const LOCALSTOR_key = document.title;
//-------------------------------------

let LOCALSTOR_content_array = [];    // contenuto da salvare su localstor  
//---------------------------------------

//------------------------------------------------
 function localStorage_setItemFile(cbc_LOCALSTOR_key, LOCALSTOR_content_array) {
	
	// Save Data to Local Storage
	
	file_LOCALSTOR_str = JSON.stringify( LOCALSTOR_content_array ); 
	
	/*
	console.log("%clocalStorage_setItemFile() cbc_LOCALSTOR_key=" + cbc_LOCALSTOR_key, "color:blue; font-weight:bold;font-size:1.2em;") 
	
	console.log("localStorage_setItemFile() --> goFunCalledByJS_mgr() --> funCalledByJs_2_writeFile --> \n\tlocaStor folder=",  file_LOCALSTOR_folder,
		"\n\tlocalStor  filename=",  file_LOCALSTOR_name, 
		"\n\tlocalStor  str 	=", file_LOCALSTOR_str,  "4","5"); 
	*/
  //goFunCalledByJS_mgr( "funCalledByJs_localStor_setItem" , "localStorage_print", LOCALSTOR_key, valStr1,  "3","4","5");  
  goFunCalledByJS_mgr( "funCalledByJs_2_writeFile" , "js_go_2_returnLocalStorWrite",  file_LOCALSTOR_folder,file_LOCALSTOR_name, file_LOCALSTOR_str,  "4","5");   
 
 }	// end of localStorage_setItemFile
 //-------------------------------------------------------------------
 function js_go_2_returnLocalStorWrite(inpStr) {
	 //console.log("return local storage write msg=", inpStr)
	 
 }
 //-------------------------------------------------
function localStorage_setData() {	
	if (go_is_ready == false) { return; }
	// Save Data to Local Storage
	
	var valStr1 = JSON.stringify( LOCALSTOR_content_array ); 
	
	//window.localStorage.setItem( LOCALSTOR_key, JSON.stringify( LOCALSTOR_content_array ) );
	
  //goFunCalledByJS_mgr( "funCalledByJs_localStor_setItem" , "localStorage_print", LOCALSTOR_key, valStr1,  "3","4","5");   
	goFunCalledByJS_mgr( "funCalledByJs_localStor_setItem" , "",                   LOCALSTOR_key, valStr1,  "3","4","5");   
	
} // end of localStorage_setData
//------------------------------------------------------------------------------
function localStorage_removeKey(key) {
	//Remove Data from Local Storage
	if (key == null) return; 
	return; 
    cambiaGo_window.localStorage.removeItem(key);
}
//------------------------------------
function localStorage_list( inpKey ) {
	// se inpKey manca, lista tutte le chiavi, se esiste solo la chiave richiesta 
	for (var i = 0, len = cambiaGo_window.localStorage.length; i < len; i++) {
		var key = cambiaGo_window.localStorage.key(i);
		if (inpKey) { if (key != inpKey) continue; }
		var value = cambiaGo_window.localStorage[key];
		console.log("window.localStorage key=" + key + " ==> " + value.replaceAll( ",[" ,    ",\n\t[") );
	}
		
} // end of list_localStorageItems()
//------------------------------------
function localStorage_keyList() {
	goFunCalledByJS_mgr( "funCalledByJs_localStor_print", "","","","","");  
	/**
	for (var i = 0, len = cambiaGo_window.localStorage.length; i < len; i++) {
		console.log( "window.localStorage key=", cambiaGo_window.localStorage.key(i)); 
	}
	**/		
} // end of list_localStorageItems()
//------------------------------------
function localStorage_print() {
	console.log("localStorage_print () "); 
  //goFunCalledByJS_mgr( "funCalledByJs_localStor_getItem" , "localStorage_getData_go2_back", LOCALSTOR_key, "2","3","4","5");   
	goFunCalledByJS_mgr( "funCalledByJs_localStor_print",    "",                              "",            "","","","");  
	/**
	for (var i = 0, len = cambiaGo_window.localStorage.length; i < len; i++) {
		console.log( "window.localStorage key=", cambiaGo_window.localStorage.key(i)); 
	}
	**/		
} // end of list_localStorageItems()
//-----------------------------------------------
//   end of localStorage_routines.js
//-----------------------------------------------