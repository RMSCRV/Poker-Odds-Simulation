#!/bin/bash

# Define the output file
OUTPUT_FILE="codebase.txt"

# Ensure the output file is empty before starting
echo "" > "$OUTPUT_FILE"

# Get the entry directory from the first argument, default to current directory
ENTRY_DIR="${1:-.}"

# Recursively find all files (excluding the output file itself)
find "$ENTRY_DIR" -type f ! -name "codebase.txt" | while read -r file; do
    # Get the relative path
    REL_PATH="${file#$ENTRY_DIR/}"
    
    # Extract filename
    FILE_NAME="$(basename "$file")"
    
    # Append header information
    echo "-------------------------" >> "$OUTPUT_FILE"
    echo "File path: $REL_PATH" >> "$OUTPUT_FILE"
    echo "File name: $FILE_NAME" >> "$OUTPUT_FILE"
    echo "" >> "$OUTPUT_FILE"
    
    # Append the file content
    cat "$file" >> "$OUTPUT_FILE"
    
    # Add a separator after each file
    echo -e "\n--------------------------------------------\n" >> "$OUTPUT_FILE"
done

echo "Codebase collected in $OUTPUT_FILE"
