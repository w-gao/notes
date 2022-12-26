# vim shortcuts

- Created on March 28, 2021
- Updated on Dec 16, 2021
- Updated on Dec 25, 2022


## Normal mode

### Move cursor (motion)

- `j`, `k`: down, up
- `h`, `l`: left, right
- `b`, `w`: left, right word
- `{`, `}`: up, down block
- `gg`: top of file
- `G`: bottom of file
- Note: prepending a number before the command would execute that command # times (e.g.: `10j`: move down 10 lines)
- More motions
  - `e`: end of word
  - `B`, `W`: left, right word (spaces only)
  - `f + <sym>`: find first occurance of symbol
    - `F + <sym>`: find left (backwards)
  - `t + <sym>`: move **till** first occurance of symbol
    - `T + <sym>`: move left (backwards)
  - `0`: to beginning of line
  - `^` (or `0w`): to first non-space character of line
  - `$`: to end of line
  - `%`: to matching braket


### Cut / copy / paste

- `d` + motion: cut (delete and copy), e.g.:
  - `dw`: cut word
  - `dd`: cut line
  - `d5j`: cut next 5 lines
  - or, `diw`: delete the word the cursor is in
- `D`: cut from cursor to end of line
- `y` + motion: copy (yank), e.g.:
  - `yw`: copy word
  - `yy`: copy line
- when cursor is inside a word, `"`, `(`, `[`, or `{`,
  - `di + <sym>`: delete everything inside block
  - `da + <sym>`: delete everything inside block (including symbols)
  - works for `y`, `c` as well, e.g.:
    - `ci[`: change text inside `[xxx]`
- `P`: paste before cursor
- `p`: paste after cursor


### Quick changes

- `x`: delete character under cursor
- `r`: replace charater under cursor
- `>>` and `<<`: indent and dedent


### Undo / redo

- `u`: undo
- `ctrl + R`: redo


### Enter visual mode

- `v`: select characters
- `V`: select lines
- `ctrl + V`: select blocks
  - useful for adding / removing multi-line comments, for example


### Enter insertion mode

- `i`: enter insertion mode
- `a`: enter insertion mode (append after cursor)
- `o`: begin new line below cursor and enter insertion mode
- `O`: begin new line above cursor and enter insertion mode
- `I`: enter insertion mode at the beginning of line
- `A`: enter insertion mode at the end of line
- `c` + motion: cut text selected by motion, then enter insertion mode
  - `cc`: cut entire line, then enter insertion mode
- `C` + motion: similar to `D`, cut from cursor to end of line, then enter insertion mode


## Visual mode

- cut and copy
  - `d` or `x`: cut
  - `y`: copy (yank)
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
- `N`: go to next, `shift + N`: go to previous


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


## Split windows

- `ctrl + W s`: split window (horizontal)
- `ctrl + W v`: split window (vertical)
- `ctrl + W q`: quit current window
- moving between windows
  - `ctrl + W <kjhl>`: move between splits
  - `ctrl + W p`: go to previous window
  - `ctrl + W w`: go to next window
  - `ctrl + W t`: go top top left window
  - `ctrl + W b`: go to bottom right window
- moving windows
  - `ctrl + W x`: exchange two windows
  - `ctrl + W r`: rotate windows
  - `ctrl + W R`: rotate (in opposite direction)
- resizing
  - `ctrl + W <num> +` or `ctrl + W <num> -`: resize currernt window height
  - `ctrl + W <num> <` or `ctrl + W <num> >`: resize current window width
  - `ctrl + W \_`: resize window to max height
  - `ctrl + W |`: resize window to max width
  - `ctrl + W =`: resize all windows to equal dimensions
- commands
  - `:sp` or `split`: split window
  - `:vsp`: vertical split
    - `:vsp <filename>`: vertical split new file
  - `:e <filename>`: change current file


## Different tabs

- To open multiple file in terminal: `vim -p file1 file2 file3`
- To open new file in vim: `:tabedit file`
- `gt`: go to next tab
- `gT`: go to previous tab
- `<num> gt`: go to tab


## Misc

Normal mode

- `gUU`: turn entire line to uppercase
- `guu`: turn entire line to lowercase

- `ctrl + A`: increment number under cursor
- `ctrl + X`: decrement number under cursor


Commands

- `:Explore` - built-in file explorer
- `:Lex` - similar to above, but stays open when opening files


