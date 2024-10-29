import Foundation

enum FileReaderError: Error {
    case fileNotFound
    case invalidEncoding
    case readError(String)
}

struct LineReader {
    /// Reads a file line by line using Foundation's cross-platform APIs
    /// - Parameters:
    ///   - path: The path to the file
    ///   - encoding: The string encoding to use (defaults to UTF-8)
    ///   - handler: Closure to process each line
    static func readFile(
        at path: String,
        encoding: String.Encoding = .utf8,
        handler: (String) throws -> Void
    ) throws {
        let url = URL(fileURLWithPath: path)

        guard FileManager.default.fileExists(atPath: path) else {
            throw FileReaderError.fileNotFound
        }

        // Use String(contentsOf:) which is available across platforms
        let contents: String
        do {
            contents = try String(contentsOf: url, encoding: encoding)
        } catch {
            throw FileReaderError.invalidEncoding
        }

        // Use components(separatedBy:) which is available on all platforms
        let lines = contents.components(separatedBy: .newlines)

        // Process each line
        for line in lines {
            // Skip empty lines (optional - remove if you want to keep them)
            guard !line.isEmpty else { continue }
            try handler(line)
        }
    }

    /// Async version of the line reader (iOS 15+, macOS 12+, Linux with Swift 5.5+)
    /// Returns an async sequence of lines
    static func lines(
        from path: String,
        encoding: String.Encoding = .utf8
    ) -> AsyncThrowingStream<String, Error> {
        AsyncThrowingStream { continuation in
            do {
                try readFile(at: path, encoding: encoding) { line in
                    continuation.yield(line)
                }
                continuation.finish()
            } catch {
                continuation.finish(throwing: error)
            }
        }
    }
}
