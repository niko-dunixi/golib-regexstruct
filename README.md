# RegexMatch

A library for parsing regex matches into `maps` and generic `T` structs when compiled with Go 1.18.

## Why?

This is really simple functionality. Like, stupid simple. So why create a library for it? The answer
lies with the DRY principal. Don't Repeat Yourself. I have copy/pasted and had to look up how to
associate named regex matches to their values countless times.

With the introduction of Generics in 1.18, this makes even more sense to implement now.
