#!/bin/sh
# Format code 
WORDBANK=wordbank.man
groff -man -Tascii $WORDBANK | less
