
HISTSIZE=1000000
SAVEHIST=1000000
export EDITOR=lvim
export PATH=$HOME/bin:$HOME/local/bin:$HOME/cargo/bin:$PATH
export GOPATH=$HOME/go
# eval "$(zoxide init zsh)"
# eval "`pip completion --zsh`"

# NNN Bookmarks
NNN_BMS="h:~;"
NNN_BMS="g:~/go/src/github.com/GoogleCloudPlatform/microservices-demo/src;$NNN_BMS"
NNN_BMS="p:~/go/src/github.com/pellepedro;$NNN_BMS"
NNN_BMS="s:~/go/src/github.com/letsramp/skyramp/;$NNN_BMS"
NNN_BMS="m:~/go/src/github.com/letsramp/microservices-demo/;$NNN_BMS"
NNN_BMS="t:~/go/src/github.com/pellepedro/devcontainer/;$NNN_BMS"
export NNN_BMS

export CLICOLOR=1
export LSCOLORS=gxFxCxDxBxegedabagaced

# NNN Plugins
NNN_PLUG="e:-!sudo -E nvim $nnn*;"
NNN_PLUG="g:!lazygit;$NNN_PLUG"
NNN_PLUG="j:autojump;$NNN_PLUG"
NNN_PLUG="r:renamer;$NNN_PLUG"
export NNN_PLUG

BLK="42"
CHR="42"
DIR="42"
EXE="F7"
REG="F7"
HARDLINK="42"
SYMLINK="48"
MISSING="42"
ORPHAN="42"
FIFO="42"
SOCK="42"
OTHER="42"

export NNN_FCOLORS="$BLK$CHR$DIR$EXE$REG$HARDLINK$SYMLINK$MISSING$ORPHAN$FIFO$SOCK$OTHER"
export NNN_COLORS='2635'
export NNN_USE_EDITOR=1
export NNN_OPENER=lvim

# ================= Alias ==============================
alias g='lazygit'

alias tdev="TERM=xterm-256color tmux attach -t dev || TERM=xterm-256color tmux new -s dev"

alias dcu='docker compose up'
alias dcb='docker compose build'
alias dcp='docker compose push'

# Colorize grep output (good for log files)
alias grep='grep --color=auto'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'

# confirm before overwriting something
alias cp="cp -i"
alias mv='mv -i'
alias rm='rm -i'

# ================= Functions ==========================
function n()
{
    # Block nesting of nnn in subshells
    if [ -n $NNNLVL ] && [ "${NNNLVL:-0}" -ge 1 ]; then
        echo "nnn is already running"
        return
    fi

    # The default behaviour is to cd on quit (nnn checks if NNN_TMPFILE is set)
    # To cd on quit only on ^G, remove the "export" as in:
    #     NNN_TMPFILE="${XDG_CONFIG_HOME:-$HOME/.config}/nnn/.lastd"
    # NOTE: NNN_TMPFILE is fixed, should not be modified
    export NNN_TMPFILE="${XDG_CONFIG_HOME:-$HOME/.config}/nnn/.lastd"

    nnn -e

    if [ -f "$NNN_TMPFILE" ]; then
            . "$NNN_TMPFILE"
            rm -f "$NNN_TMPFILE" > /dev/null
     fi
}

function kl() {
  IMAGE=$(docker images | grep worker | awk '{print $1 ":" $2}' | fzf)
  kind load docker-image $IMAGE --name skyramp-local-local-cluster
}

# ================= Prompt =============================

autoload -Uz vcs_info
autoload -U colors && colors

# enable only git
zstyle ':vcs_info:*' enable git

# setup a hook that runs before every prompt.
precmd_vcs_info() { vcs_info }
precmd_functions+=( precmd_vcs_info )
setopt prompt_subst

bindkey -v
export KEYTIMEOUT=1

# ================= Vim Key ==============================


# Use vim keys in tab complete menu:
bindkey -M menuselect '^h' vi-backward-char
bindkey -M menuselect '^k' vi-up-line-or-history
bindkey -M menuselect '^l' vi-forward-char
bindkey -M menuselect '^j' vi-down-line-or-history
bindkey -M menuselect '^[[Z' vi-up-line-or-history
bindkey -v '^?' backward-delete-char

# Change cursor shape for different vi modes.
function zle-keymap-select () {
    case $KEYMAP in
        vicmd) echo -ne '\e[1 q';;      # block
        viins|main) echo -ne '\e[5 q';; # beam
    esac
}
zle -N zle-keymap-select
zle-line-init() {
    zle -K viins # initiate `vi insert` as keymap (can be removed if `bindkey -V` has been set elsewhere)
    echo -ne "\e[5 q"
}
zle -N zle-line-init
echo -ne '\e[5 q' # Use beam shape cursor on startup.
preexec() { echo -ne '\e[5 q' ;} # Use beam shape cursor for each new prompt.

# add a function to check for untracked files in the directory.
# from https://github.com/zsh-users/zsh/blob/master/Misc/vcs_info-examples
zstyle ':vcs_info:git*+set-message:*' hooks git-untracked
#
+vi-git-untracked(){
    if [[ $(git rev-parse --is-inside-work-tree 2> /dev/null) == 'true' ]] && \
        git status --porcelain | grep '??' &> /dev/null ; then
        hook_com[staged]+='!' # signify new files with a bang
    fi
}

zstyle ':vcs_info:*' check-for-changes true
zstyle ':vcs_info:git:*' formats " %{$fg[blue]%}(%{$fg[red]%}%m%u%c%{$fg[yellow]%}%{$fg[magenta]%} %b%{$fg[blue]%})"

PROMPT="%B%{$fg[blue]%}[%{$fg[white]%}%n%{$fg[red]%}@%{$fg[white]%}%m%{$fg[blue]%}] %(?:%{$fg_bold[green]%}➜ :%{$fg_bold[red]%}➜ )%{$fg[cyan]%}%c%{$reset_color%}"
PROMPT+="\$vcs_info_msg_0_ "


