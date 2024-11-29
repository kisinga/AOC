// swift-tools-version: 6.0
public protocol Solvable {
    func solve() async throws -> Int
}

public protocol Logger {
    /// Log a message if verbose is enabled
    func log(_ level: LogLevel)
    /// Log a message regardless of verbose
    func forceLog(_ message: String)
}

@frozen public enum LogLevel {
    case info(_ message: String)
    case warn(_ message: String)
    case error(_ message: String)
    case debug(_ message: String)
}
