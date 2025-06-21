# wordup
simple CLI tool that fetches word definitions and stores them locally in `wordbank.man` file. You can later view all saved words in a manpage-style format 


## Features 
- Get word definitions via the command line 
- Saves looked-up words to a local `wordbank.man` file
- View saved word definitions in Unix manpage style 

## Usage
```bash
wordup
```
Displays the saved word definitions in manpage format 

```bash
wordup [WORD]
```
Fetches the definition for the specified `WORD`, prints it to the console, and appends it to `wordbank.man`

## Example
```bash
$ wordup hello
word: hello
definitions:
(1) "Hello!" or an equivalent greeting.
(2) To greet with "hello".
(3) A greeting (salutation) said when meeting someone or acknowledging someone’s arrival or presence.

$ wordup
.SH hello
(1) "Hello!" or an equivalent greeting.
(2) To greet with "hello".
(3) A greeting (salutation) said when meeting someone or acknowledging someone's arrival or presence.
```

## Installation
```bash
$ git clone https://github.com/jeredwong/wordup.git

$ cd wordup

$ go install 
```