#!/bin/bash

# # Find all the `.sh` files in the current directory and its subdirectories.
# SH_FILES=$(find . -type f -name "*.sh")

# # Sort the `.sh` files in descending order.
# SH_FILES_SORTED=$(sort -r <<< "$SH_FILES")

# # Print the names of the `.sh` files to the console.
# for SH_FILE in $SH_FILES_SORTED; do
#   echo "$SH_FILE" | sed 's/\.sh$//'
# done

find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r