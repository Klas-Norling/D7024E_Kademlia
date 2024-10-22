#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <file_path> <string>"
    exit 1
fi

# Assign arguments to variables
FILE_PATH=$1
NEW_CONTENT=$2

# Check if the file exists
if [ ! -f "$FILE_PATH" ]; then
    echo "Error: File '$FILE_PATH' does not exist."
    exit 1
fi

# Replace the content of the file with the string
echo "$NEW_CONTENT" > "$FILE_PATH"

echo "File content replaced with: '$NEW_CONTENT'"