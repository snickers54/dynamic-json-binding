# dynamic-json-binding

I just wanted to play a bit with the AST, unfortunately there is no way to instantiate a type only from it's name (a string). So I had to register my types .. 

It's absolutely something to do, it's not reliable, not handling incomplete json payload, you could have more than one struct matching your payload. 

But it was fun to work with the AST. In itself, because I've to register my type, I don't need to parse the AST.. I could just use reflection .. Whatever ..