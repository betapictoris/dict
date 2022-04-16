# Dictionary CLI [![Go](https://github.com/BetaPictoris/dict/actions/workflows/go.yml/badge.svg)](https://github.com/BetaPictoris/dict/actions/workflows/go.yml)
View the dictionary through the CLI

![dict-final](https://user-images.githubusercontent.com/65696362/163287317-f2e6d271-7785-4baa-b8b0-ac69a63bba41.png)


## Installation
### From release
```bash
curl -LO https://github.com/BetaPictoris/dict/releases/latest/download/dict    # Download the latest binary.
sudo install -Dt /usr/local/bin -m 755 dict                                    # Install Dictionary CLI to "/usr/local/bin" with the mode "755"
```

### Build from source 

If you're running Arch/Manjaro (or derivatives) you can install the git build from the AUR with your preferred AUR helper. e.g
```
paru -S dict-git
```
For all other distros, read below.

#### Dependencies

You need Go installed to build this program. You can install it from your distro's repository using one of the following commands:

```bash
# Arch/Manjaro (and derivatives)
sudo pacman -S go

# Debian/Ubuntu (and derivatives)
sudo apt install golang-go
```

Alternatively, you can install it from go's official website: https://go.dev/doc/install

```bash
git clone git@github.com:BetaPictoris/dict.git      # Clone the repository
cd dict                                             # Change into the repository's directory
make                                                # Build Wiki CLI
sudo make install                                   # Install Wiki CLI to "/usr/local/bin" with the mode "755"
sudo make uninstall                                 # Uninstall
```

### User install
If you don't have access to `sudo` on your system you can install to your user's `~/.local/bin` directory with this command: 
```bash
install -Dt ~/.local/bin -m 755 dict
```
