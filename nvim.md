# Awesome Configs
[Allaman](https://github.com/Allaman/nvim)
<BR/>
[Keymappings](https://github.com/Lazytangent/nvim-conf)
<BR/>

[Plugins & golangci](https://github.com/fablol/.cfg/tree/master/.config/nvim)

# Keymappings
<details>
  <summary>Keymappings</summary>
    
  ```
  \2    Telescope Buffers
  \t    Trouble
  \w    Telescope search <cword>
  \s    Telescope search word
  \3    LazyGit
  R     Replace
    
  ```
</details>

<details>
  <summary>go.nvim Keymappings</summary>
    
  ```  
  :ReplToggle
  :GoBreakToggle
  :GoDebug [OPTIONS]
    -c, --compile         compile and run\n"
    -r, --run             run\n"
    -t, --test            run tests\n"
    -R, --restart         restart\n"
    -s, --stop            stop\n"
    -h, --help            display this help and exit\n"
    -n, --nearest         debug nearest file\n"
    -p, --package         debug package\n"
    -f, --file            display file\n"
    -b, --breakpoint      set breakpoint\n"
    -T, --tag             set tag"
  
  DAP
    r = run
    c = continue
    n = step_over
    s = step_into
    o = step_out
    S = stop
    u = up
    D = down
    C = run_to_cursor
    b = toggle_breakpoint
    P = pause
```  
</details>
  
<details>
  <summary>Built in Keymappings</summary>
    
  ```
  tabe  Open New tab
  tabn  Go to next tab
  tabp  Go to previous tab
  tabm  Move tab
  tabc  Close tab
  gt    Goto next tab
  gp    Goto previous tab    
  ```
</details>  
  
# Feature
  - Folding
  - Telescope (projects, buffers, filemanager, files, bookmarks, grep, grep <cword>)

# Plugins
  
```
Buffer      Telescope
Files       Telescope
Marks       Builtin
Go          go.nvim
Rename      Spectre
LSP         null-ls, lsp, lspinstall
Folding     Builtin
Terminal    Togglterm
Comment
Todo
  
use { "nvim-telescope/telescope-file-browser.nvim" }
use "tom-anders/telescope-vim-bookmarks.nvim"
use {'nvim-telescope/telescope-ui-select.nvim' }
use {'theHamsta/nvim-dap-virtual-text'}
use { "ghillb/cybu.nvim" }
use { "nyngwang/NeoZoom.lua", branch = "neo-zoom-original" }
 
```
<details>
  <summary>Spectre</summary>
    
  ```
  <CR>  Goto Current File
  c     Input Replace
  t     Toggle Line
  o     Show Options
  R     Replace
    
  ```
</details>
<details>
  <summary>Keymapping</summary>
    
  ```
  ```
</details>

