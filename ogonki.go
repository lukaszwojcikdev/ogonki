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
		"sq": "çëë̈f̈g̈šïx̌ÿžÇËË̈F̈G̈ŠÏŽX̌ŸŽ", 
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
    flagLang := flag.String("lang", "pl", "The language of diacritic marks: \n (ao, ay, ba, baq, br, ch-it, co, cr, crh, crn, csb, cy, cz, di, dk, el, es, et, eu, fi, fiu, fo, fr, fy, ga, gd, gl, gn, gsw, hr, hsb, hu, is, it, kaz, kv, ku, kw, la, lb, liv, lt, lv, mdf, md, mxz, nl, no, oc, op, pl, pt, qu, rm, rom, rup, sc, sk, sl, sq, sr, sv, tr, tut, vi, x-lmk)")
    flag.Parse()

    if *flagHelp {
        printHelp()
        return
    }
    inputFiles := flag.Args()
    if len(inputFiles) == 0 {
        fmt.Println("OGONKI => Error: No file path specified. Displaying help: OGONKI -help")
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
        fmt.Printf("Diacritics swap completed successfully. New file saved as %s.\n", outputFile)
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

func getReplacement(letter string) string {
    replacements := map[string]string{
        "ą,à,á,ã,ả,ạ,ă,ắ,ằ,ẵ,ẳ,ặ,â,ấ,ầ,ẫ,ẩ,ậ,ä,å,α,ă,ā,ǎ,ạ,a̱,ä,â,ä̂,ạ,à́,a̧,ȃ,ḁ,ą̇":"a","Ą,À,Á,Ã,Ả,Ạ,Ă,Ắ,Ằ,Ẵ,Ẳ,Ặ,Â,Ấ,Ầ,Ẫ,Ẩ,Ậ,Ä,Å,Ă,Ā,Ǎ,Ạ,A̱,Ä,Â,Ä̂,Ậ,À́,A̧,Ȃ,Ą̇":"A","ƀ,β,ḃ":"b","ß,Ḃ":"ẞ","č,ć,ç,ċ,ĉ":"c","Č,Ć,Ç,Ċ,Ĉ,Ċ":"C","ð,đ,ď,δ,ɗ́,ɗ,ḋ":"d","Ď,Đ,Δ,Ɗ́,Ɗ,Ḋ":"D","ė,ę,ē,è,é,ê,ë,ẽ,ẻ,ẹ,ε,ē,ě,ẹ,e̱,ë̂,ệ,é̄,è́,ë̈,ȩ,ĕ,ȇ,ẹ̄,ế,ề,ể,ễ,ệ,ė́,̣ɛ,ɛ̄,ɛ̈,ɛ̈̈,̈̈̈̈ɛ":"e","Ė,Ę,Ē,Ẹ,È,É,Ê,Ë,Ẽ,Ẻ,Ẹ,Ě,E̱,Ë̂,Ệ,É̄,È́,Ë̈,Ȩ,Ĕ,Ê,Ẹ̄,Ế,Ề,Ể,Ễ,Ệ,Ė́":"E","f̈,φ,ḟ":"f","F̈,Φ,Ḟ":"F","ģ,ġ,γ,ğ,g̈":"g","Ģ,Ġ,Γ,Ğ,Ǧ,G̈":"G","ȟ,ħ,ḥ":"h","Ȟ,Ħ,Ḥ":"H","į,i̇,ì,í,î,î,ï,ĩ,ỉ,ị,ī,ι,ı,ǐ,i̱,ï̂,i̧,ĭ,ȋ,ị̄,į́,į̃,ï̈,ị̈,ɨ,ɨ̄,ɨ̈,ɨ̈,̈ɨ,į̇":"i","Į,İ,Ì,Í,Î,Ï,Ĩ,Ỉ,Ị,Ī,Ǐ,I̱,Ï̂,I̧,Ĭ,Ȋ,Ī,Į́,Į̣̃,Į̇":"I","ǰ,ь,ĵ,j́":"j","Ь,Ĵ,J́":"J","κ,ķ,ḳ":"k","Ķ,Ḳ":"K","ĺ,ľ,ļ,ł,λ":"l","Ĺ,Ľ,Ļ,Ł,Λ":"L","ḿ,μ,ṁ":"m","Ḿ,Ṁ":"M","ń,ñ,ň,ņ,ŉ,ņ̌":"n","Ń,Ñ,Ň,Ņ,Ŋ,Ņ̌":"N","õ,ø,ó,ô,ò,ö,ő,ō,ω,ȯ,ȱ,ỏ,ọ,ố,ồ,ỗ,ổ,ộ,ơ,ớ,ờ,ỡ,ở,ợ,ǒ,ǫ,o̱,ö̂,ó̄,ò́,ò̂,ŏ,ȏ,ọ̄,ö̈,ö,̣ɔ,ɔ̄,ɔ̈,ɔ̈̈,ɔ̣̣":"o","Õ,Ø,Ó,Ô,Ò,Ö,Ő,Ō,Ω,Ȯ,Ȱ,Ỏ,Ọ,Ố,Ồ,Ỗ,Ổ,Ộ,Ơ,Ớ,Ờ,Ỡ,Ở,Ợ,Ǫ,Ǒ,O̱,Ö̂,Ó̄,Ò́,Ò̂,Ŏ,Ȏ,Ọ̄":"O","π,ṗ":"p","Π,Ṗ":"P","ř,ŕ,ŗ,σ":"r","Ř,Ŕ,Ŗ":"R","š,ś,ş,ș,τ,ŝ,ś́,ṣ,ṡ":"s","Š,Ś,Ş,Ș,Σ,Ŝ,Ś́,Ṣ,Ṡ":"S","ṭ,ť,ţ,ț,ŧ,ṫ":"t","Ť,Ţ,Ț,Ṭ,Ŧ,Ṫ":"T","ü,ú,ů,û,ù,ũ,ų,ū,ű,υ,ǘ,ǚ,ǜ,ǔ,ŭ,ǖ,ụ,u̱,ü̂,ъ,u̧,ȗ,ụ̄,ụ,ứ,ừ,ử,ữ,ự,ų̃,ų̄,ų̃,̄ų̃̌,ų̄,ü̈,ụ̈,̣ʉ,ʉ̄,ʉ̈,ʉ̈̈,ʉ":"u","Ü,Ú,Ů,Û,Ù,Ũ,Ų,Ū,Ű,Ǖ,Ǘ,Ǚ,Ǜ,Ǔ,Ụ,U̱,Ŭ,Ü̂,Ъ,U̧,Ŭ,Ụ̄,Ụ,Ứ,Ừ,Ử,Ữ,Ų̃,Ų̄,Ų̃̄,Ų̃̌,Ų̄̌":"U","w̌,ŵ,ẃ,ẁ":"w","W̌,Ŵ,Ẃ,Ẁ":"W","x̌,ξ":"x","X̌,Ξ":"X","ý,ỳ,ỹ,ỷ,ỵ,ŷ,ȳ,ẏ,ÿ":"y","Ӯ,Ý,Ỳ,Ỹ,Ỷ,Ỵ,Ŷ,Ȳ,Ẏ,Ÿ":"Y","ż,ź,ž,ζ":"z","Ż,Ź,Ž,Ź":"Z","Θ,Þ":"TH","Ψ":"PS","η":"ee","θ,þ":"th","χ":"ch","ψ":"ps","ѓ":"gj","ќ":"kj","љ":"lj","њ":"nj","џ":"dz","ǿ":"OE","Ǿ":"OE","æ,ǣ,ǽ":"ae","Æ,Ǣ,Ǽ":"AE","œ":"oe","Œ":"OE","ə̄":"ə","Ə̄":"Ə",
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
