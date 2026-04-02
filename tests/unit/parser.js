const fs = require('fs');
const path = require('path');
const { parse } = require('yaml');

class Parser {
    constructor(filePath) {
        this.filePath = path.resolve(filePath);
    }

    readFile() {
        try {
            return fs.readFileSync(this.filePath, 'utf8');
        } catch (error) {
            throw new Error(`Failed to read file: ${error.message}`);
        }
    }

    parseYAML() {
        const fileContent = this.readFile();
        try {
            return parse(fileContent);
        } catch (error) {
            throw new Error(`Failed to parse YAML: ${error.message}`);
        }
    }

    parseJSON() {
        const fileContent = this.readFile();
        try {
            return JSON.parse(fileContent);
        } catch (error) {
            throw new Error(`Failed to parse JSON: ${error.message}`);
        }
    }
}

module.exports = Parser;