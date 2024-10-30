import Foundation

enum FileReaderError: Error {
    case fileNotFound
    case invalidEncoding
    case readError(String)
}

struct LineReader {
    /// Async version of the line reader (iOS 15+, macOS 12+, Linux with Swift 5.5+)
    /// Returns an async sequence of lines
    static func lines(from path: String) -> AsyncThrowingStream<String, Error> {
        AsyncThrowingStream { continuation in
            do {
                let fileHandle = try FileHandle(forReadingFrom: URL(fileURLWithPath: path))
                Task {
                    defer { try? fileHandle.close() }
                    for try await line in fileHandle.bytes.lines {
                        guard !line.isEmpty else { continue }
                        continuation.yield(line)
                    }
                    continuation.finish()
                }
            } catch {
                continuation.finish(throwing: error)
            }
        }
    }
}
