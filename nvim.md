# Awesome Configs
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
  \4    ndap-ui    
  \5    dap
  \6.   dap debug
  R     Replace
    
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
  <summary>Indent Blankline</summary>
    
  ```
  ```
</details>
<details>
  <summary>DAP Debugging</summary>
    
  ```
  <F6>    Debug Test
  <F5>    Continue
  <F10>   Step Over
  <F11>   Step Into
  ```
</details>
