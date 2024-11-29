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
struct Colors {
    static let reset = "\u{001B}[0;0m"
    static let black = "\u{001B}[0;30m"
    static let red = "\u{001B}[0;31m"
    static let green = "\u{001B}[0;32m"
    static let yellow = "\u{001B}[0;33m"
    static let blue = "\u{001B}[0;34m"
    static let magenta = "\u{001B}[0;35m"
    static let cyan = "\u{001B}[0;36m"
    static let white = "\u{001B}[0;37m"
}

extension ConsoleLogger {
    func log(_ level: LogLevel) {
        switch level {
        case let .debug(message):
            if debug {
                print(Colors.green + message + Colors.reset)
            }
        case let .error(message):
            print(Colors.red + message)
        case let .info(message):
            if verbose {
                print(Colors.white + message + Colors.reset)
            }
        case let .warn(message):
            if verbose && debug {
                print(Colors.yellow + message + Colors.reset)
            }
        }
    }

    func forceLog(_ message: String) {
        print(message)
    }
}
