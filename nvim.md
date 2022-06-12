# Awesome Configs
[Allaman](https://github.com/Allaman/nvim)
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
  <summary>DAP Keymappings</summary>
    
  ```  
  vim.keymap.set("n", "<F5>", ":lua require'dap'.continue()<CR>")
  vim.keymap.set("n", "<F10>", ":lua require'dap'.step_over()<CR>")
  vim.keymap.set("n", "<F11>", ":lua require'dap'.step_into()<CR>")
  vim.keymap.set("n", "<F12>", ":lua require'dap'.step_out()<CR>")
  vim.keymap.set("n", "<Leader>b", ":lua require'dap'.toggle_breakpoint()<CR>")
  vim.keymap.set("n", "<Leader>B", ":lua require'dap'.set_breakpoint(vim.fn.input('Breakpoint condition: '))<CR>")
  vim.keymap.set("n", "<Leader>lp", ":lua require'dap'.set_breakpoint(nil, nil, vim.fn.input('Log point message: '))<CR>")
  vim.keymap.set("n", "<Leader>dr", ":lua require'dap'.repl.open()<CR>")
  vim.keymap.set("n", "<Leader>dl", ":lua require'dap'.run_last()<CR>")
  vim.keymap.set("n", "<Leader>td", ":lua require('dap-go').debug_test()<CR>")

require('dap-go').setup()
require("dapui").setup()
require('nvim-dap-virtual-text').setup()

local dap, dapui = require("dap"), require("dapui")
dap.listeners.after.event_initialized["dapui_config"] = function()
  dapui.open()
end
dap.listeners.before.event_terminated["dapui_config"] = function()
  dapui.close()
end
dap.listeners.before.event_exited["dapui_config"] = function()
  dapui.close()
end
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
