# Text Completion, Editing, and Auto-Correction Tool

## Introduction

- **Language**: The project is written in Go.
- **Coding Standards**: Code adherence to established good practices.
- **Testing**: Test files are also created for unit testing purposes.

### Text Modifications

The program receives input in the form of a file containing text requiring modifications and an output file where the modified text is stored. Here are some of the modifications the program should executes:

1. **Hexadecimal Conversion**: Replace every instance of `(hex)` with the decimal version of the preceding word.
   
   Example: `"1E (hex) files were added" -> "30 files were added"`

2. **Binary Conversion**: Replace every instance of `(bin)` with the decimal version of the preceding word.
   
   Example: `"It has been 10 (bin) years" -> "It has been 2 years"`

3. **Uppercase Conversion**: Convert the preceding word to uppercase for every instance of `(up)`.

   Example: `"Ready, set, go (up)!" -> "Ready, set, GO!"`

4. **Lowercase Conversion**: Convert the preceding word to lowercase for every instance of `(low)`.

   Example: `"I should stop SHOUTING (low)" -> "I should stop shouting"`

5. **Capitalization**: Capitalize the preceding word for every instance of `(cap)`.

   Example: `"Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"`

   - Additionally, if a number is specified after `(low)`, `(up)`, or `(cap)`, the program should apply the conversion to the specified number of words.

     Example: `"This is so exciting (up, 2)" -> "This is SO EXCITING"`

6. **Punctuation Formatting**: Ensure punctuation marks (., ,, !, ?, :, ;) are properly spaced.

   Example: `"I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"`

   - Handle groups of punctuation like `...` or `!?` appropriately.

     Example: `"I was thinking ... You were right" -> "I was thinking... You were right"`

7. **Single Quote Placement**: Place single quotation marks (' ') adjacent to the corresponding words.

   Example: `"I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"`

   - If there are multiple words between the single quotes, adjust accordingly.

     Example: `"As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'"`

8. **Indefinite Article Correction**: Change 'a' to 'an' if the next word begins with a vowel (a, e, i, o, u) or 'h'.

   Example: `"There it was. A amazing rock!" -> "There it was. An amazing rock!"`

## Usage

1. **Input Preparation**: Create a text file (`sample.txt`) containing sentences with specified modifications indicated in parentheses.

2. **Execution**: Users run the program using the command `go run . sample.txt result.txt`, where `sample.txt` is the input file and `result.txt` is the output file.

3. **Processing**: The program reads the input file line by line, identifies modification instructions within parentheses, and applies the specified modifications to the corresponding parts of the sentences.

4. **Output Generation**: After applying modifications, the program writes the modified sentences to the output file (`result.txt`).

5. **Example Usage**:
   - Input File (`sample.txt`):
     ```
     Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
     ```
   - Command:
     ```
     go run . sample.txt result.txt
     ```
   - Output File (`result.txt`):
     ```
     Simply add 66 and 2 and you will see the result is 68.
     ```

Users can repeat this process with different input files and modification instructions as needed, obtaining modified text in the output file each time.


## Conclusion

You can contribute to this program by testing for edge cases and recommending potential upgrades where applicable.