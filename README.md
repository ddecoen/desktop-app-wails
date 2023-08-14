# desktop-app-wails
Build a Vale-like linter for prose

## Attempt to build a Vale-like linter for prose
1. Be able to read at least 200 words
2. Run Vale or something similar against the entered text or markdown file
3. Provide the user with the analysis and suggestions from the linter

## Examples of Vale using CLI
(![Vale](<Vale CLI.png>))

## Examples of Frontend App making suggestions
![Wails](<Wails App Spell Check.png>)

## Assignment Findings and Issues
The key **takeaway** from this assignment is that Vale does **not** work well with Wails. It did not seem to want to work as an executable file. The TL;DR is that the build did not work.  
### Solutions  
The workaround is to download [Vale](https://github.com/errata-ai/vale) and install it to one's local machine. After that it would be as easy as using Vale's built in CLI to run vale on any text or markdown file. Here is an example of how one company uses this in practice: [Red Hat](https://redhat-documentation.github.io/vale-at-red-hat/docs/main/user-guide/using-vale-cli/).  
### Another option
It would also be possible just to have Wails up in the Dev environment. On MacOS (or Windows) the Wails Dev will open up as an app (or project) that allows the user to interact with Wails.
![App](<App Screen.png>)

