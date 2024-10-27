// swift-tools-version: 6.0

public class Logger {
    let verbose: Bool

    public init(verbose: Bool) {
        self.verbose = verbose
    }

    public func log(_ message: String) {
        if verbose {
            print(message)
        }
    }
}

public protocol Solution {
    func solve(filePath: String, logger: Logger)
}
