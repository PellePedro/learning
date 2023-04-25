# Astro Nvim
[rafaelderolez](https://github.com/rafaelderolez/astronvim)  
[mehalter](https://code.mehalter.com/AstroNvim_user/~files)  
[A-Lamia](https://github.com/A-Lamia/AstroNvim-conf)  


# Awesome Configs
https://github.com/herschel-ma/nvim-config

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

# Plugins
<details>
  <summary>Plugins</summary>
    
  ```
  Better Quick Fix (bqf)
  go.nvim
  Spectre
  Telescope (file, buffer, projects)
    
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

  
# Building
<details>
  <summary>MAC</summary>
    
  ```
#!/bin/bash

SDKROOT=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX11.0.sdk
INSTALL_PATH=$HOME/.local/apps/nvim

rm -rf neovim
rm -rf /usr/local/bin/nvim
rm -rf /usr/local/share/nvim
rm -rf /usr/local/lib/nvim

if [ -d "$INSTALL_PATH" ]; then rm -rf ${INSTALL_PATH}; fi

mkdir -p ${INSTALL_PATH}
brew install ninja libtool automake cmake pkg-config gettext
git clone https://github.com/neovim/neovim.git && cd neovim
git tag -d nightly
git tag nightly
make CMAKE_BUILD_TYPE=Release CMAKE_INSTALL_PREFIX=$INSTALL_PATH SDKROOT=$SDKROOT MACOSX_DEPLOYMENT_TARGET=11.0
make install
cd ..
rm -rf ~/neovim
ln -s ${INSTALL_PATH}/bin/nvim /usr/local/bin/nvim
ln -s ${INSTALL_PATH}/share/nvim /usr/local/share/nvim
ln -s ${INSTALL_PATH}/lib/nvim /usr/local/lib/nvim
  ```
</details>

