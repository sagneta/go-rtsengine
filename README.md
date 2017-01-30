
# GO-RTSENGINE

An RTS engine written in golang. Detaches the engine from the UI via a UDP API. Any UI could be hung off of this. 
Tailored for large numbers of simultaneous players within a large world. I am writing this basically to 
continue to improve my golang skills and prevent professional boredom. 

However, suggestions and collaberation of all kinds is appreciated. Please find the architectural design document
for this endeavor within the ./doc directory. All documents are github markdown.


## Installation

The install.sh script downloads all the go tools I use for golang development within emacs. That may or may not entice you.
I also added the env.sh script which, when sourced, will set the GOPATH to the working directory (PWD) in which the script
is run as a convenience.


