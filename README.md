hu
==

Hu is an experiment in parsing. A "term" has a label (I think of it as a headword) and zero or more definitions. Deciding which definition the term should use in any given sentence is what I mean by "parsing" in this context.

This experimental strategy for parsing is to give each definition a signature, which includes a list of other term signatures that can be taken as inputs and list that can be taken as outputs. Since it is convenient for multiple definitions to use the same signature (one could think of it as a part of speech, e.g. "verb"), we also give the signature a label.

Using only the terms, signatures, and definitions, it should be possible to determine the valid parsing(s) for a sentence.  The combinatorics can quickly get out of hand, but hopefully a computer can deal with this better than a human.
