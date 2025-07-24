# portalctl

`portalctl` and `portal` make diving into "deep" directories easy!

# Installation

```shell
go install github.com/sirkon/portalctl@latest # Install portalctl.
portalctl setup                               # Adds portal to your shell and do some other stuff.
```

# What it is and how to use.

Imagine we have a project with deep hierachy of folders. You may do something like 
`cd project/internal/libray/internal/sublibrary/…/destination` one time, two times, but it quickly becomes
bothersome once you need to do this again and again and again. Meh.

`portalctl` and `portal` are means to do this pretty easy. The usage scenario with them is:

- Once in a directory you run
  ```shell
  portalctl here <dest> 
  ```
  Where `<dest>` is some identifier.
- When you need to cd in that folder again you just
  ```shell
  portal <dest>
  ```
And your `bash` or `zsh` will move you there.

# More on the `portalctl`.

You can do some management with it:

| command           | info                                                                                                                                                                         |
|-------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `here <name>`     | Adds current directory under the given name - a new portal.                                                                                                                  |
| `delete <name>`   | Removes both name and path once registered with `here` under the given name.                                                                                                 |
| `show <name>`     | Prints a path once registered under the given name.                                                                                                                          |
| `list`            | List all available "portals" with name and its path.                                                                                                                         |
| `prefix <prefix>` | List all names having the given `<prefix>`                                                                                                                                   |
| `log-compact`     | Removes all log entries related to removed portals.                                                                                                                          |
| `version`         | Prints version and exits.                                                                                                                                                    |
| `setup`           | Registers `<portal>` shell function. Registers shell completion for `<portal>`. `<portal>` is just `portal` by default. Can be customized, see help for the `setup` command. |

# What is under the hood.

`portalctl` works like a database: 

- once you do "write" type operations, it adds a new record to a log file placed in `<user cache directory>/portalctl/`.
- it reads records in the log file and do a response logic once you do "read" type operation. It is always "full scan",
 so that's why `log-compact` is provided. Albeit, the log shouldn't be that huge to really care about this: there are
 only a few directories you need to go into regularly in any reasonable scenario.
