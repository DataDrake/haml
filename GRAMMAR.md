```
<opt-whitespace> ::= ' ' | '\t'
     <name-char> ::= <alphanum> | "-"
          <name> ::= <name-char>+
   <proper-name> ::= <alpha> <name>
          <text> ::= <char>* EOL
           <tab> ::= <space> <space> | '\t'
        <indent> ::= <tab>*

      <div-body> ::= <classes> | <attributes> | <text>
     <identifer> ::= "#" <proper-name>
            <id> ::= <identifier> <div-body>
  <element-name> ::= "%" <proper-name>
       <element> ::= <element-name> <identifier>? <div-body>
  <classes-body> ::= <attributes> | <text>
         <class> ::= "." <proper-name>
       <classes> ::= <class>+ <classes-body>
     <attr-char> ::= <alphanum> | <non-doublequote-symbol>
    <attr-value> ::= "'" <attr-char>* "'"
     <attr-name> ::= <opt-whitespace>* <proper-name> ":" <opt-whitespace>*
     <attribute> ::= <attr-name> <attr-value> <opt-whitespace>* ","
    <attributes> ::= "{" <attribute>* "}" <text>

           <tag> ::= <id> | <element> | <classes> | <attributes>

      <contents> ::= <tag> | <text>

          <line> ::= <indent>? <contents>
```
