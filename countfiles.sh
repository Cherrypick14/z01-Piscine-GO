#!/bin/bash

# Count the number of regular files in the current directory and its subdirectories.
NUM_FILES=$(find . -type f | wc -l)

# Count the number of folders in the current directory and its subdirectories.
NUM_FOLDERS=$(find . -type d | wc -l)

# Print the number of regular files and folders to the console.
# echo "Number of regular files: $NUM_FILES"
# echo "Number of folders: $NUM_FOLDERS"
echo $(($NUM_FILES + $NUM_FOLDERS ))
