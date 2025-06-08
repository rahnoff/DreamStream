# DreamStream

To extract employees' ids from a CSV, copy this column in Excel, paste it in a file, then run these commands:
1. `sed "s/.*/'&'/" ids_as_a_column_without_quotes.txt 1>ids_as_a_column_with_quotes.txt`
2. `sed ':a; N; $!ba; s/\n/,/g' ids_as_a_column_with_quotes.txt 1>customer-emulator.py`.
After that wrap a resulting line with () in Vim.
