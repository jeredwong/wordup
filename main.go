package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	
	// Constants and paths 
	const BasePath = "https://api.dictionaryapi.dev/api/v2/entries/en/"
	const WordBankFile = "wordbank.man"
	const WordBankScript = "open-wordbank.sh"

	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)

	wordBankFilePath := filepath.Join(exeDir, WordBankFile)
	wordBankScriptPath := filepath.Join(exeDir, WordBankScript)

	// Structs 
	type Definition struct {
		Definition string
	}

	type Meaning struct {
		Definitions []Definition
	}

	type Word struct {
		Word string 
		Meanings []Meaning
	}

	// Initialise wordbank.man if does not exist 
	if _, err := os.Stat(wordBankFilePath); errors.Is(err, os.ErrNotExist) {
		f, err := os.OpenFile(wordBankFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer f.Close()

		if _, err := fmt.Fprint(f, ".TH WORDBANK\n"); err != nil {
			f.Close() // ignore error; Write error takes precedence
			log.Fatalf("Error appending to file: %v", err)
		}
	}

	// Initialise open-wordbank.sh if does not exist 
	if _, err := os.Stat(wordBankScriptPath); errors.Is(err, os.ErrNotExist) {
		f, err := os.OpenFile(wordBankScriptPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer f.Close()

		if _, err := fmt.Fprintf(f, "#!/bin/sh\nWORDBANK=%s\ngroff -man -Tascii $WORDBANK | less\n", wordBankFilePath); err != nil {
			f.Close() // ignore error; Write error takes precedence
			log.Fatalf("Error appending to file: %v", err)
		}
	}

	if len(os.Args) == 2 { // Retrieve word definition

		word := os.Args[1]

		// log.Printf("INFO: making GET request for word \"%s\"\n", word)

		// Making GET request
		resp, err := http.Get(BasePath + word)
		if err != nil { 
			log.Fatalf("Error making GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			log.Fatalf("Definition for \"%s\" not found", word)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Status Code not 200: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Unmarshalling response
		var words []Word
		if err := json.Unmarshal(body, &words); err != nil {
			log.Fatalf("Error unmarshalling response body: %v", err)
		}

		fmt.Printf("word: %v\n", words[0].Word)

		fmt.Println("definitions:")
		for i, meaning := range words[0].Meanings {
			definition := meaning.Definitions[0].Definition
			fmt.Printf("(%d) %v\n",i+1 , definition)
		}

		// Add word to word bank 
		f, err := os.OpenFile(wordBankFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer f.Close()

		if _, err := fmt.Fprintf(f, ".SH %s\n", word); err != nil {
			f.Close() // ignore error; Write error takes precedence
			log.Fatalf("Error appending to file: %v", err)
		}

		for i, meaning := range words[0].Meanings {
			definition := meaning.Definitions[0].Definition
			s := fmt.Sprintf("(%d) %v\n",i+1 , definition)
			if _, err := fmt.Fprintf(f, "%s\n", s); err != nil {
				f.Close() 
				log.Fatal(err)
			}
		}

	} else if len(os.Args) == 1 { // Display wordbank

		cmd := exec.Command(wordBankScriptPath)

		// Attach current terminal's input/output to the command 
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("Invalid arguments provided")
	}

}
