// +build utf8

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	//adding support for access paths
	"path/filepath"
)

var ogonki = map[string]string{
    "pl": "ąćłńóśźżęĄĆŁŃÓŚŹŻĘ",
    "de": "äöüßÄÖÜẞ",
    "fr": "éèàçêôûÉÈÀÇÊÔÛ",
    "es": "áíóúéüñÁÍÓÚÉÜÑ",
    "it": "àèéìòùÀÈÉÌÒÙ",
    "cz": "čďěňřšťůýžČĎĚŇŘŠŤŮÝŽ",
    "hu": "áéíóöőúüűÁÉÍÓÖŐÚÜŰ",
    "se": "åäöÅÄÖ",
    "dk": "æøåÆØÅ",
    "no": "æøåÆØÅ",
    "fi": "äöÄÖ",
}

func main() {
	flagHelp := flag.Bool("help", false, "View Help")
	flagLang := flag.String("lang", "pl", "The language of diacritic marks: (pl, de, fr, es, it, cz, hu, se, dk, no, fi)")

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
			case ".txt", ".html", ".htm", ".csv", ".svg", ".srt", ".yaml", ".xml", ".json", ".log", ".cpp", ".js", ".c", ".go", ".java", ".py", ".php", ".css", ".sql", ".md", ".bat", ".ini", ".cfg", ".tex", ".properties", ".sh", ".cs", ".ts", ".rb", ".swift", ".r", ".lua", ".pl", ".scala", ".matlab", ".kotlin", ".rust", ".shell":
			content, err := ioutil.ReadFile(inputFile)
			if err != nil {
				fmt.Printf("Error loading file %s: %s\n", inputFile, err)
		continue
		}

		newContent := replaceOgonki(string(content), *flagLang)

		outputFile := strings.TrimSuffix(inputFile, ".txt") + "_modified.txt"

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
		text = strings.ReplaceAll(text, replacement, getReplacement(replacement))
	}
	return text
}

func getReplacement(letter string) string {
	switch letter {
	case "ą":
		return "a"
	case "ć":
		return "c"
	case "ę":
		return "e"
	case "ł":
		return "l"
	case "ń":
		return "n"
	case "ó":
		return "o"
	case "ś":
		return "s"
	case "ź":
		return "z"
	case "ż":
		return "z"
	case "Ą":
		return "A"
	case "Ć":
		return "C"
	case "Ę":
		return "E"
	case "Ł":
		return "L"
	case "Ń":
		return "N"
	case "Ó":
		return "O"
	case "Ś":
		return "S"
	case "Ź":
		return "Z"
	case "Ż":
		return "Z"
	case "ä":
		return "a"
	case "ö":
		return "o"
	case "ü":
		return "u"
	case "Ä":
		return "A"
	case "Ö":
		return "O"
	case "Ü":
		return "U"
	case "ß":
		return "ẞ"	
	case "é":
		return "e"
	case "è":
		return "e"
	case "ê":
		return "e"
	case "ë":
		return "e"	
	case "à":
		return "a"		
	case "â":
		return "a"
	case "ô":
		return "o"
	case "ù":
		return "u"
	case "û":
		return "u"		
	case "ç":
		return "c"
	case "î":
		return "i"
	case "ï":
		return "i"
	case "ÿ":
		return "y"
	case "œ":
		return "oe"
	case "É":
		return "E"
	case "È":
		return "E"			
	case "Ê":
		return "e"		
	case "Ë":
		return "E"	
	case "À":
		return "A"		
	case "Â":
		return "A"
	case "Ô":
		return "O"
	case "Ù":
		return "U"
	case "Û":
		return "U"
	case "Ç":
		return "C"
	case "Î":
		return "I"	
	case "Ï":
		return "I"
	case "Ÿ":
		return "Y"
	case "Œ":
		return "OE"
	case "á":
		return "a"
	case "í":
		return "i"
	case "ú":
		return "u"
	case "ñ":
		return "ñ"
	case "Á":
		return "A"
	case "Í":
		return "I"
	case "Ú":
		return "U"
	case "Ñ":
		return "N"
	case "ì":
		return "i"
	case "ò":
		return "o"
	case "Ì":
		return "I"
	case "Ò":
		return "O"
	case "č":
		return "c"
	case "ď":
		return "d"
	case "ě":
		return "e"
	case "ň":
		return "n"
	case "ř":
		return "r"
	case "š":
		return "s"
	case "ť":
		return "t"
	case "ů":
		return "u"
	case "ý":
		return "y"
	case "ž":
		return "z"
	case "Č":
		return "C"
	case "Ď":
		return "D"
	case "Ě":
		return "E"
	case "Ň":
		return "N"
	case "Ř":
		return "R"
	case "Š":
		return "S"
	case "Ť":
		return "T"
	case "Ů":
		return "U"
	case "Ý":
		return "Y"
	case "Ž":
		return "Z"
	case "ő":
		return "o"
	case "ű":
		return "u"
	case "Ő":
		return "O"
	case "Ű":
		return "U"
	case "å":
		return "a"
	case "Å":
		return "A"
	case "æ":
		return "a"
	case "ø":
		return "a"
	case "Æ":
		return "A"
	case "Ø":
		return "A"
	default:
		return letter
	}
}

func printHelp() {
  fmt.Println()                                                                
  fmt.Println("     _/_/      _/_/_/    _/_/    _/      _/  _/    _/  _/_/_/  ") 
  fmt.Println("  _/    _/  _/        _/    _/  _/_/    _/  _/  _/      _/     ")
  fmt.Println(" _/    _/  _/  _/_/  _/    _/  _/  _/  _/  _/_/        _/     ")
  fmt.Println("_/    _/  _/    _/  _/    _/  _/    _/_/  _/  _/      _/     ")   
  fmt.Println(" _/_/      _/_/_/    _/_/    _/      _/  _/    _/  _/_/_/    ")   
	fmt.Println()
	fmt.Println("OGONKI v1.0 (c) by Łukasz Wójcik")
	fmt.Println("Program for converting diacritical characters in unformatted text files.")
	fmt.Println()
	fmt.Println("It supports languages: Polish, Spanish, German, French, Italian, Czech, Hungarian, Swedish, Danish, Norwegian, Finnish")
	fmt.Println()
  fmt.Println("Use: ogonki [-help] [-lang flag] [file1.txt file2.txt ...]")
	fmt.Println("Flags: [pl] [es] [de] [fr] [it] [cz] [hu] [se] [dk] [no] [fi]")
	fmt.Println("Example: ogonki -lang fr file.txt")
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Site:")
	fmt.Println("https://github.com/lukaszwojcikdev/ogonki.git")
}

