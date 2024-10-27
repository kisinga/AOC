// swift-tools-version: 6.0

import types

class Day1: types.Solution {
    let filePath: String?
    let logger: Logger

    init(filePath: String?, logger: Logger) {
        self.filePath = filePath
        self.logger = logger
    }

    func solve() throws -> Int {
        self.logger.log("Solving day 1")
        return 0
    }
}
