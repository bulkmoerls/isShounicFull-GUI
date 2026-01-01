# IsShounicFullGUI

A GUI version of IsShounicFull that checks if Shounic Trenches (A TF2 Server) is full. \
\
![Built with Go](https://badges.penpow.dev/badges/built-with/go/cozy.svg)

## "What is the meaning of this?"

If you have been playing TF2 Community Servers, you may have seen an 100 player server named "Shounic Trenches (USA Chicago)". And be honest, the player count is big. And I mean **BIG** when I say it. This server is **IMPOSSIBLE** to join when the server is at peak times. Although not really, I can't sit here waiting and waiting and waiting for a person to finally leave the server so I can join and have a good time. That's why I built IsShounicFull. \
ISF-GUI originally uses [Rust](https://rust-lang.org/) as the programming language as it is robust and fast, but I find it harder to learn since I am still working on learning the basics of code, so I currently use [Go](https://go.dev) for now. The original ISF is coded using [Python](https://www.python.org), and I still have the original file. If you want, I can post it on GitHub right now.

## Installation

## Windows

Since it's built with GTK3, you may normally experience errors like libcairo-2.dll not found. To fix that, install [GTK for Windows Runtime](https://github.com/tschoonj/GTK-for-Windows-Runtime-Environment-Installer).

## Linux

If you get errors saying that some .so file isn't present, you didn't install GTK3. basically the same thing, but it takes one sudo command to have it installed.
| Distro        | Command                       |
| ------------- | ----------------------------- |
| Ubuntu/Debian | `sudo apt install libgtk-3-0` |
| Arch Linux    | `sudo pacman -S gtk3`         |
| Fedora        | `sudo dnf install gtk3`       |

## Building

_soon..._

## faq

Why is the full server ip exposed in the source code?
: The Server IP is required to Query Info.

## credits

go-a2s for making this project possible \
gotk3 for the same thing, the ui and borrowing code from example repo. \
@shounic for making the chaotic 100 player server (it was cinema) \
