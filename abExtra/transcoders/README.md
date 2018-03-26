### About

#### Notes on using chidley
To generate the structs to serialize the DataCite XML schema (kernel 4) we are using a package
called chidely.  We need only generate the structs since the code generation aspect of the chidley
generates way more code than is necessary for our use.  

Simply being able to mashal the XML data into the structs is fine.  From there we can parse
and rebuild a schema.org/Dataset representation of it.
