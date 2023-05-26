# go-funge98

An implementation of the [Funge-98 spec, Befunge-98-flavored](https://github.com/catseye/Funge-98) in Go.

## Implementations Notes

### Funge-Space

- The indexing of the Funge-Space starts at `0,0`. **TODO?:** map to `(-x,x)`

### Stack

- This implementation uses a single stack.