# Ascii Art Web

## Description
The **Ascii Art Web** is a web-based tool that takes user input (text) and generates corresponding ASCII art. The server processes text inputs to create a visual representation of the text using ASCII characters. This tool can be used to convert any textual input into a styled ASCII format for creative or decorative purposes.

## Authors
- **cgkaldan** - Christos Gkaldanidis
- **sgougoul** - Stelios Gougoulakis

Creators and primary developers

## Usage: How to Run
1. **Clone the repository**:
    ```bash
    git clone https://platform.zone01.gr/git/cgkaldan/ascii-art-web.git
    cd ascii-art-web

2. **Start the server**
    ```bash
    go run .

3. **Open the server in your browser**
    Open the web browser and go to the following URL:
    http://localhost:8080
    
4. **Have fun**


## Implementation Details: Algorithm

The ASCII Art Generator works by transforming a given text input into a visual representation using ASCII characters. The following outlines the high-level steps and considerations for the algorithm:

### 1. **Text Processing**
   - **Input Parsing**: The server parses the input text, breaking it down into individual words or lines.
   - **Text Formatting**: Depending on the desired style, the input might be split into rows or columns to better fit the ASCII art layout. This also helps in managing very long inputs, potentially breaking them into chunks for easier processing.

### 2. **Output Display**
   - **Final Formatting**: Once the ASCII art is generated, the server returns the ASCII art to the user. The output can be formatted in a simple text box or styled using CSS to improve readability.
     - The ASCII art is returned as plain text wrapped in HTML tags for styling and display.
   
   - **Error Handling**: If something goes wrong during the generation (e.g., memory issues, invalid input), the server returns a descriptive error message and logs the error for further debugging.

### 3. **Additional Features**
   - **Custom Styling**: The user can specify text color and background color, which are applied to the output via CSS. This allows users to personalize the appearance of their generated ASCII art.

### Algorithm Flow:
1. **Receive input** from the user.
2. **Reads a banner file** from a directory with banners
2. **Generate ASCII art** for each character using a predefined set of patterns.
4. **Assemble the ASCII art** line by line.
5. **Return the ASCII art** to the user, with optional error handling or progress indication.



