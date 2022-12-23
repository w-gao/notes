# vim shortcuts

- Created on March 28, 2021
- Updated on Dec 16, 2021
- Updated on Dec 21, 2022


## Normal mode

### Move cursor

- `j`, `k`: up, down
- `h`, `l`: left, right
- `b`, `w`: left, right word
- `{`, `}`: up, down block
- `gg`: top of file
- `G`: bottom of file
- Note: prepending a number before the command would execute that command # times (e.g.: `10j`: move down 10 lines)


### Cut / copy / paste

- `dd`: cut line (delete line and copy to buffer)
- `yy`: copy (yank) line
- `P`: paste before cursor
- `p`: paste after cursor
- Note: in _visual mode_, `y`: copy block


### Selection (visual mode)

- `v`: select characters
- `V`: select lines
- `ctrl + V`: select blocks
  - useful for adding / removing multi-line comments, for example
- once selected, cut and copy are:
  - `d` or `x`: cut (delete and copy to buffer)
  - `y`: copy


### Undo / redo

- `u`: undo
- `ctrl + R`: redo


### Insertion mode

- `i`: go to insertion mode
- `o`: go to insertion mode at next line


## Search

- `/` + term, hit enter
- `N`: next, `shift + N`: previous


## Adding / removing multi-line comments


### Approach 1: using visual block mode

- Add comment block
  1. `ctrl + V` - enter visual block mode
  2. select block (characters will be inserted at the left of block)
  3. `shift + I` - enter insertion mode from visual mode
  4. insert any text, e.g.: `// `, `# `, etc
  5. `Esc Esc` - exit insertion mode, then visual block mode
- Remove comment block
  1. `ctrl + V`
  2. select block
  3. `d` - delete

### Approach 2: using norm

- Add comment block
  1. `shift + V` - enter visual mode
  2. select lines
  3. type `:norm i` + characters to insert before each line, e.g.: `:norm i // `
    - Note: the selected range will be present, so it will look like `:'<.'>norm i// `
- Remove comment block
  1. `shift + V`
  2. type `:norm xxx`, the number of characters to remove before each line


See: https://stackoverflow.com/a/15588798


