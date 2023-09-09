package main
import (
    "flag"
    "fmt"
    "time"
    "io/ioutil"
    "strings"
    "path/filepath"
)
var ogonki = map[string]string{
	 //-------------------------------------------------------
	 //Albański | Albanian (sq) 
		"sq": "çéëÇËÉ",
	 //Baskijksi | Basque (eu)
		"eu": "áéíóúüñÁÉÍÓÚÜÑ",
	 //Bretoński | Breton (br) 
		"br": "âãäæçéèêëêôöœûüùñŷýÿìïîŕśÂÄÃÆÇÉÈÊËÔÖŒÑÙÛÜŶÌÏÎŔŚÝŸ",
	 //Kataloński | Catalan (ca) 
		"ca": "áàà́èéè́íìòóò́ïúüùýÁÀÀ́ÈÉÈ́ÏÍÌÒÒ́ÓÚÜÙÝ",
	 //Chorwacki | Croatian (hr)
		"hr": "čćđšžČĆĐŠŽ",
	 //Czeski | Czech (cz) 
		"cz":"áčďéěíňóřšťúůýžÁČĎÉĚÍŇÓŘŠŤÚŮÝŽ",
	 //Duński | Danish (dk) 
		"dk":"æøåÆØÅ",
	 //Niderlandzki (Holenderski) | Dutch/Netherlands (nl) 
		"nl": "äáéëïíóöüúÄÁÉËÍÏÓÖÜÚ",
     //Estoński | Estonian (et) 
		"et": "äöõüšžÄÖÕÜŠŽ",
     //Fiński | Finnish (fi) 
		"fi":"äåöÄÅÖ",
     //Francuski | French (fr)
		"fr":"àâäçéèêëîïôûùüÿœæÀÄÂÇÉÈÊËÎÏÔÛÙÜŸŒÆ",
	 //Galicyjski/Galisyjski | Galician (gl)
		"gl": "áéíóúñüÁÉÍÓÚÑÜ",
     //Węgierski | Hungarian (hu)
		"hu":"áéíóöőúüűÁÉÍÓÖŐÚÜŰ",
     //Islandzki | Icelandic (is)
		"is" : "áðéíóöúýþæÁÐÉÍÓÖÚÝÞÆ",
     //Irlandzki | Irish (ga) 
		"ga": "áéíóúàèéìòóùúċḃḋḟġṁṗṡṫāēīíōūǽǿÁÉÍÓÚÀÈÍÌÒÓÙÚĊḂḊḞĠṀṖṠṪĀĒĪŌŪǼǾ",
     //Włoski | Italian (it) 
		"it":"àèéìòóùüÀÈÉÌÒÓÙÜ",
     //Łotewski | Latvian (lv)
		"lv": "āčēģīķļņšūžĀČĒĢĪĶĻŅŠŪŽ",
     //Litewski | Lithuanian (lt)
		"lt": "ąčęėįį̇šųūžĄČĘĖĮĮ̇ŠŲŪŽ",
	 //Maltański/Maltyjski | Maltese/Malta (mt)
		"mt": "ċġħżĊĠĦŻ",
     //Norweski | Norwegian (no)
		"no":"áéíæøóåúýÆØÁÅÉÍÓÚÝ",
     //Polski | Polish (pl)
		"pl":"ąćęłńóśźżĄĆĘŁŃÓŚŹŻ",
     //Rumuński | Romanian (ro)
		"ro": "ăâîșțţĂÂÎȘȚŢ",
     //Słowacki | Slovakia (sk)
		"sk": "áäčďéíĺľňóôŕšťúýžÁÄČĎÉÍĹĽŇÓÔŔŠŤÚÝŽ",
     //Słoweński | Slovenian (sl) 
		"sl": "ą̊ą̇ąáčćđéęíįĺľńŕšśóúųžźĄ̊Ą̇ĄÁČĆĐÉĘÍĮĹĽŃŔŠŚÓÚŲŽŹ",
     //Hiszpański | Spanish (es)
		"es":"áéíóúüñÁÉÍÓÚÜÑ",
     //Turecki | Turkish (tr)
		"tr": "şığüöçŞİĞÜÖÇ",
     //Wietnamski | Vietnam (vi) 
		"vi": "áàãảạăắằẵẳặâấầẫẩậđéèẽẻẹêếềễểệíìĩỉịóòõỏọôốồỗổộơớờỡởợúùũủụưứừữửựýỳỹỷỵÁÀÃẢẠĂẮẰẴẲẶÂẤẦẪẨẬĐÉÈẼẺẸÊẾỀỄỂỆÍÌĨỈỊÓÒÕỎỌÔỐỒỖỔỘƠỚỜỠỞỢÚÙŨỦỤƯỨỪỮỬỰÝỲỸỶỴ",
     //Walijski | Welsh (cy) 
		"cy": "äâêëîïôûŵŷÂÄÊËÎÏÔÛŴŶ",
     //Bośniacki | Bosnia and Herzegovina (ba)
		"ba": "ćčđšžıĆČĐŠŽ",
     //Grecki | Greek (el)
		"el":"άαβγδεέζηθικλμνξοπρστήίϊΐυύΰϋόφχψωώΆΓΔΘΛΈΞΠΉΊΪΣΌΦΨΏΩΎΫ",
	 //Serbski | Serbian (sr)
		"sr": "čćđšžüČĆĐŠŽÜ", 
     //Łaciński | Latin (la)
		"la": "áąāéęēīōūȳǖǘǚǜċćġħǰŝźćģķļłņńŗśźżǎěǐǒóúǔǖǘǚǜǧğıíįķļŉǫřśŝţťūǔůẏýžæœÁĄĀĒÉĘĪŌÓŪȲǕǗǙǛĊĆĠĦŜŹĆĢĶĻŁŅŃŖŚŹŻǍĚÍǏǑǓǕǗǙǛǦĞİĮĶĻŊǪŘŚŜŢŤÚŪǓŮẎÝŽÆŒ",
     //Kurdyjski | Kurdish (ku)
		"ku": "âáçéêğîíışţôûýÁÂÇÉÊĞÎÍŞŢÔÛÝ",
     //Farerski | Faroese (fo)
		"fo": "áðéíóúýæøÁÐÉÍÓÚÝÆǾØ",
     //Fryzyjski | Frisian (fy)
		"fy": "áäâêéëïîíôóöøûúüÿýÁÄÂÊÉËÏÎÍÖÓÔØÚÜÛÝŸ",
     //Górnołużycki | Upper Sorbian (hsb)
		"hsb": "ǎáćčďđěéęȟíj́ḱĺłḿńŉóöőǫǒřŕśšŭůw̌ŵźžӯǍÁĆČĎĐĚÉȞÍJ̌J́ḰĹŁḾŃŊÓǪǑӦŐǓŬŮŔŘŚŠŴW̌ӮȲŹŽ",
     //Kaszubski | Kashubian (csb)
		"csb": "ąâãćêěëęèéé̄łïńôòóó̄ò́ò̂ôśś́üùûżźĄÂÃĆÈÊĚËĘÉÉ̄ŁÏŃÔÒÒ́Ò̂ÓÓ̄ÔÜÙÛŚŚ́ŻŹ",
     //Kosowski | Kosovan (sq)
		"sq-kosovo": "çëë̈f̈g̈šïx̌ÿžÇËË̈F̈G̈ŠÏŽX̌ŸŽ", 
     //Kornijski | Cornish (kw)
		"kw": "āâēêëôōîīïûūŵẃŷȳÿǣǽĀÂĒÊËÔŌÎĪÏÛŪŴẂŸȲŶǢǼ",
     //Krymskotatarski | Crimean Tatar (crh) 
		"crh": "âäƀçèéëğıìíñöşûüÂÄÇÈÉËĞÌÍÑÖŞÛÜ",
     //Luksemburski | Luxembourgish (lb)
		"lb": "âäéëîïôöûüŷÂÄÉËÎÏÔÖÛÜŶ",
     //Mołdawski    | Moldovan (md)
		"md": "ăâîșşțţĂÂÎȘŞȚŢ", 
	 //Romski/Cygański | Romani (rom) 
		"rom": "ạảấầẩẫậắằẳẵặáa̧ǎăȃâãǣćčđȩéěĕȇêẹ̄ẻếềểễệêëîỉịíi̧ǐĭȋîị̄ļłňỏốồổỗộớờởỡợóo̧ǒŏȏôọ̄ǿöőřśšţťüůűùúu̧ǔŭȗûụ̄ụủứừửữựýźžẠÃẢẤẦÁẨẪẬẮẰẲẴẶA̧ǍĂȂÂÃǢĆČĐÉȨĚĔÊẸ̄ẺẾỀỂỄỆĚËỈÎỊI̧ÍǏĬȊÎỊ̄ĻŁŇỎỐỒỔỖỘỚỜỞỠỢÓO̧ǑŎȎÔỌ̄ǾÖŐŘŚŠŤŢŮÜŰÚU̧ǓŬUÙÛỤ̄ỤỦỨỪỬỮỰÝŹŽ",
	 //Szwajcarski niemiecki | Swiss German (gsw)
		"gsw": "àäëïöüÿÀÄËÏÖÜŸ",
	 //Szwajcarski włoski | Swiss Italian (ch-it) 
		"ch-it": "àèéìòóùÀÈÉÌÒÓÙ",
     //Łemkowski | Lemko (x-lmk) 
		"x-lmk": "åäǎąćďęěįǐĺľłńňóǒöøŕřśšťųüůýźžÅÄǍĄĆĎĘĚǏĮĹĽŁŃŇÓǑÖØŔŘŚŠŤŮŲÜÝŹŽ",	
	 //Ao (języka Naga) | Ao language (ao)
		"ao": "ǎǐěǒǔñńňřłǍĚǑǓÑŃŇŘŁ",
	 //Gaskoński | Gascon (gn)
		"gn" : "âéèêëïîòôöüùûÂÉÈÊËÏÎÒÔÖÜÙÛ",
	 //Staropruski | Old Prussian (op)
		"op": "ąčęėēģīķłńõōšūžĄČĘĖĒĢĪĶŁŃÕŌŠŪŽ",
	 //Kuriński | Kurin (kv) 
		"kv": "āēīōūǣǫœġḥḳṣṭŋĀĒĪŌŪǢǪŒĠḤḲṢṬŊ",
	 //Degestański | Dagestani (di)
		"di": "āə̄ēīōūȳĀƏ̄ĒĪŌŪȲ",
	 //Szwedzki | Swedish (sv)
		"sv": "åäöÅÄÖ",
	 //Prowansalski | Occitan (oc)
		"oc": "áàâéèêíìîóòôúùûëïüçÁÀÂÉÈÊÍÌÎÓÒÔÚÙÛËÏÜÇ",
	 //Sardyński | Sardinian (sc)
		"sc": "àèéìòùÀÈÉÌÒÙ",
	 //Korsykański | Corsican (co)
		"co": "àèìòùÀÈÌÒÙ",
	 //Retoromański | Romansh (rm)
		"rm": "áàâǎäąéèêěëęíìîǐïįóòôǒöǫúùûǔüųçčśŝšţñņňŋåŧðłÁÀÂǍÄĄÉÈÊĚËĘÍÌÎǏÏĮÓÒÔǑÖǪÚÙÛǓÜŲÇČŚŜŠŢÑŅŇŊÅŦÐŁ",
	 //Arumuński | Aromanian (rup)
		"rup": "âăćĕëîńŏřśţŭźÂĂĆĔËÎŃŎŘŚŢŬŹ",
	 //Szkocki | Scottish Gaelic (gd)
		"gd": "àáâäéèêëìíîïòóôöùúûüýÀÁÂÄÈÊËÉÌÍÎÏÔÓÒÖÙÚÛÜÝ",
	 //Kornwalijski | Cornish (crn)
		"crn": "âêĵôûŵŷÂÊĴÔÛŴŶ",
	 //Liwski | Livonian (liv)
		"liv": "áčėęė́ĩįį́į̃ĺņņ̌õóšųų̃ų̄ų̃̄ų̃̌ų̄̌ỹžÁČĖĘĖ́ĨĮĮ́Į̃ĹŅŅ̌ÕÓŠŲŲ̃Ų̄Ų̃̄Ų̃̌Ų̄̌ỸŽ",
	 //Mordwiński | Moksha (mdf)
		"mdf": "ĺćńśẃẁŕḿǵźj́ɗ́ɗťĹĆŃŚẂẀŔḾǴŹJ́Ɗ́ƊŤ",
	 //Woro | Võro (fiu)
		"fiu": "äöüõšžåÄÖÜÕŠŽÅ",
	 //Kazachski | Kazakh (kaz)
		"kaz": "äïöüÄÏÖÜ",
	 //Gagauski | Gagauz (tut)
		"tut": "ăâäêıîöőûüűĂÂÄÊİÎÖŐÛÜŰ",
	 //Karaimski | Karaim (kdr)
		"kdr": "ėäöüńşźáéíóúĖÄÖÜŃŞŹÁÉÍÓÚ",
	 //Baskijski | Basque (baq)
		"baq": "áéíóúüñÁÉÍÓÚÜÑ",
	 //Keczucki | Quechua (qu)
		"qu": "áéíóúñÁÉÍÓÚÑ",
	 //Ajmarski | Aymara (ay)
		"ay": "ąęįǫųĄĘĮǪŲ",
	 //Portugalski | Portuguese (pt)
		"pt": "áâàãçéêèíîìóôòõôúûùüñÂÁÂÀÃÇÉÊÈÍÎÌÓÔÒÕÚÙÛÜÑ",
	 //Kreolski | Creole (cr)
		"cr": "áéíóúñüÁÉÍÓÚÑÜ",
	 //Mixe-Zoque (mxz)
		"mxz":"āäȧḁēëėẹ̄īïï̈ị̈ōöö̈ọ̈ūüü̈ụ̈ɨɨ̄ɨ̈ɨ̈̈ɨ̣ʉʉ̄ʉ̈ʉ̈̈ʉ̣ɛɛ̄ɛ̈ɛ̈̈ɛ̣ɔɔ̄ɔ̈ɔ̈̈ɔ̣",	
	 //-----------------------------------------------------
}

