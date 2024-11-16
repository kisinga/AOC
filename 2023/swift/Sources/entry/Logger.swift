// swift-tools-version: 6.0

import Solutions

class ConsoleLogger: Logger {
    let verbose: Bool
    let debug: Bool

    init(verbose: Bool, debug: Bool) {
        self.verbose = verbose
        self.debug = debug
    }
}

extension ConsoleLogger {
    func log(_ level: Level) {
        switch level {
        case let .debug(message):
            if debug {
                print(message)
            }
        case let .error(message):
            print(message)
        case let .info(message):
            if verbose {
                print(message)
            }
        case let .warn(message):
            if verbose && debug {
                print(message)
            }
        }
    }

    func forceLog(_ message: String) {
        print(message)
    }
}
