
<details>
  <summary>Go Workspaces</summary>
  
  https://go.dev/doc/tutorial/workspaces
  
  Topology
  ```golang
  /top-directory
    /module1
      go.mod
    /module2
      go.mod
  go.work
  ```
  Commands
  ```
  go work init ./module1
  go work use ./module2
  ```
  
</details>