func main() {
    flagHelp := flag.Bool("help", false, "View Help")
    flagLang := flag.String("lang", "pl", "The language of diacritic marks: \n (ao, ay, ba, baq, br, ch-it, co, cr, crh, crn, csb, cy, cz, di, dk, el, es, et, eu, fi, fiu, fo, fr, fy, ga, gd, gl, gn, gsw, hr, hsb, hu, is, it, kaz, kv, ku, kw, la, lb, liv, lt, lv, mdf, md, mxz, nl, no, oc, op, pl, pt, qu, rm, rom, rup, sc, sk, sl, sq, sq-kosovo, sr, sv, tr, tut, vi, x-lmk)")
    flag.Parse()

    if *flagHelp {
        printHelp()
        return
    }
	
    inputFiles := flag.Args()
    if len(inputFiles) == 0 {
        fmt.Println("OGONKI => Error: No file path specified. \nDisplaying help: OGONKI -help")
        return
    }

    for _, inputFile := range inputFiles {
        fileExt := strings.ToLower(filepath.Ext(inputFile))
            switch fileExt {
            case ".ads", ".adb", ".as", ".asm", ".asp", ".aspx", ".au3", ".avs", ".avsi", ".awk", ".bash" ,".bash_profile" ,".bashrc", ".bat" ,".bb" ,".bi" ,".c" ,".cba" ,".cbf", ".cbh" ,".cfg", ".cbl", ".cd", ".cl" ,".cln", ".cmd", ".cob", ".copy", ".cpy", ".cs", ".csd", ".csh", ".ctg", ".csv", ".cw", ".cxx" ,".d" ,".diff" ,".em", ".epd" ,".erl", ".f", ".f2k", ".f23", ".f77", ".f90", ".f95", ".fen" ,".for", ".forth" ,".gd", ".git", ".gitconfig", ".go" ,".groovy", ".gui", ".h", ".hh", ".hrl", ".hcl", ".hws", ".html", ".hta", ".hex", ".hs",".inf", ".info", ".ini" ,".ino", ".iss", ".java" ,".js", ".jsm", ".json5", ".jsonc", ".jsp", ".kix" ,".kml" ,".kt", ".las", ".lhs", ".lisp", ".log", ".lst", ".lua", ".lpr" ,".lsp", ".mak", ".m", ".matlab" ,".md", ".mib", ".ml" ,".mli", ".mm" ,".mms" ,".mot", ".mxml", ".nfo", ".nim", ".nsi" ,".nsh", ".nt", ".nosql", ".orc", ".osx" ,".out", ".pack" , ".pas", ".pb", ".p", ".php", ".php3" ,".php4" ,".php5", ".phps", ".phpt", ".phtml" ,".plx" , ".pl", ".pm", ".pp" , ".properties" ,".ps1", ".psd1", ".psm1", ".ps" ,".pgn", ".py" ,".pyw",  ".r" ,".r2", ".r3", ".raku", ".rb", ".rbw", ".reg" ,".reb", ".rs" ,".rust", ".s", ".scm" ,".smd", ".shell", ".sh", ".si4", ".sml" ,".splus", ".srt", ".sql" ,".sqlite", ".src", ".srec" ,".ss" ,".stp", ".st", ".sty", ".svg", ".swift", ".shtm" ,".shtml", ".t2t", ".tab" , ".tcl", ".tek", ".tex" ,".thy", ".tsq", ".ts" ,".tsx", ".txt", ".url", ".vb" ,".vba" ,".vbs" ,".v", ".vala" ,".vhdl", ".vh", ".vhd" ,".wer", ".xhtml", ".xht" ,".xml" ,".xsd", ".xsl", ".xslt", ".xul" ,".yaml", ".yml":
            content, err := ioutil.ReadFile(inputFile)
            if err != nil {
                fmt.Printf("Error loading file %s: %s\n", inputFile, err)
        continue
        }

        newContent := replaceOgonki(string(content), *flagLang)
        outputFile := strings.TrimSuffix(inputFile, fileExt) + "_modified" + fileExt

        err = ioutil.WriteFile(outputFile, []byte(newContent), 0644)
        if err != nil {
            fmt.Printf("Error for adding support access paths %s: %s\n", outputFile, err)
            continue
        }
		
        fmt.Printf("Diacritics swap completed successfully. \nNew file saved as %s.\n", outputFile)
        default:
        fmt.Printf("Unsupported file format: %s\n", fileExt)
    }
 }
}

