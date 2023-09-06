## Making Life Easier for Programmers: 
A Simple Solution to Diacritical Marks in Text Files ##

# OPIS
Ten program jest narzędziem do zamiany `ogonków` w plikach tekstowych niesformatowanych (najlepiej kodowanych w UTF-8).

Obsługuje 66 języków.

Na początku program definiuje mapę `ogonki`, która zawiera zestawy znaków specyficznych dla każdego z obsługiwanych języków.

Główna funkcja `main` rozpoczyna od obsługi flag. 
Dostępne flagi to `-help` (wyswietlenie pomocy) i `-lang <jezyk>` (wybór języka ogonków, domyślnie `pl`).

Następnie program pobiera ścieżki dostępu do plików jako argumenty wiersza poleceń.

Jeśli nie zostanie podana żadna ścieżka, program wyświetli komunikat o błędzie i zakończy działanie.

Dla każdej podanej ścieżki dostępu program sprawdza rozszerzenie pliku.

Jeśli jest to jedno z obsługiwanych rozszerzeń, program wczytuje zawartość pliku, zamienia ogonki przy użyciu funkcji `replaceOgonki` dla zadanego języka, a następnie zapisuje zmodyfikowaną zawartość do nowego pliku o nazwie `<oryginalna_nazwa_pliku>_modified.txt`.

Proces zamiany ogonków odbywa się za pomocą funkcji `replaceOgonki`, która iteruje przez wszystkie znaki ogonków dla danego języka i zamienia je na odpowiednie zastępcze znaki.

Funkcja `getReplacement` określa zastępcze znaki dla poszczególnych ogonków. 
Zamianie ulegają, diakrytyki, np: `ą` na `a` , `ć` na `c`, itd z pozostałymi znakami. 
Pozostałe znaki wytępujące w tekście zamieniane są na siebie same.

Funkcja `printHelp` wyświetla pomoc dotyczącą korzystania z programu, opisuje obsługiwane flagi i języki oraz podaje przykłady użycia.

# INSTRUKCJA

Aby uruchomić program, musisz mieć zainstalowany kompilator języka Go na swoim urządzeniu. 

Następnie wykonaj następujące kroki:

1. **Sklonuj repozytorium na swoje urządzenie:**
   
git clone https://github.com/lukaszwojcikdev/ogonki.git

2. **Przejdź do katalogu z projektem:**
   
cd ogonki

4. **Skompiluj program:**
   
go build ogonki.go

4. **Uruchomienie programu:**
   
*Linux:*

./ogonki.go

*Windows:*

ogonki.exe

## Autor

Ten program został stworzony przez [Łukasz Wójcik]. 

Jeśli masz jakiekolwiek pytania lub uwagi, skontaktuj się ze mną pod adresem kontakt(at)lukaszwojcik.eu

## Licencja

Ten projekt jest objęty licencją [[MIT](https://opensource.org/license/mit/)]. 

Szczegóły znajdują się w pliku [[LICENSE](https://github.com/lukaszwojcikdev/ogonki/blob/main/LICENSE)].

Mam nadzieję, że program Ci się spodobał i będzie Ci pomocny.

Jeśli masz jakieś pytania, śmiało pytaj, postaram się odpowiedzieć jak najszybciej.
