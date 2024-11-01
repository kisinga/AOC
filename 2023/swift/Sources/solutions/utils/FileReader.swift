import Foundation

enum FileReadError: Error {
    case fileNotFound
    case invalidEncoding
    case readError(String)
}

struct FileReader {
    /// Returns an async stream of lines from a file at the given path.
    /// - Parameter path: The file path to read from
    /// - Returns: An async stream that yields lines from the file
    static func lines(from path: String) -> AsyncThrowingStream<String, Error> {
        AsyncThrowingStream { continuation in
            // Create FileHandle outside of Task to handle errors immediately
            guard let fileHandle = FileHandle(forReadingAtPath: path) else {
                continuation.finish(throwing: FileReadError.fileNotFound)
                return
            }

            // Don't allocate buffer until needed inside the task
            Task {
                do {
                    var lineBuffer = Data()

                    // Ensure file handle is closed even if task is cancelled
                    defer {
                        fileHandle.closeFile()
                    }

                    while true {
                        // Check for cancellation
                        try Task.checkCancellation()

                        let data = try fileHandle.read(upToCount: 64 * 1024) ?? Data()
                        if data.isEmpty {
                            // Process any remaining data in lineBuffer
                            if !lineBuffer.isEmpty {
                                guard let line = String(data: lineBuffer, encoding: .utf8) else {
                                    throw FileReadError.invalidEncoding
                                }
                                continuation.yield(line)
                            }
                            continuation.finish()
                            break
                        }

                        // Append new data to existing line buffer
                        lineBuffer.append(data)

                        // Process lines
                        while let newlineRange = lineBuffer.range(of: Data([0x0A])) {
                            let lineData = lineBuffer.subdata(in: 0..<newlineRange.lowerBound)
                            guard let line = String(data: lineData, encoding: .utf8) else {
                                throw FileReadError.invalidEncoding
                            }
                            continuation.yield(line)

                            // Remove processed line from buffer
                            lineBuffer.removeSubrange(0...newlineRange.upperBound - 1)
                        }
                    }
                } catch is CancellationError {
                    continuation.finish()
                } catch let error as FileReadError {
                    continuation.finish(throwing: error)
                } catch {
                    continuation.finish(
                        throwing: FileReadError.readError(error.localizedDescription))
                }
            }
        }
    }
}
