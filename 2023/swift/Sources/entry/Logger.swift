// swift-tools-version: 6.0

import Solutions

class ConsoleLogger: Logger {
    let verbose: Bool

    init(verbose: Bool) {
        self.verbose = verbose
    }
}

extension ConsoleLogger {
    func log(_ message: String) {
        if verbose {
            print(message)
        }
    }

    func forceLog(_ message: String) {
        print(message)
    }
}
