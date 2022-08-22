
## Installation
```
brew install git-delta
```


## .gitignore
```

[user]
	name = pellepedro
	email = Per.Pettersson@gmail.com
[filter "lfs"]
	smudge = git-lfs smudge -- %f
	process = git-lfs filter-process
	required = true
	clean = git-lfs clean -- %f
[init]
	defaultBranch = main
[core]
    pager = delta
[interactive]
    diffFilter = delta --color-only
[add.interactive]
    useBuiltin = false # required for git 2.37.0

[delta]
    features = side-by-side line-numbers decorations
    syntax-theme = Dracula
    navigate = true    # use n and N to move between diff sections
    light = false      # set to true if you're in a terminal w/ a light background color (e.g. the default macOS terminal)
    plus-style = syntax "#003800"
    minus-style = syntax "#3f0001"

[delta "decorations"]
    commit-decoration-style = bold yellow box ul
    file-style = bold yellow ul
    file-decoration-style = none
    hunk-header-decoration-style = cyan box ul

[delta "line-numbers"]
    line-numbers-left-style = cyan
    line-numbers-right-style = cyan
    line-numbers-minus-style = 124
    line-numbers-plus-style = 28

[pager]
    diff = delta
    log = delta
    reflog = delta
    show = delta

```



## Lazygit
```
git:
  paging:
    colorArg: always
    pager: delta --dark --paging=never
  pull:
    mode: 'rebase'
gui:
  theme:
    activeBorderColor:
      - blue
      - bold
    inactiveBorderColor:
      - white
    optionsTextColor:
      - blue
    selectedLineBgColor:
      - default # set to `default` to have no background colour
    selectedRangeBgColor:
      - default
    cherryPickedCommitBgColor:
      - cyan
    cherryPickedCommitFgColor:
      - blue
    unstagedChangesColor:
      - red

  showFileTree: true # for rendering changes files in a tree format
  showListFooter: false # for seeing the '5 of 20' message in list panels
  showRandomTip: false
  showBottomLine: false # for hiding the bottom information line (unless it has important information to tell you)
  showCommandLog: true
  showIcons: true

disableStartupPopups: true
notARepository: 'skip' # one of: 'prompt' | 'create' | 'skip'
os:
  editCommand: 'lvim'
```
