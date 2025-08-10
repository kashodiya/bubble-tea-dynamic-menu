
#!/bin/bash

echo "=== Bubble Tea TUI Command Runner Demo ==="
echo "This script will demonstrate how to use the TUI application."
echo ""

echo "1. Building the application..."
go build -o tui-app

echo ""
echo "2. Configuration file (config.json) contains:"
cat config.json | sed 's/^/   /'

echo ""
echo "3. To run the application, use: ./tui-app"
echo "   - Use arrow keys or j/k to navigate"
echo "   - Press Enter to execute a command"
echo "   - Press q to go back or quit"
echo ""

echo "4. Example commands in the configuration:"
echo "   - List Files: Lists files in the current directory"
echo "   - Check Disk Space: Shows disk space usage"
echo "   - System Info: Displays system information"
echo "   - Run Custom Script: Executes custom_script.sh"
echo ""

echo "Demo completed. Run ./tui-app to start the application."
