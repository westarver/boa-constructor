package boaconstructor 

Package boaconstructor is the implementation of the boa-constructor
GUI app that lets the user lay out the cli of their command line app
by filling in text fields, checking boxes and selecting from drop down
selectors. The names of allowable commands and flags along with the
number, type and optional/required status of the parameters to said
commands and flags can be defined. Where possible validation is done
to catch command line input errors by users of the app being defined
During the design stages your work can be saved and recalled for
editing. The preferred format is now JSON, although the original
format is still viable and not going away any time soon. Saving 
and using the original format may still be the best choice if
you intend to edit it by hand using a text editor. The input script 
format is based on the format used by docopt where the user creates a
usage/help text that is parsed and a map created of the command line 
args actually received. Now that the GUI is available the hand written 
input script has been largely superseded by the JSON format generated 
by the GUI. The JSON data generated by boa-constructor can be passed 
to the Boa package function boa.FromJSON(jsonData string, args []
string) from your app and your app will receive a data structure 
containing the command line args with parameters etc. The same map 
will be obtained by passing to boa.FromHelp(help string,args []
string), a proper input script aka usage string aka help string. Go
code to start your app, get and evaluate the commands received, and
implement the help command can be generated from the GUI. The current
implementation creates three files, main.go, runner.go and help.go

