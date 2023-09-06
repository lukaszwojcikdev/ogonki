## Making Life Easier for Programmers: A Simple Solution to Diacritical Marks in Text Files ##

# DESCRIPTION

This program is a tool for replacing `tails` in unformatted text files (preferably encoded in UTF-8).

It supports 66 languages.

At the beginning, the program defines a `tail map`, which contains sets of characters specific to each of the supported languages.

The main function: `main` starts with handling flags. The available flags are `-help` (display help) and `-lang` <language> (choose the `tail` language, `pl` by default).

The program then takes file paths as command line arguments.

If no path is specified, the program displays an error message and exits.

For each given access path, the program checks the file extension (recognizes up to 180 extensions of non-software text files).

If it is one of the supported extensions, the program reads the content of the file, replaces the tails using the `replaceTags` function for the given language, and then saves the modified content to a new file named `<original_file_name>_modified.txt`.

The process of replacing tails is done using the `replaceTails` function, which iterates through all the tailing characters for a given language and replaces them with the appropriate replacement characters.

The `getReplacement` function specifies replacement characters for individual tails.
Diacritics are changed, e.g., `ą` to `a`, `ć` to `c`, etc., with the remaining characters.
Other characters in the text are replaced by themselves.

The `printHelp` function displays help on how to use the program, describes the supported flags and languages, and gives examples of use.

# MANUAL

To run the program, you must have a Go compiler installed on your device.

Then follow these steps:

1. **Clone the repository to your device:**
   
git clone https://github.com/lukaszwojcikdev/ogonki.git

2. **Go to the project directory:**
   
cd tails

4. **Compile the program:**
   
go build ogonki.go

4. **Starting the program:**
   
*Linux:*

./ogonki.go

*Windows:*

ogonki.exe

## Author

This program was created by [Lukasz Wójcik].

If you have any questions or comments, please contact me at kontakt(at)lukaszwojcik.eu

## License

This project is licensed under [[MIT](https://opensource.org/license/mit/)].

Details can be found in the file [[LICENSE](https://github.com/lukaszwojcikdev/ogonki/blob/main/LICENSE)].

I hope you liked the program and it will be helpful to you.

If you have any questions, feel free to ask, I will try to answer as soon as possible.