func replaceOgonki(text string, lang string) string {
    ogonkiLetters, ok := ogonki[lang]
    if !ok {
        fmt.Printf("Error: Unknown diacritics language %s.\n", lang)
        return text
    }
    for _, ogonkiLetter := range ogonkiLetters {
        replacement := string(ogonkiLetter)
        text = strings.ReplaceAll(text, replacement, getReplacement(string(ogonkiLetter)))
    }
    return text
}

func getReplacement(letter string) string {replacements := map[string]string{ 	
	"ą": "a", "à": "a", "á": "a", "ã": "a", "ả": "a", "ạ": "a", "ắ": "a", "ằ": "a", "ẵ": "a", "ẳ": "a", "ặ": "a", "â": "a", "ấ": "a", "ầ": "a", "ẫ": "a", "ẩ": "a", "ậ": "a", "ä": "a", "å": "a", "α": "a", "ă": "a", "ā": "a", "ǎ": "a", "a": "a", "a̱": "a", "â": "a", "ä̂": "a", "ạ": "a", "à́": "a", "a̧": "a", "ȃ": "a", "ḁ": "a", "ą̇": "a", "Ą": "A", "À": "A", "Á": "A", "Ã": "A", "Ả": "A", "Ạ": "A", "Ắ": "A", "Ằ": "A", "Ẵ": "A", "Ẳ": "A", "Ặ": "A", "Â": "A", "Ấ": "A", "Ầ": "A", "Ẫ": "A", "Ẩ": "A", "Ậ": "A", "Ä": "A", "Å": "A", "Ă": "A", "Ā": "A", "Ǎ": "A", "Ạ": "A", "A̱": "A", "Ä": "A", "Â": "A", "Ä̂": "A", "Ậ": "A", "À́": "A", "A̧": "A", "Ȃ": "A", "Ą̇": "A", "ƀ": "b", "β": "b", "ḃ": "b", "ß": "B", "Ḃ": "ẞ", "č": "c", "ć": "c", "ç": "c", "ċ": "c", "ĉ": "c", "Č": "C", "Ć": "C", "Ç": "C", "Ĉ": "C", "Ċ": "C", "ð": "d", "đ": "d", "ď": "d", "δ": "d", "ɗ́": "d", "ɗ": "d", "ḋ": "d", "Ď": "D", "Đ": "D", "Δ": "D", "Ɗ́": "D", "Ɗ": "D", "Ḋ": "D", "ė": "e", "ę": "e", "è": "e", "é": "e", "ê": "e", "ë": "e", "ẽ": "e", "ẻ": "e", "ẹ": "e", "ε": "e", "ē": "e", "ě": "e", "ẹ": "e", "e̱": "e", "ë̂": "e", "ệ": "e", "é̄": "e", "è́": "e", "ë̈": "e", "ȩ": "e", "ĕ": "e", "ȇ": "e", "ẹ̄": "e", "ế": "e", "ề": "e", "ể": "e", "ễ": "e", "ệ": "e", "ė́": "e", "̣ɛ": "e", "ɛ̄": "e", "ɛ̈": "e", "ɛ̈̈": "e", "̈̈̈̈ɛ": "e", "Ė": "E", "Ę": "E", "Ē": "E", "È": "E", "É": "E", "Ë": "E", "Ẽ": "E", "Ẻ": "E", "Ẹ": "E", "Ě": "E", "E̱": "E", "Ë̂": "E", "Ệ": "E", "É̄": "E", "È́": "E", "Ë̈": "E", "Ȩ": "E", "Ĕ": "E", "Ê": "E", "Ẹ̄": "E", "Ế": "E", "Ề": "E", "Ể": "E", "Ễ": "E", "Ệ": "E", "Ė́": "E", "f̈": "f", "φ": "f", "ḟ": "f", "F̈": "F", "Φ": "F", "Ḟ": "F", "ģ": "g", "ġ": "g", "γ": "g", "ğ": "g", "g̈": "g", "Ģ": "G", "Ġ": "G", "Γ": "G", "Ğ": "G", "Ǧ": "G", "G̈": "G", "ȟ": "h", "ħ": "h", "ḥ": "h", "Ȟ": "H", "Ħ": "H", "Ḥ": "H", "į": "i", "i̇": "i", "ì": "i", "í": "i", "î": "i", "î": "i", "ï": "i", "ĩ": "i", "ỉ": "i", "ị": "i", "ī": "i", "ι": "i", "ı": "i", "ǐ": "i", "i̱": "i", "ï̂": "i", "i̧": "i", "ĭ": "i", "ȋ": "i", "ị̄": "i", "į́": "i", "į̃": "i", "ï̈": "i", "ị̈": "i", "ɨ": "i", "ɨ̄": "i", "ɨ̈": "i", "į̇": "i", "Į": "I", "İ": "I", "Ì": "I", "Í": "I", "Î": "I", "Ï": "I", "Ĩ": "I", "Ỉ": "I", "Ị": "I", "Ǐ": "I", "I̱": "I", "Ï̂": "I", "I̧": "I", "Ĭ": "I", "Ȋ": "I", "Ī": "I", "Į́": "I", "Į̣̃": "I", "Į̇": "I", "ǰ": "j", "ь": "j", "ĵ": "j", "j́": "j", "Ь": "J", "Ĵ": "J", "J́": "J", "κ": "k", "ķ": "k", "ḳ": "k", "Ķ": "K", "Ḳ": "K", "ĺ": "l", "ľ": "l", "ļ": "l", "ł": "l", "λ": "l", "Ĺ": "L", "Ľ": "L", "Ļ": "L", "Ł": "L", "Λ": "L", "ḿ": "m", "μ": "m", "ṁ": "m", "Ḿ": "M", "Ṁ": "M", "ń": "n", "ñ": "n", "ň": "n", "ņ": "n", "ŉ": "n", "ņ̌": "n", "Ń": "N", "Ñ": "N", "Ň": "N", "Ņ": "N", "Ŋ": "N", "Ņ̌": "N", "õ": "o", "ø": "o", "ó": "o", "ô": "o", "ò": "o", "ö": "o", "ő": "o", "ō": "o", "ω": "o", "ȯ": "o", "ȱ": "o", "ỏ": "o", "ọ": "o", "ố": "o", "ồ": "o", "ỗ": "o", "ổ": "o", "ộ": "o", "ơ": "o", "ớ": "o", "ờ": "o", "ỡ": "o", "ở": "o", "ợ": "o", "ǒ": "o", "ǫ": "o", "o̱": "o", "ö̂": "o", "ó̄": "o", "ò́": "o", "ò̂": "o", "ŏ": "o", "ȏ": "o", "ọ̄": "o", "ö̈": "o", "̣ɔ": "o", "ɔ̄": "o", "ɔ̈": "o", "ɔ̈̈": "o", "ɔ̣̣": "o", "Õ": "O", "Ø": "O", "Ó": "O", "Ô": "O", "Ò": "O", "Ö": "O", "Ő": "O", "Ō": "O", "Ω": "O", "Ȯ": "O", "Ȱ": "O", "Ỏ": "O", "Ọ": "O", "Ố": "O", "Ồ": "O", "Ỗ": "O", "Ổ": "O", "Ộ": "O", "Ơ": "O", "Ớ": "O", "Ờ": "O", "Ỡ": "O", "Ở": "O", "Ợ": "O", "Ǫ": "O", "Ǒ": "O", "O̱": "O", "Ö̂": "O", "Ó̄": "O", "Ò́": "O", "Ò̂": "O", "Ŏ": "O", "Ȏ": "O", "Ọ̄": "O", "π": "p", "ṗ": "p", "Π": "P", "Ṗ": "P", "ř": "r", "ŕ": "r", "ŗ": "r", "σ": "r", "Ř": "R", "Ŕ": "R", "Ŗ": "R", "Σ": "R", "š": "s", "ś": "s", "ş": "s", "ș": "s", "τ": "s", "ŝ": "s", "ś́": "s", "ṣ": "s", "ṡ": "s", "Š": "S", "Ś": "S", "Ş": "S", "Ș": "S", "Ŝ": "S", "Ś́": "S", "Ṣ": "S", "Ṡ": "S", "ṭ": "t", "ť": "t", "ţ": "t", "ț": "t", "ŧ": "t", "ṫ": "t", "Ť": "T", "Ţ": "T", "Ț": "T", "Ṭ": "T", "Ŧ": "T", "Ṫ": "T", "ü": "u", "ú": "u", "ů": "u", "û": "u", "ù": "u", "ũ": "u", "ų": "u", "ū": "u", "ű": "u", "υ": "u", "ǘ": "u", "ǚ": "u", "ǜ": "u", "ǔ": "u", "ŭ": "u", "ǖ": "u", "ụ": "u", "u̱": "u", "ü̂": "u", "ъ": "u", "u̧": "u", "ȗ": "u", "ụ̄": "u", "ụ": "u", "ứ": "u", "ừ": "u", "ử": "u", "ữ": "u", "ự": "u", "ų̃": "u", "̄ų̃̌": "u", "ų̄": "u", "ü̈": "u", "ụ̈": "u", "̣ʉ": "u", "ʉ̄": "u", "ʉ̈": "u", "ʉ̈̈": "u", "ʉ": "u", "Ü": "U", "Ú": "U", "Ů": "U", "Û": "U", "Ù": "U", "Ũ": "U", "Ų": "U", "Ū": "U", "Ű": "U", "Ǖ": "U", "Ǘ": "U", "Ǚ": "U", "Ǜ": "U", "Ǔ": "U", "Ụ": "U", "U̱": "U", "Ŭ": "U", "Ü̂": "U", "Ъ": "U", "U̧": "U", "Ụ̄": "U", "Ụ": "U", "Ứ": "U", "Ừ": "U", "Ử": "U", "Ữ": "U", "Ų̃": "U", "Ų̄": "U", "Ų̃̄": "U", "Ų̃̌": "U", "Ų̄̌": "U", "w̌": "w", "ŵ": "w", "ẃ": "w", "ẁ": "w", "W̌": "W", "Ŵ": "W", "Ẃ": "W", "Ẁ": "W", "x̌": "x", "ξ": "x", "X̌": "X", "Ξ": "X", "ý": "y", "ỳ": "y", "ỹ": "y", "ỷ": "y", "ỵ": "y", "ŷ": "y", "ȳ": "y", "ẏ": "y", "ÿ": "y", "Ӯ": "Y", "Ý": "Y", "Ỳ": "Y", "Ỹ": "Y", "Ỷ": "Y", "Ỵ": "Y", "Ŷ": "Y", "Ȳ": "Y", "Ẏ": "Y", "Ÿ": "Y", "ż": "z", "ź": "z", "ž": "z", "ζ": "z", "Ż": "Z", "Ź": "Z", "Ž": "Z", "Θ": "TH", "Þ": "TH", "Ψ": "PS", "η": "ee", "θ": "th", "þ": "th", "χ": "ch", "ψ": "ps", "ѓ": "gj", "ќ": "kj", "љ": "lj", "њ": "nj", "џ": "dz", "ǿ": "OE", "Ǿ": "OE", "æ": "ae", "ǣ": "ae", "ǽ": "ae", "Æ": "AE", "Ǣ": "AE", "Ǽ": "AE", "œ": "oe", "Œ": "OE", "ə̄": "ə", "Ə̄": "Ə",
    }
		replacement, ok := replacements[letter]
		if ok {
			return replacement
    }
    return letter
}

