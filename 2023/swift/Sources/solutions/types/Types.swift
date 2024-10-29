// swift-tools-version: 6.0
public protocol Solvable {
    func solve() async throws -> Int
}

public protocol Logger {
    func log(_ message: String)
}
