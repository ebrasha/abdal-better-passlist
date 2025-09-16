# Abdal Better PassList

<div align="center">
  <img src="shot.png" alt="Abdal Better PassList Screenshot" width="600">
</div>

**English** | [فارسی](README.fa.md)

## 🚀 About


**Abdal Better PassList** is a very powerful application for generating password lists using individuals’ identity information. The software can generate the probable passwords used by a victim for different countries by taking into account those countries’ documents and systems. It is designed for cybersecurity professionals, penetration testers, and security researchers. This powerful tool creates targeted password lists based on personal information and helps security professionals identify weak passwords and improve overall security posture.


## 🎯 Why This Software Was Created

In today's digital landscape, weak passwords remain one of the most significant security vulnerabilities. Traditional password attacks often fail because they rely on generic wordlists that don't account for personal information patterns. **Abdal Better PassList** addresses this gap by:

- **Personalized Targeting**: Generates password lists based on specific personal information
- **Multi-Country Support**: Handles document types and naming conventions from various countries
- **Intelligent Combinations**: Creates realistic password variations using names, dates, and numbers
- **Scalable Generation**: Supports multiple complexity levels for different testing scenarios

## ✨ Features

### 🌍 Multi-Country Support
- **All Countries**: Iran, USA, UK, Germany, France, Sweden, India, Japan, UAE, Canada, Australia, and other countries 
- **Country-Specific Documents**: Handles unique document types for each country
- **Localized Patterns**: Understands naming conventions and document formats

### 🔐 Three Complexity Levels
- **🟢 Normal**: Basic combinations for quick testing
- **🟡 Sensitive**: Complete combinations for thorough analysis  
- **🔴 Aggressive**: Maximum combinations for comprehensive coverage

### ⚡ Performance Features
- **Concurrent Processing**: Multi-threaded password generation (1-16 workers)
- **Progress Tracking**: Real-time generation progress with visual indicators
- **Memory Efficient**: Optimized algorithms for large password lists

### 🎮 User Experience
- **Interactive Mode**: Step-by-step guided interface
- **CLI Mode**: Command-line interface for automation
- **Cyberpunk UI**: Modern, colorful interface with phosphorescent styling
- **Error Logging**: Comprehensive error tracking and logging system

### 📁 Output Management
- **Smart Naming**: Files named as `FirstName_LastName_YYYY-MM-DD.txt`
- **Organized Storage**: Configurable output directories
- **Multiple Sessions**: Generate unlimited password lists in one session

## 🛠️  Build from Source

### Prerequisites
- Go 1.21 or later
- Windows, Linux, or macOS
 
### Building for Different Platforms
```bash
./build/build.bat
```

## 📖 Usage

### Interactive Mode
```bash
abdal-better-passlist --interactive
```

### Command Line Mode
```bash
# Basic usage
abdal-better-passlist --country ir --complexity sensitive --workers 4

# Advanced usage
abdal-better-passlist --country us --complexity aggressive --workers 8 --output /path/to/output
```

### Available Options
- `--country, -c`: Country code (ir, us, uk, de, fr, se, in, jp, ae, ca, au, general)
- `--complexity, -l`: Complexity level (normal, sensitive, aggressive)
- `--workers, -w`: Number of concurrent workers (1-16)
- `--output, -o`: Output directory
- `--interactive, -i`: Run in interactive mode

### Commands
- `test`: Test application functionality
- `help-detailed`: Show comprehensive help information

## 🌍 Supported Countries

| Country | Code | Key Documents |
|---------|------|---------------|
| 🇮🇷 Iran | `ir` | National ID, Birth Certificate, Passport, Driving License |
| 🇺🇸 United States | `us` | Driver's License, SSN, Passport, Birth Certificate |
| 🇬🇧 United Kingdom | `uk` | Passport, Driving Licence, NINO, Birth Certificate |
| 🇩🇪 Germany | `de` | Personalausweis, Passport, Tax ID, Registration |
| 🇫🇷 France | `fr` | Carte Nationale, Passport, Tax Number, Birth Certificate |
| 🇸🇪 Sweden | `se` | Personnummer, ID-card, Passport, BankID |
| 🇮🇳 India | `in` | Aadhaar, PAN, Passport, Voter ID, Driving Licence |
| 🇯🇵 Japan | `jp` | My Number, Passport, Driver's Licence, Pension IDs |
| 🇦🇪 UAE | `ae` | Emirates ID, Passport, Driving Licence, Residence Visa |
| 🇨🇦 Canada | `ca` | Driver's Licence, SIN, Passport, Provincial ID |
| 🇦🇺 Australia | `au` | Driver's Licence, TFN, Medicare Card, Passport |
| 🌍 General | `general` | National ID, Passport, Tax ID, Local Registration |

## 🔒 Security & Legal

### ⚠️ Important Disclaimer
This tool is designed for **educational purposes** and **authorized security testing only**. Users must:

- Comply with local laws and regulations
- Obtain proper authorization before testing
- Use responsibly and ethically
- Accept full responsibility for any misuse

### 🛡️ Ethical Use Guidelines
- Only test systems you own or have explicit permission to test
- Respect privacy and data protection laws
- Use findings to improve security, not exploit vulnerabilities
- Report security issues responsibly

## 🧪 Testing

```bash
# Test the application
abdal-better-passlist test

# Run with test mode (skip disclaimer)
abdal-better-passlist --test-mode
```

## 📊 Examples

### Example 1: Interactive Mode
```bash
abdal-better-passlist --interactive
```
- Guided country selection
- Step-by-step information collection
- Real-time complexity selection
- Progress tracking during generation

### Example 2: High-Performance Generation
```bash
abdal-better-passlist --country us --complexity aggressive --workers 12 --output /tmp/passwords
```
- Maximum complexity for comprehensive coverage
- 12 concurrent workers for speed
- Custom output directory

### Example 3: Quick Testing
```bash
abdal-better-passlist --country ir --complexity normal --workers 4
```
- Fast generation for initial testing
- Moderate complexity for efficiency
- Standard worker count

 

 
## 🐛 Reporting Issues

If you encounter any issues or have configuration problems, please reach out via email at Prof.Shafiei@Gmail.com. You can also report issues on GitLab or GitHub.

## ❤️ Donation

If you find this project helpful and would like to support further development, please consider making a donation:
- [Donate Here](https://alphajet.ir/abdal-donation)

## 🤵 Programmer

Handcrafted with Passion by **Ebrahim Shafiei (EbraSha)**
- **E-Mail**: Prof.Shafiei@Gmail.com
- **Telegram**: [@ProfShafiei](https://t.me/ProfShafiei)

## 📜 License

This project is licensed under the GPLv2 or later License.