func printHelp() {
    fmt.Println()                                                                
    fmt.Println("     _/_/      _/_/_/    _/_/    _/      _/  _/    _/  _/_/_/   ")
    time.Sleep(300 * time.Millisecond)	
    fmt.Println("  _/    _/  _/        _/    _/  _/_/    _/  _/  _/      _/     ")
    time.Sleep(300 * time.Millisecond)
    fmt.Println(" _/    _/  _/  _/_/  _/    _/  _/  _/  _/  _/_/        _/     ")
    time.Sleep(300 * time.Millisecond)
    fmt.Println("_/    _/  _/    _/  _/    _/  _/    _/_/  _/  _/      _/     ")
    time.Sleep(300 * time.Millisecond)
    fmt.Println(" _/_/      _/_/_/    _/_/    _/      _/  _/    _/  _/_/_/   ")
    time.Sleep(300 * time.Millisecond)
    fmt.Println()
    fmt.Println("OGONKI v1.0 (c) by Łukasz Wójcik 2023")
    fmt.Println("Program for converting diacritics in unformatted text files.")
    fmt.Println()
    fmt.Println("     Use: ogonki [-help] [-lang flag] [file1.txt file2.txt ...]")
    fmt.Println()
    fmt.Println("    Flag: [ao]  [ay] [ba] [baq] [br] [ch-it] [co] [cr] [crh] [crn]")
    fmt.Println("          [csb] [cy] [cz] [di]  [dk] [el] [es] [et] [eu] [fi]")
    fmt.Println("          [fo]  [fr] [fy] [ga]  [gd] [gl] [gn] [gsw] [hr] [hsb]")
    fmt.Println("          [hu]  [is] [it] [kaz] [kv] [ku] [kw] [la] [lb] [liv]")
    fmt.Println("          [lt]  [lv] [mdf] [md] [mxz] [nl] [no] [oc] [op] [pl]")
    fmt.Println("          [pt]  [qu] [rm] [rom] [rup] [sc] [sk] [sl] [sq] [sr]")
    fmt.Println("          [sv]  [tr] [tut] [vi] [x-lmk] [fiu]")
    fmt.Println()
    fmt.Println(" Example: ogonki -lang fr file.txt")
    flag.PrintDefaults()
    fmt.Println()
    fmt.Println("   Site := [https://github.com/lukaszwojcikdev/ogonki.git]")
    fmt.Println("License := [MIT]")
}
