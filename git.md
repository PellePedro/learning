
## Sign Commit

```
gpg --full-generate-key

gpg --list-secret-keys --keyid-format=long

git config --global user.signingkey <key>
git config --global user.signingkey 3AA5C34371567BD2
```

### if you aren't using the GPG suite 

run the following command in the zsh shell to add the GPG key to your .zshrc file, if it exists, or your .zprofile file:
```
if [ -r ~/.zshrc ]; then echo 'export GPG_TTY=$(tty)' >> ~/.zshrc; \
  else echo 'export GPG_TTY=$(tty)' >> ~/.zprofile; fi
```

### Alt
```
brew install pinentry-mac
echo "pinentry-program $(which pinentry-mac)" >> ~/.gnupg/gpg-agent.conf
killall gpg-agent
```
