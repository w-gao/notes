# vim shortcuts

- Created on March 28, 2021
- Updated on Dec 16, 2021
- Updated on Dec 21, 2022


## Normal mode

### Move cursor (motion)

- `j`, `k`: down, up
- `h`, `l`: left, right
- `b`, `w`: left, right word
- `{`, `}`: up, down block
- `gg`: top of file
- `G`: bottom of file
- Note: prepending a number before the command would execute that command # times (e.g.: `10j`: move down 10 lines)


### Cut / copy / paste

- `d` + motion: cut (delete and copy), e.g.:
  - `dw`: cut word
  - `dd`: cut line
  - `d5j`: cut next 5 lines
- `y` + motion: copy (yank), e.g.:
  - `yw`: copy word
  - `yy`: copy line
- `P`: paste before cursor
- `p`: paste after cursor


### Undo / redo

- `u`: undo
- `ctrl + R`: redo


### Enter visual mode

- `v`: select characters
- `V`: select lines
- `ctrl + V`: select blocks
  - useful for adding / removing multi-line comments, for example


### Enter insertion mode

- `i`: go to insertion mode
- `o`: go to insertion mode at next line


## Visual mode

- cut and copy
  - `d` or `x`: ut (delete and copy to buffer)
  - `y`: copy
- `>` and `<`: indent and dedent
- e.g.: `2>`: indent twice
- `~`: flip characters (upper <-> lower)
- `U`: turn all to uppercase
- `u`: turn all to lowercase


## Markers

- `m<reg>`: mark cursor and store to any registers
  - e.g.: `ma`: mark and store to register a
- `\`<reg>`: jump to marker at register (exact place)
- `'<reg>`: jump to marker at register (line)


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



## Misc

Normal mode

- `gUU`: turn entire line to uppercase
- `guu`: turn entire line to lowercase

- `ctrl + A`: increment number under cursor
- `ctrl + X`: decrement number under cursor

