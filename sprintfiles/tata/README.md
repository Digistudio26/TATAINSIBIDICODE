# NOTES TOOL

## OVERVIEW
It is a command-line tool that allows users to manage short single-line notes. Using this application, users can create collections of notes, open and view them, add new notes, or remove existing notes. The tool takes exactly one argument, which is the name of the collection.

To create or manage a collection coding_ideas the command would be:
```
$ ./notestool coding_ideas
```

## TOOL FUNCTIONALITY
1. Display notes from the collection
2. Add a note to the collection
3. Remove a note from the collection
4. Exit the program

## DATABASE
For each collection, a separate database is created. The database is a plain text file with the same name as the collection, where notes are stored in separate rows. If the collection does not exist, it is created, if it does, it is loaded. Notes persist between the runs of the tool.

## USAGE
```
$ ./notestool testtag
Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two
003 - note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
3

Enter the number of note to remove or 0 to cancel:
3

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
```