
# PA - Project Aliaser

> Navigate your filesystem with ease using custom shortcuts to your most frequently visited directories.

## 📋 Overview

Project Aliaser (PA) is a powerful command-line utility that enables you to create, manage, and use aliases for your directory paths. Instead of typing lengthy paths repeatedly, create short, memorable aliases and navigate to those locations with a single command.

## 🔧 Project Requirements

- Windows OS with PowerShell
- Go 1.22.1 or higher
- SQLite3 database (handled automatically)

## 📦 Dependencies

The project uses several Go packages:

- **Core functionality**:
  - `github.com/mattn/go-sqlite3` - SQLite database driver
  - `github.com/agnivade/levenshtein` - String similarity for fuzzy search
  
- **UI components**:
  - `github.com/charmbracelet/bubbletea` - Terminal UI framework
  - `github.com/charmbracelet/lipgloss` - Styling for terminal applications

## 🚀 Getting Started

### Installation

#### Option 1: Building from Source

1. **Build the project**:
   ```bash
   go build
   ```

2. **Setup PowerShell module**:
   - Ensure the PowerShell module and script are in the same directory as the executable
   - The module will be imported automatically when using the `pa.ps1` script
  
#### Option 2: Downloading the Pre-built Executable

1. **Download** the latest release (`project-aliaser@v0.1.0.zip`).  
2. **Extract** the files in a folder (e.g., `~\sufumi\project-aliaser\`)
3. **Add** the folder to your `$PATH` .  
4. **Run in PowerShell**:  
   ```ps1
   pa  # Launches the interactive TUI
   ```

> Note:
> Files  Included are as follows
> aliases.db, fa.exe, lib.psm1, mid_file.json, pa.ps1


### Creating Your First Alias

```powershell
# Add an alias to the current directory
./pa.ps1 add projects .

# Add an alias to a specific path
./pa.ps1 add documents "C:\Users\username\Documents" 
```

## 🔍 Usage

### Command-Line Interface

FA supports both direct command-line arguments and an interactive TUI:

```powershell
# List all aliases
./pa.ps1 list

# Jump to an aliased location
./pa.ps1 projects

# Delete an alias
./pa.ps1 delete projects
```

### Interactive UI

Launch the interactive UI by running FA without arguments:

```powershell
./pa.ps1
```

In the TUI:
- Use arrow keys (or j/k) to navigate
- Press Enter to select
- Press 'q' to quit
- Select "Search" to find and navigate to an aliased location

## 🔄 Core Features

### Alias Management

```powershell
# Create new alias
./pa.ps1 add dev "C:\Development"

# Alternative shorthand
./pa.ps1 ad dev "C:\Development"

# List all aliases
./pa.ps1 ls

# Delete an alias
./pa.ps1 dl dev
```

### Fuzzy Search

The fuzzy search capability allows you to find aliases even if you don't remember their exact names:

1. Launch FA without arguments
2. Select "Search" option
3. Start typing, and it will show matches based on Levenshtein distance
4. Navigate results using arrow keys
5. Press Enter to select and navigate to that location

### Implementation Details

FA stores aliases in a SQLite database (`aliases.db`) that is automatically created in the same directory as the executable. The database schema includes:

```sql
CREATE TABLE IF NOT EXISTS aliases (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    location TEXT NOT NULL UNIQUE
);
```

Navigating to an aliased location works via a temporary JSON file (`mid_file.json`) that facilitates communication between the Go executable and PowerShell.

## 📝 Examples

### Example 1: Setting up aliases for common development directories

```powershell
# Create aliases for different project directories
./pa.ps1 add go "C:\Development\Go"
./pa.ps1 add python "C:\Development\Python"
./pa.ps1 add web "C:\Development\Web"

# Now navigate using short commands
./pa.ps1 go    # Changes to C:\Development\Go
./pa.ps1 web   # Changes to C:\Development\Web
```

### Example 2: Using the interactive search

1. Run `./pa.ps1` to open the interactive UI
2. Select "Search"
3. Type part of an alias name (e.g., "dev")  
4. Use arrow keys to select from matching results
5. Press Enter to navigate to that location

## 🧰 Project Structure

- **`main.go`**: Entry point, handles command-line arguments
- **`api/api.go`**: Core functionality for alias management
- **`ui/ui.go`**: Terminal UI implementation
- **`pa.ps1`**: PowerShell script for user interaction
- **`lib.psm1`**: PowerShell module with helper functions

## 📈 Future Improvements

Potential enhancements for the project:

- Cross-platform support (Linux/macOS)
- Command history
- Alias categories/tags
- Export/import alias configurations
- Bash/Zsh shell support

## 🔚 Conclusion

FA brings the convenience of bookmarks to your command line, making filesystem navigation faster and more intuitive. Whether you're a developer juggling multiple project directories or a power user who values efficiency, FA streamlines your workflow by eliminating the need to memorize and type long paths.

Get started today and reclaim your time and mental energy for more important tasks!